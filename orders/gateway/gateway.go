package gateway

import (
	"context"

	pb "github.com/BooRuleDie/Microservice-in-Go/common/api"
)

type StockGateway interface {
	CheckIfItemIsInStock(ctx context.Context, customerID string, items []*pb.ItemsWithQuantity) (bool, []*pb.Item, error)
}
