package grpcservers

import (
	"github.com/google/wire"
	"github.com/gorillazer/ginny/transports/grpc"

	stdgrpc "google.golang.org/grpc"
)

func CreateInitServerFn(
	ps *DetailsServer,
) grpc.InitServers {
	return func(s *stdgrpc.Server) {
		// proto.RegisterDetailsServer(s, ps)
	}
}

var ProviderSet = wire.NewSet(
	NewDetailsServer,
	CreateInitServerFn,
)
