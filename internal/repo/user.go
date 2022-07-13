package repo

import (
	"context"

	"MODULE_NAME/internal/config"
	"MODULE_NAME/internal/repo/entity"

	"github.com/google/wire"
	mysql "github.com/goriller/ginny-mysql"
	"github.com/goriller/ginny/logger"
	"go.uber.org/zap"
	// mongo "github.com/goriller/ginny-mongo"
	// redis "github.com/goriller/ginny-redis"
	// DATABASE_LIB 锚点请勿删除! Do not delete this line!
)

// UserRepoProvider
var UserRepoProvider = wire.NewSet(NewUserRepo,
	wire.Bind(new(IUserRepo), new(*UserRepo)))

// IUserRepo
type IUserRepo interface {
	GetUser(ctx context.Context) (*entity.UserEntity, error)
}

// UserRepo
type UserRepo struct {
	config *config.Config
	// redis *redis.Manager
	mysql *mysql.SqlBuilder
	// mongo *mongo.Manager
	// STRUCT_ATTR 锚点请勿删除! Do not delete this line!
}

// NewUserRepo
func NewUserRepo(
	config *config.Config,
	// redis *redis.Manager,
	mysql *mysql.SqlBuilder,
	// mongo *mongo.Manager,
	// FUNC_PARAM 锚点请勿删除! Do not delete this line!
) *UserRepo {
	return &UserRepo{
		config: config,
		// redis: redis,
		mysql: mysql,
		// mongo: mongo,
		// FUNC_ATTR 锚点请勿删除! Do not delete this line!
	}
}

func (p *UserRepo) GetUser(ctx context.Context) (*entity.UserEntity, error) {
	r := &entity.UserEntity{}
	log := logger.WithContext(ctx).With(zap.String("action", "GetUser"))
	// if err := p.mysql.Find(ctx, r, r.TableName(), nil); err != nil {
	// 	log.Error("", zap.Error(err))
	// 	return nil, err
	// }

	// if _, err := p.mongo.Database.Collection(r.TableName()).Find(ctx, nil); err != nil {
	// 	return nil, err
	// }
	// p.redis.DB().Get(ctx, r.TableName()).Result()
	log.Debug("GetUser")
	return r, nil
}
