package handlers

import (
	"moduleName/api/proto"
	"moduleName/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// TestHandlerProvider
var TestHandlerProvider = wire.NewSet(NewTestHandler, wire.Bind(new(ITestHandler), new(*TestHandler)))

// ITestHandler
type ITestHandler interface {
	Get(c *gin.Context)
	GetRPC(c *gin.Context)
}

// TestHandler
type TestHandler struct {
	v            *viper.Viper
	logger       *zap.Logger
	testService  *services.TestService
	detailClient proto.DetailsClient
}

// NewTestHandler
func NewTestHandler(
	v *viper.Viper,
	logger *zap.Logger,
	testService *services.TestService,
	detailClient proto.DetailsClient,
) *TestHandler {
	return &TestHandler{
		v:            v,
		logger:       logger.With(zap.String("type", "TestHandler")),
		testService:  testService,
		detailClient: detailClient,
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

func (t *TestHandler) GetRPC(c *gin.Context) {
	req := &proto.GetDetailRequest{
		Id: 1,
	}
	t.logger.Info(t.v.GetString("consul.address"))
	p, err := t.detailClient.Get(c, req)
	if err != nil {
		t.logger.Error("TestHandler.GetRPC", zap.Error(err))
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, p.Name)
}
