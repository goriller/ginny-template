package rpc_servers

import (
	"MODULE_NAME/api/proto"

	"github.com/google/wire"
	"github.com/gorillazer/ginny-serve/grpc"
	stdGrpc "google.golang.org/grpc"
)

// CreateInitServerFn
func CreateInitServerFn(
	d *DetailsServer,
) grpc.InitServers {
	return func(s *stdGrpc.Server) {
		proto.RegisterDetailsServer(s, d)
	}
}

// ProviderSet
var ProviderSet = wire.NewSet(
	NewDetailsServer,
	CreateInitServerFn,
)
