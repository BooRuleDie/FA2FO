package discovery

import (
	"context"
	"math/rand/v2"

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

	// create the connection and handle both return values
	conn, err := grpc.NewClient(
		target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
