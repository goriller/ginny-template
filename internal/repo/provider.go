package repo

import (
	"github.com/google/wire"
	mysql "github.com/goriller/ginny-mysql"
	redis "github.com/goriller/ginny-redis"
	// DATABASE_LIB 锚点请勿删除! Do not delete this line!
)

var ProviderSet = wire.NewSet(

	redis.Provider,
	RedisCacheProvider,
	mysql.Provider,
	// mongo.Provider,
	// DATABASE_PROVIDER 锚点请勿删除! Do not delete this line!
	UserRepoProvider,
	// REPO_PROVIDER 锚点请勿删除! Do not delete this line!
)
