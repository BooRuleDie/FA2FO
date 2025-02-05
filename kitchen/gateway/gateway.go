package gateway

import (
	"context"

	pb "github.com/BooRuleDie/Microservice-in-Go/common/api"
)

type KitchenGateway interface {
	UpdateOrder(ctx context.Context, o *pb.Order) error
}
