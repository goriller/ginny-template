// +build wireinject

package main

import (
	// "moduleName/internal/grpcservers"
	"moduleName/internal/handlers"
	"moduleName/internal/services"

	"github.com/google/wire"
	"github.com/gorillazer/ginny"
	"github.com/gorillazer/ginny/config"
	"github.com/gorillazer/ginny/log"
	"github.com/gorillazer/ginny/naming/consul"
	"github.com/gorillazer/ginny/tracing/jaeger"
	"github.com/gorillazer/ginny/transports/grpc"
	"github.com/gorillazer/ginny/transports/http"
)

// providerSet
var providerSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	consul.ProviderSet,
	jaeger.ProviderSet,
	http.ProviderSet, grpc.ProviderSet,
	handlers.ProviderSet,
	// grpcservers.ProviderSet,
	services.ProviderSet,
	appProvider,
)

var appProvider = wire.NewSet(newServe, ginny.AppProviderSet)

// Create http/grpc Serve
func newServe(
	hs *http.Server,
	// gs *grpc.Server,
) ([]ginny.Serve, error) {
	return []ginny.Serve{
		ginny.HttpServe(hs),
		// ginny.GrpcServe(gs),
	}, nil
}

// CreateApp
func CreateApp(name string) (*ginny.Application, error) {
	panic(wire.Build(providerSet))
}
