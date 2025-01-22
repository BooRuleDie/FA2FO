package stripe

import (
	"github.com/BooRuleDie/Microservice-in-Go/common"
	pb "github.com/BooRuleDie/Microservice-in-Go/common/api"
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/checkout/session"
)

var (
	gatewaySuccessUrl string = common.EnvString("GATEWAY_PAYMENT_SUCCESS", "http://localhost:8080/success.html")
)

type Stripe struct {
}

func NewStripe() *Stripe {
	return &Stripe{}
}

func (s *Stripe) CreatePaymentLink(p *pb.Order) (string, error) {

	items := []*stripe.CheckoutSessionLineItemParams{}
	for _, item := range p.Items {
		items = append(items, &stripe.CheckoutSessionLineItemParams{
			Price:    stripe.String(item.PriceID),
			Quantity: stripe.Int64(int64(item.Quantity)),
		})
	}

	params := &stripe.CheckoutSessionParams{
		LineItems: items,
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(gatewaySuccessUrl),
	}
	result, err := session.New(params)
	if err != nil {
		return "", err
	}
	return result.URL, nil
}
