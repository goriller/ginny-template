package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/gorillazer/ginny-serve/http"
	"github.com/gorillazer/ginny/res"
)

func CreateInitHandlerFn(
	test *TestHandler,
) http.InitHandlers {
	return func(r *gin.Engine) {
		r.GET("/test/:id", res.Wrapper(test.Get))
		r.GET("/test1/:id", res.Wrapper(test.GetRPC))
	}
}

var ProviderSet = wire.NewSet(
	http.ProviderSet,
	TestHandlerProvider,
	CreateInitHandlerFn,
)
