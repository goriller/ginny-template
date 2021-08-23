package repositories

import (
	"github.com/google/wire"
	"github.com/gorillazer/ginny/db/mysql"
)

var ProviderSet = wire.NewSet(
	mysql.Provider,
	UserRepositoryProvider,
)
