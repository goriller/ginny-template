package rpc_servers

import (
	"moduleName/api/proto"

	"github.com/google/wire"
	"github.com/gorillazer/ginny-serve/grpc"

	stdgrpc "google.golang.org/grpc"
)

func CreateInitServerFn(
	ps *DetailsServer,
) grpc.InitServers {
	return func(s *stdgrpc.Server) {
		proto.RegisterDetailsServer(s, ps)
	}
}

var ProviderSet = wire.NewSet(
	NewDetailsServer,
	CreateInitServerFn,
)
