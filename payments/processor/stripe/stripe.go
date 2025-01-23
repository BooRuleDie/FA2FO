package stripe

import (
	"fmt"

	"github.com/BooRuleDie/Microservice-in-Go/common"
	pb "github.com/BooRuleDie/Microservice-in-Go/common/api"
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/checkout/session"
)

var (
	gatewaySuccessUrl string = common.EnvString("GATEWAY_PAYMENT_SUCCESS", "http://localhost:8081/success.html?customerID=%s&orderID=%s")
	gatewayCancelUrl string = common.EnvString("GATEWAY_PAYMENT_CANCEL", "http://localhost:8081/cancel.html")
)

type Stripe struct {
}

func NewStripe() *Stripe {
	return &Stripe{}
}

func (s *Stripe) CreatePaymentLink(o *pb.Order) (string, error) {

	items := []*stripe.CheckoutSessionLineItemParams{}
	for _, item := range o.Items {
		items = append(items, &stripe.CheckoutSessionLineItemParams{
			Price:    stripe.String(item.PriceID),
			Quantity: stripe.Int64(int64(item.Quantity)),
		})
	}

	gatewaySuccessUrl = fmt.Sprintf(gatewaySuccessUrl, o.CustomerID, o.ID) 

	params := &stripe.CheckoutSessionParams{
		LineItems: items,
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(gatewaySuccessUrl),
		CancelURL: stripe.String(gatewayCancelUrl),
	}
	result, err := session.New(params)
	if err != nil {
		return "", err
	}
	return result.URL, nil
}
