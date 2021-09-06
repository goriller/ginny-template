package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/gorillazer/ginny-serve/http"
	// "github.com/gorillazer/ginny/res"
)

func CreateInitHandlerFn(
// HANDLE 锚点请勿删除! Do not delete this line!
) http.InitHandlers {
	return func(r *gin.Engine) {
		// 在此定义路由规则 Define routing rules here, exp:
		// r.GET("/test/:id", res.Wrapper(test.Get))
	}
}

var ProviderSet = wire.NewSet(
	// HANDLE_PROVIDER 锚点请勿删除! Do not delete this line!
	CreateInitHandlerFn,
)
