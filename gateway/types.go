package main

import pb "github.com/BooRuleDie/Microservice-in-Go/common/api"

type CreateOrderRequest struct {
	Order         *pb.Order
	RedirectToUrl string
}
