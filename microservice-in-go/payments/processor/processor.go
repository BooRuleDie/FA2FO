package processor

import pb "github.com/BooRuleDie/Microservice-in-Go/common/api"

type PaymentProcessor interface {
	CreatePaymentLink(*pb.Order) (string, error)
}
