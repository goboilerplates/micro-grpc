package config

import (
	"github.com/goboilerplates/core"
	"github.com/goboilerplates/micro-rpc/api"
	"github.com/goboilerplates/micro-rpc/proto"
	"google.golang.org/grpc"
)

// RegisterServiceServers .
func RegisterServiceServers(server *grpc.Server) {
	proto.RegisterGetSamplesServer(server, &api.GetSamplesAPI{
		Interactor: core.CreateDefaultGetSamples(),
	})
}
