package repo

import (
	"context"
	"errors"
	"strings"

	"MODULE_NAME/internal/repo/entity"

	"github.com/google/wire"
	orm "github.com/goriller/ginny-gorm"
	"github.com/goriller/ginny-util/validation"
	"gorm.io/gorm"
	// DATABASE_LIB 锚点请勿删除! Do not delete this line!
)

// UserRepoProvider
var UserRepoProvider = wire.NewSet(NewUserRepo)

// UserRepo
type UserRepo struct {
	orm *orm.ORM
	// mongo *mongo.Manager
	entity *entity.UserEntity
	// STRUCT_ATTR 锚点请勿删除! Do not delete this line!
}

// NewUserRepo
func NewUserRepo(
	// redis *redis.Manager,
	orm *orm.ORM,
	userEntity *entity.UserEntity,
	// mongo *mongo.Manager,
	// FUNC_PARAM 锚点请勿删除! Do not delete this line!
) (*UserRepo, error) {
	return &UserRepo{
		orm:    orm,
		entity: userEntity,
		// FUNC_ATTR 锚点请勿删除! Do not delete this line!
	}, nil
}

// Count
func (p *UserRepo) Count(ctx context.Context, where Query) (total int64, err error) {
	db := p.orm.RDB().Table(p.entity.TableName())
	if where.QueryStr != "" && where.Attrs != nil {
		db = db.Where(where.QueryStr, where.Attrs...)
	}
	err = db.Count(&total).Error
	return
}

// Find
func (p *UserRepo) Find(ctx any, where Query, order []string) (
	result *entity.UserEntity, err error) {
	if order == nil {
		order = []string{"id desc"}
	}
	db := p.orm.RDB().Table(p.entity.TableName())
	if where.QueryStr != "" && where.Attrs != nil {
		db = db.Where(where.QueryStr, where.Attrs...)
	}
	err = db.Order(strings.Join(order, ",")).First(&result).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return
}

// FindAll
func (p *UserRepo) FindAll(ctx context.Context, where Query,
	order []string, opt ...int) (result []entity.UserEntity, err error) {
	if order == nil {
		order = []string{"id desc"}
	}
	db := p.orm.RDB().Table(p.entity.TableName())
	if where.QueryStr != "" && where.Attrs != nil {
		db = db.Where(where.QueryStr, where.Attrs...)
	}
	db = db.Order(strings.Join(order, ","))
	var (
		limit  = 1000
		offset = 0
	)
	if len(opt) > 0 {
		limit = opt[0]
	}
	db = db.Limit(limit)
	if len(opt) == 2 {
		offset = opt[1]
	}
	db = db.Offset(offset)
	err = db.Find(&result).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return
}

// Insert
func (p *UserRepo) Insert(ctx context.Context,
	entity *entity.UserEntity) (int64, error) {
	if err := validation.Validate(entity); err != nil {
		return 0, err
	}
	result := p.orm.RDB().Table(p.entity.TableName()).Create(entity)
	return entity.Id, result.Error
}

// Update
func (p *UserRepo) Update(ctx context.Context, where Query,
	update entity.UserEntity) (int64, error) {
	if err := validation.Validate(update); err != nil {
		return 0, err
	}
	db := p.orm.RDB().Table(p.entity.TableName())
	if where.QueryStr != "" && where.Attrs != nil {
		db = db.Where(where.QueryStr, where.Attrs...)
	}
	result := db.Updates(update)
	return result.RowsAffected, result.Error
}

// Delete
func (p *UserRepo) Delete(ctx context.Context,
	where Query) (int64, error) {
	var t *entity.UserEntity
	db := p.orm.RDB().Table(p.entity.TableName())
	if where.QueryStr != "" && where.Attrs != nil {
		db = db.Where(where.QueryStr, where.Attrs...)
	}
	result := db.Delete(t)
	return result.RowsAffected, result.Error
}

// PDelete physical deletion
func (p *UserRepo) PDelete(ctx context.Context,
	where Query) (int64, error) {
	var t *entity.UserEntity
	db := p.orm.RDB().Table(p.entity.TableName())
	if where.QueryStr != "" && where.Attrs != nil {
		db = db.Where(where.QueryStr, where.Attrs...)
	}
	result := db.Unscoped().Delete(t)
	return result.RowsAffected, result.Error
}
