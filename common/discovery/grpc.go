package discovery

import (
	"context"
	"math/rand/v2"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ServiceConnection(ctx context.Context, serviceName string, registry Registry) (*grpc.ClientConn, error) {
	addrs, err := registry.Discover(ctx, serviceName)
	if err != nil {
		return nil, err
	}

	// pick one target randomly
	targetIndex := rand.IntN(len(addrs))
	target := addrs[targetIndex]

	// create tracker handler
	handler := otelgrpc.NewClientHandler()

	// create the connection and handle both return values
	conn, err := grpc.NewClient(
		target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// middleware
		grpc.WithStatsHandler(handler),
	)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
