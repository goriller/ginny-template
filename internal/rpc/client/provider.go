package client

import (
	"github.com/google/wire"
)

// ProviderSet
var ProviderSet = wire.NewSet(

	NewDetailsClient,
)
