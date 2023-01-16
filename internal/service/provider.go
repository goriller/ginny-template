package service

import (
	"context"

	pb "MODULE_NAME/api/proto"
	"MODULE_NAME/internal/cache"
	"MODULE_NAME/internal/repo"

	"github.com/google/wire"
	"github.com/goriller/ginny"
	"github.com/goriller/ginny/errs"
)

// ProviderSet
var ProviderSet = wire.NewSet(
	cache.ProviderSet,
	NewService,
	RegisterService,
)

// Service the instance for grpc proto.
type Service struct {
	pb.UnimplementedSayServer
	// Introduce new dependencies here, exp:
	cache          *cache.RedisCache
	userRepository *repo.UserRepo
}

// NewService new service that implement hello
func NewService(
	cache *cache.RedisCache,
	userRepository *repo.UserRepo,
) (*Service, error) {
	return &Service{
		cache:          cache,
		userRepository: userRepository,
	}, nil
}

// RegisterService
func RegisterService(ctx context.Context, sev *Service) ginny.RegistrarFunc {
	return func(app *ginny.Application) error {
		// 注入错误码
		errs.RegisterErrorCodes(pb.ErrorCode_name)
		// 注册gRPC服务
		app.Server.RegisterService(ctx, &pb.Say_ServiceDesc, sev)

		if app.Option.HttpAddr != "" {
			// 注册http服务
			if err := pb.RegisterSayHandlerServer(ctx, app.Server.ServeMux(), sev); err != nil {
				return err
			}
		}
		return nil
	}
}
