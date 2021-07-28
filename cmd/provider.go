// +build wireinject

package main

import (
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
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// providerSet
var providerSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	consul.ProviderSet,
	jaeger.ProviderSet,
	http.ProviderSet,
	grpc.ProviderSet,
	handlers.ProviderSet,
	services.ProviderSet,
	appProviderSet,
)

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

// CreateApp
func CreateApp(cf string) (*ginny.Application, error) {
	panic(wire.Build(providerSet))
}
