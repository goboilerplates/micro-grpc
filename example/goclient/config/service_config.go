package config

import (
	"context"

	"github.com/goboilerplates/micro-rpc/example/goclient/request"
	"github.com/goboilerplates/micro-rpc/example/goclient/runner"
	"google.golang.org/grpc"
)

// SetRequesters sets all requesters and run these requesters.
func SetRequesters(ctx context.Context, conn *grpc.ClientConn) {
	getSampleRequester := request.GetSampleRequesterImpl{
		Context: ctx,
		Conn:    conn,
	}
	runner.RunGetSampleRequests(getSampleRequester)
}
