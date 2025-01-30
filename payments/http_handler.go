package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	pb "github.com/BooRuleDie/Microservice-in-Go/common/api"
	"github.com/BooRuleDie/Microservice-in-Go/common/broker"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/webhook"
	"go.opentelemetry.io/otel"
)

type PaymentHTTPHandler struct {
	channel *amqp.Channel
}

func NewPaymentHTTPHandler(channel *amqp.Channel) *PaymentHTTPHandler {
	return &PaymentHTTPHandler{channel: channel}
}

func (p *PaymentHTTPHandler) registerRoutes(router *http.ServeMux) {
	router.HandleFunc("/webhook", p.handleCheckoutWebhook)
}

func (p *PaymentHTTPHandler) handleCheckoutWebhook(w http.ResponseWriter, req *http.Request) {
	const MaxBodyBytes = int64(65536)
	req.Body = http.MaxBytesReader(w, req.Body, MaxBodyBytes)

	body, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading request body: %v\n", err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	// fmt.Printf("Got body: %s", body)

	// Pass the request body and Stripe-Signature header to ConstructEvent, along with the webhook signing key
	// Use the secret provided by Stripe CLI for local testing
	// or your webhook endpoint's secret.
	event, err := webhook.ConstructEvent(body, req.Header.Get("Stripe-Signature"), endpointStripeSecret)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error verifying webhook signature: %v\n", err)
		w.WriteHeader(http.StatusBadRequest) // Return a 400 error on a bad signature
		return
	}

	if event.Type == stripe.EventTypeCheckoutSessionCompleted ||
		event.Type == stripe.EventTypeCheckoutSessionAsyncPaymentSucceeded {
		var cs stripe.CheckoutSession
		err := json.Unmarshal(event.Data.Raw, &cs)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if cs.PaymentStatus == stripe.CheckoutSessionPaymentStatusPaid {
			log.Printf("payment for checkout session %v succeeded!", cs.ID)

			ctx, cancel := context.WithTimeout(req.Context(), 5*time.Second)
			defer cancel()

			// marshal the order for fanout
			o := &pb.Order{
				ID:          cs.Metadata["orderID"],
				CustomerID:  cs.Metadata["customerID"],
				Status:      "paid",
				PaymentLink: "",
			}
			marshalledOrder, err := json.Marshal(o)
			if err != nil {
				log.Fatal(err.Error())
			}

			tr := otel.Tracer("amqp")
			spanName := fmt.Sprintf("AMQP - publish - %s", broker.OrderPaidEvent)
			amqpContext, messageSpan := tr.Start(ctx, spanName)
			defer messageSpan.End()
			headers := broker.InjectAMQPHeaders(amqpContext)

			// publish the message
			p.channel.PublishWithContext(ctx, broker.OrderPaidEvent, "", false, false, amqp.Publishing{
				ContentType:  "application/json",
				Body:         marshalledOrder,
				DeliveryMode: amqp.Persistent,
				Headers:      headers,
			})

			log.Println("Message published order.paid")
		}
	}

	w.WriteHeader(http.StatusOK)
}
