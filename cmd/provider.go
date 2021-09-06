// +build wireinject

package main

import (
	"MODULE_NAME/internal/handlers"
	// CMD_IMPORT 锚点请勿删除! Do not delete this line!

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
	jaeger.ProviderSet,
	http.ProviderSet,
	handlers.ProviderSet,
	// CMD_PROVIDERSET 锚点请勿删除! Do not delete this line!

	// consul.ProviderSet,
	// grpc.ProviderSet,
	// server.ProviderSet,
	// client.ProviderSet,
	// services.ProviderSet,
	// repositories.ProviderSet,
	appProvider,
)

var appProvider = wire.NewSet(newServe, ginny.AppProviderSet)

// Create http/grpc Serve
func newServe(
	hs *http.Server,
	// cli *consul.Client,
	// gs *grpc.Server,
	// CMD_SERVEPARAM 锚点请勿删除! Do not delete this line!
) ([]ginny.Serve, error) {
	return []ginny.Serve{
		ginny.HttpServe(hs),
		// ginny.GrpcServeWithConsul(gs, cli),
		// CMD_SERVEFUNC 锚点请勿删除! Do not delete this line!
	}, nil
}

// CreateApp
func CreateApp(name string) (*ginny.Application, error) {
	panic(wire.Build(providerSet))
}
