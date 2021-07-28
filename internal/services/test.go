package services

import (
	"github.com/google/wire"
	"go.uber.org/zap"
)

// TestServiceProvider
var TestServiceProvider = wire.NewSet(NewTestService, wire.Bind(new(ITestService), new(*TestService)))

// ITestService
type ITestService interface{}

// TestService
type TestService struct {
	logger *zap.Logger
}

// NewTestService
func NewTestService() *TestService {
	return &TestService{}
}

// GetInfo
func (p *TestService) GetInfo() string {
	return "test"
}
