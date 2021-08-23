package services

import (
	"context"
	"moduleName/internal/repositories"

	"github.com/google/wire"
	"go.uber.org/zap"
)

// TestServiceProvider
var TestServiceProvider = wire.NewSet(NewTestService, wire.Bind(new(ITestService), new(*TestService)))

// ITestService
type ITestService interface{}

// TestService
type TestService struct {
	logger         *zap.Logger
	userRepository *repositories.UserRepository
}

// NewTestService
func NewTestService(logger *zap.Logger,
	userRepository *repositories.UserRepository) *TestService {
	return &TestService{
		logger:         logger,
		userRepository: userRepository,
	}
}

// GetInfo
func (p *TestService) GetInfo(ctx context.Context) (string, error) {
	user, err := p.userRepository.GetUser(ctx)
	if err != nil {
		return "", err
	}
	return user.Name, nil
}
