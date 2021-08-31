package repositories

import (
	"context"

	"github.com/google/wire"
	mysql "github.com/gorillazer/ginny-mysql"
	"go.uber.org/zap"
)

var UserRepositoryProvider = wire.NewSet(NewUserRepository, wire.Bind(new(IUserRepository), new(*UserRepository)))

type IUserRepository interface {
	GetUser(ctx context.Context) (*UserRepository, error)
}
type UserRepository struct {
	Id   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`

	sqlBuilder *mysql.SqlBuilder
	logger     *zap.Logger
}

func NewUserRepository(sqlBuilder *mysql.SqlBuilder,
	logger *zap.Logger) *UserRepository {
	return &UserRepository{
		logger:     logger,
		sqlBuilder: sqlBuilder,
	}
}

func (p *UserRepository) GetUser(ctx context.Context) (*UserRepository, error) {
	user := &UserRepository{}
	err := p.sqlBuilder.Find(ctx, user, "user", nil)
	if err != nil {
		p.logger.Error("UserRepository.GetUser", zap.Error(err))
		return nil, err
	}
	return user, nil
}
