// +build wireinject

package main

import (
	"MODULE_NAME/internal/handlers"
	// CMD_IMPORT 锚点请勿删除! Do not delete this line!

	"github.com/google/wire"
	"github.com/gorillazer/ginny"
	config "github.com/gorillazer/ginny-config"
	jaeger "github.com/gorillazer/ginny-jaeger"
	log "github.com/gorillazer/ginny-log"
	http "github.com/gorillazer/ginny-serve/http"
)

// Create http/grpc Serve
func newServe(
	hs *http.Server,
	// CMD_SERVEPARAM 锚点请勿删除! Do not delete this line!

) ([]ginny.Serve, error) {
	return []ginny.Serve{
		ginny.HttpServe(hs),
		// CMD_SERVEFUNC 锚点请勿删除! Do not delete this line!

	}, nil
}

// appProvider
var appProvider = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	jaeger.ProviderSet,
	newServe, ginny.AppProviderSet)

// CreateApp
func CreateApp(name string) (*ginny.Application, error) {
	panic(wire.Build(wire.NewSet(
		handlers.ProviderSet,
		// CMD_PROVIDERSET 锚点请勿删除! Do not delete this line!

		appProvider,
	)))
}
