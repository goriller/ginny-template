// +build wireinject

package main

import (
	// "moduleName/internal/grpcservers"
	"moduleName/internal/handlers"
	"moduleName/internal/repositories"
	"moduleName/internal/services"

	"github.com/google/wire"
	"github.com/gorillazer/ginny"
	config "github.com/gorillazer/ginny-config"

	// consul "github.com/gorillazer/ginny-consul"
	jaeger "github.com/gorillazer/ginny-jaeger"
	log "github.com/gorillazer/ginny-log"

	// grpc "github.com/gorillazer/ginny-serve/grpc"
	http "github.com/gorillazer/ginny-serve/http"
)

// providerSet
var providerSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	// consul.ProviderSet,
	jaeger.ProviderSet,
	http.ProviderSet,
	// grpc.ProviderSet,
	handlers.ProviderSet,
	// grpcservers.ProviderSet,
	services.ProviderSet,
	repositories.ProviderSet,
	appProvider,
)

var appProvider = wire.NewSet(newServe, ginny.AppProviderSet)

// Create http/grpc Serve
func newServe(
	hs *http.Server,
	// cli *consul.Client
	// gs *grpc.Server,
) ([]ginny.Serve, error) {
	return []ginny.Serve{
		ginny.HttpServe(hs),
		// ginny.GrpcServeWithConsul(gs, cli),
	}, nil
}

// CreateApp
func CreateApp(name string) (*ginny.Application, error) {
	panic(wire.Build(providerSet))
}
