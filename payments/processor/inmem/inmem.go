package inmem

import pb "github.com/BooRuleDie/Microservice-in-Go/common/api"

type inmem struct {
}

func NewInmem() *inmem {
	return &inmem{}
}

func (i *inmem) CreatePaymentLink(*pb.Order) (string, error) {
	return "http://example.com", nil
}
