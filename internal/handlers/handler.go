package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/gorillazer/ginny/transports/http"
)

func CreateInitHandlerFn(
	test *TestHandler,
) http.InitHandlers {
	return func(r *gin.Engine) {
		r.GET("/test/:id", test.Get)
	}
}

var ProviderSet = wire.NewSet(
	TestHandlerProvider,
	CreateInitHandlerFn,
)
