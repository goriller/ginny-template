// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/gorillazer/ginny"
	"github.com/gorillazer/ginny/config"
	"github.com/gorillazer/ginny/log"
	"github.com/gorillazer/ginny/naming/consul"
	"github.com/gorillazer/ginny/tracing/jaeger"
	"github.com/gorillazer/ginny/transports/grpc"
	"github.com/gorillazer/ginny/transports/http"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"moduleName/internal/handlers"
	"moduleName/internal/services"
)

// Injectors from provider.go:

// CreateApp
func CreateApp(cf string) (*ginny.Application, error) {
	viper, err := config.New(cf)
	if err != nil {
		return nil, err
	}
	logOptions, err := log.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	logger, err := log.New(logOptions)
	if err != nil {
		return nil, err
	}
	mainOptions, err := newOptions(viper, logger)
	if err != nil {
		return nil, err
	}
	httpOptions, err := http.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	testService := services.NewTestService()
	testHandler := handlers.NewTestHandler(logger, testService)
	initHandlers := handlers.CreateInitHandlerFn(testHandler)
	configuration, err := jaeger.NewConfiguration(viper, logger)
	if err != nil {
		return nil, err
	}
	tracer, err := jaeger.New(configuration)
	if err != nil {
		return nil, err
	}
	engine := http.NewRouter(httpOptions, logger, initHandlers, tracer)
	consulOptions, err := consul.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	client, err := consul.New(consulOptions)
	if err != nil {
		return nil, err
	}
	server, err := http.New(httpOptions, logger, engine, client)
	if err != nil {
		return nil, err
	}
	application, err := newApp(mainOptions, logger, server)
	if err != nil {
		return nil, err
	}
	return application, nil
}

// provider.go:

// providerSet
var providerSet = wire.NewSet(log.ProviderSet, config.ProviderSet, consul.ProviderSet, jaeger.ProviderSet, http.ProviderSet, grpc.ProviderSet, handlers.ProviderSet, services.ProviderSet, appProviderSet)

var appProviderSet = wire.NewSet(newApp, newOptions)

// options
type options struct {
	Name string
}

// newOptions
func newOptions(v *viper.Viper, logger *zap.Logger) (*options, error) {
	var err error
	o := new(options)
	if err = v.UnmarshalKey("app", o); err != nil {
		return nil, errors.Wrap(err, "unmarshal app option error")
	}

	logger.Info("load application options success")

	return o, err
}

// newApp
func newApp(o *options, logger *zap.Logger, hs *http.Server) (*ginny.Application, error) {
	a, err := ginny.New(o.Name, logger, ginny.HttpServerOption(hs))

	if err != nil {
		return nil, errors.Wrap(err, "new app error")
	}

	return a, nil
}