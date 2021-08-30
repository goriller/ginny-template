package handlers

import (
	"moduleName/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go.uber.org/zap"
)

// TestHandlerProvider
var TestHandlerProvider = wire.NewSet(NewTestHandler, wire.Bind(new(ITestHandler), new(*TestHandler)))

// ITestHandler
type ITestHandler interface{}

// TestHandler
type TestHandler struct {
	logger      *zap.Logger
	testService *services.TestService
}

// NewTestHandler
func NewTestHandler(
	logger *zap.Logger,
	testService *services.TestService,
) *TestHandler {
	return &TestHandler{
		logger:      logger.With(zap.String("type", "TestHandler")),
		testService: testService,
	}
}

func (t *TestHandler) Get(c *gin.Context) {
	t.logger.Debug("TestHandler.Get", zap.Any("TestHandler.Get", c.Params))
	name, err := t.testService.GetInfo(c)
	if err != nil {
		t.logger.Error("TestHandler.Get", zap.Error(err))
		c.JSON(http.StatusBadGateway, err.Error())
		return
	}
	c.JSON(http.StatusOK, name)
}
