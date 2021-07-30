package grpcservers

import "github.com/google/wire"

var DetailsServerProvider = wire.NewSet(NewDetailsServer, wire.Bind(new(IDetailsServer), new(*DetailsServer)))

type IDetailsServer interface{}
type DetailsServer struct{}

func NewDetailsServer() *DetailsServer {
	return &DetailsServer{}
}
