package repositories

import (
	"github.com/google/wire"
	mysql "github.com/gorillazer/ginny-mysql"
)

var ProviderSet = wire.NewSet(
	mysql.Provider,
	UserRepositoryProvider,
)
