package main

import (
	"net/http"

	"github.com/BooRuleDie/Microservice-in-Go/common"
	pb "github.com/BooRuleDie/Microservice-in-Go/common/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type handler struct {
	client pb.OrderServiceClient
}

func NewHandler(client pb.OrderServiceClient) *handler {
	return &handler{client: client}
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customerID}/orders", h.HandleCreateOrder)
}

func (h *handler) HandleCreateOrder(rw http.ResponseWriter, r *http.Request) {
	customerID := r.PathValue("customerID")

	var items []*pb.ItemsWithQuantity
	if err := common.ReadJSON(r, &items); err != nil {
		common.WriteError(rw, http.StatusBadRequest, err.Error())
		return // don't continue executing rest of the code if it's an error
	}

	if err := validateItems(items); err != nil {
		common.WriteError(rw, http.StatusBadRequest, err.Error())
		return // don't continue executing rest of the code if it's an error
	}

	o, err := h.client.CreateOrder(r.Context(), &pb.CreateOrderRequest{
		CustomerID: customerID,
		Items:      items,
	})
	// we shouldn't write the grpc server's error directly if we want a
	// consistent error format in the response as grpc package manipulates the
	// error response just like that:
	// { "error": "rpc error: code = Unknown desc = heeeyo" }
	// but the format I want to use is this:
	// { "error": "heeeyo"}
	// to achieve that we need to handle the error like that:
	rStatus := status.Convert(err)
	if rStatus != nil {
		if rStatus.Code() != codes.InvalidArgument {
			common.WriteError(rw, http.StatusBadRequest, rStatus.Message())
			return // don't continue executing rest of the code if it's an error
		}

		common.WriteError(rw, http.StatusInternalServerError, err.Error())
		return // don't continue executing rest of the code if it's an error
	}

	common.WriteJSON(rw, http.StatusCreated, o)
}

func validateItems(items []*pb.ItemsWithQuantity) error {
	if len(items) == 0 {
		return common.ErrNoItems
	}

	for _, i := range items {
		if i.ID == "" {
			return common.ErrNoID
		}

		if i.Quantity <= 0 {
			return common.ErrInvalidQuantity
		}
	}

	return nil
}
