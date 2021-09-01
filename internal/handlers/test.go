package handlers

import (
	"MODULE_NAME/api/proto"
	"MODULE_NAME/configs"
	"MODULE_NAME/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/gorillazer/ginny/errs"
	"github.com/gorillazer/ginny/res"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// TestHandlerProvider
var TestHandlerProvider = wire.NewSet(NewTestHandler, wire.Bind(new(ITestHandler), new(*TestHandler)))

// ITestHandler
type ITestHandler interface {
	Get(c *gin.Context) (*res.Response, error)
	GetRPC(c *gin.Context) (*res.Response, error)
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

func (t *TestHandler) Get(c *gin.Context) (*res.Response, error) {
	t.logger.Debug("TestHandler.Get", zap.Any("TestHandler.Get", c.Params))
	name, err := t.testService.GetInfo(c)
	if err != nil {
		t.logger.Error("TestHandler.Get", zap.Error(err))
		return nil, errs.New(configs.ERR_GETINFO, configs.GetErrMsg(configs.ERR_GETINFO))
	}
	return res.Success(name), nil
}

func (t *TestHandler) GetRPC(c *gin.Context) (*res.Response, error) {
	req := &proto.GetDetailRequest{
		Id: 1,
	}
	t.logger.Info(t.v.GetString("consul.address"))
	p, err := t.detailClient.Get(c, req)
	if err != nil {
		t.logger.Error("TestHandler.GetRPC", zap.Error(err))
		return res.Fail(errs.New(configs.ERR_GETINFO, configs.GetErrMsg(configs.ERR_GETINFO))), nil
	}
	return res.Success(p.Name), nil
}
