package rpc

import (
	"github.com/google/wire"
	consul "github.com/gorillazer/ginny-consul"
	"github.com/gorillazer/ginny-serve/grpc"
)

// ProviderSet
var ProviderSet = wire.NewSet(
	grpc.ProviderSet,
	consul.ProviderSet,
)
