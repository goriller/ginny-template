package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/gorillazer/ginny-serve/http"
)

func CreateInitHandlerFn(
	test *TestHandler,
) http.InitHandlers {
	return func(r *gin.Engine) {
		r.GET("/test/:id", test.Get)
		r.GET("/test1/:id", test.GetRPC)
	}
}

var ProviderSet = wire.NewSet(
	TestHandlerProvider,
	CreateInitHandlerFn,
)
