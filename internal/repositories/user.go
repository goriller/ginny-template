package repositories

import "github.com/google/wire"

// UserRepositoryProvider
var UserRepositoryProvider = wire.NewSet(wire.Struct(new(UserRepository), "*"), wire.Bind(new(IUserRepository), new(*UserRepository)))

// IUserRepository
type IUserRepository interface{}

// UserRepository
type UserRepository struct {
	Id   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
}
