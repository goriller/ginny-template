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
type ITestService interface {
	Get(ctx context.Context, Id uint64) (string, error)
	GetInfo(ctx context.Context) (string, error)
}

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

// Get
func (p *TestService) Get(ctx context.Context, Id uint64) (string, error) {
	return "user.Name", nil
}

// GetInfo
func (p *TestService) GetInfo(ctx context.Context) (string, error) {
	user, err := p.userRepository.GetUser(ctx)
	if err != nil {
		return "", err
	}
	return user.Name, nil
}
