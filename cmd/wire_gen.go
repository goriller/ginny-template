// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"MODULE_NAME/internal/handlers"
	"github.com/google/wire"
	"github.com/gorillazer/ginny"
	"github.com/gorillazer/ginny-config"
	"github.com/gorillazer/ginny-jaeger"
	"github.com/gorillazer/ginny-log"
	"github.com/gorillazer/ginny-serve/http"
)

// Injectors from provider.go:

// CreateApp
func CreateApp(name string) (*ginny.Application, error) {
	viper, err := config.New(name)
	if err != nil {
		return nil, err
	}
	options, err := log.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	logger, err := log.New(options)
	if err != nil {
		return nil, err
	}
	option, err := ginny.NewOption(viper, logger)
	if err != nil {
		return nil, err
	}
	serverOption, err := http.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	configuration, err := jaeger.NewConfiguration(viper, logger)
	if err != nil {
		return nil, err
	}
	tracer, err := jaeger.New(configuration)
	if err != nil {
		return nil, err
	}
	initHandlers := handlers.CreateInitHandlerFn()
	engine := http.NewRouter(serverOption, logger, tracer, initHandlers)
	server, err := http.NewServer(serverOption, logger, engine)
	if err != nil {
		return nil, err
	}
	v, err := newServe(server)
	if err != nil {
		return nil, err
	}
	application, err := ginny.NewApp(option, logger, v...)
	if err != nil {
		return nil, err
	}
	return application, nil
}

// provider.go:

// providerSet
var providerSet = wire.NewSet(log.ProviderSet, config.ProviderSet, jaeger.ProviderSet, handlers.ProviderSet, appProvider)

var appProvider = wire.NewSet(newServe, ginny.AppProviderSet)

// Create http/grpc Serve
func newServe(
	hs *http.Server,

) ([]ginny.Serve, error) {
	return []ginny.Serve{ginny.HttpServe(hs)}, nil
}
