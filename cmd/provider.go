// +build wireinject

package main

import (
	"MODULE_NAME/internal/handlers"
	"MODULE_NAME/internal/repositories"
	"MODULE_NAME/internal/rpc_clients"
	"MODULE_NAME/internal/rpc_servers"
	"MODULE_NAME/internal/services"

	"github.com/google/wire"
	"github.com/gorillazer/ginny"
	config "github.com/gorillazer/ginny-config"
	jaeger "github.com/gorillazer/ginny-jaeger"
	log "github.com/gorillazer/ginny-log"

	consul "github.com/gorillazer/ginny-consul"
	grpc "github.com/gorillazer/ginny-serve/grpc"
	http "github.com/gorillazer/ginny-serve/http"
)

// providerSet
var providerSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	jaeger.ProviderSet,
	http.ProviderSet,
	grpc.ProviderSet,
	handlers.ProviderSet,
	consul.ProviderSet,
	rpc_servers.ProviderSet,
	rpc_clients.ProviderSet,
	services.ProviderSet,
	repositories.ProviderSet,
	appProvider,
)

var appProvider = wire.NewSet(newServe, ginny.AppProviderSet)

// Create http/grpc Serve
func newServe(
	hs *http.Server,
	cli *consul.Client,
	gs *grpc.Server,
) ([]ginny.Serve, error) {
	return []ginny.Serve{
		ginny.HttpServe(hs),
		ginny.GrpcServeWithConsul(gs, cli),
	}, nil
}

// CreateApp
func CreateApp(name string) (*ginny.Application, error) {
	panic(wire.Build(providerSet))
}
