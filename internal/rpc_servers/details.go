package rpc_servers

import (
	"context"
	"moduleName/api/proto"
	"moduleName/internal/services"
	"time"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.uber.org/zap"
)

// DetailsServerProvider
var DetailsServerProvider = wire.NewSet(NewDetailsServer, wire.Bind(new(IDetailsServer), new(*DetailsServer)))

// IDetailsServer
type IDetailsServer interface {
	Get(ctx context.Context, req *proto.GetDetailRequest) (*proto.Detail, error)
}

// DetailsServer
type DetailsServer struct {
	logger      *zap.Logger
	testService *services.TestService
}

// NewDetailsServer
func NewDetailsServer(
	logger *zap.Logger,
	testService *services.TestService,
) (*DetailsServer, error) {
	return &DetailsServer{}, nil
}

func (s *DetailsServer) Get(ctx context.Context, req *proto.GetDetailRequest) (*proto.Detail, error) {
	p, err := s.testService.Get(ctx, req.Id)
	if err != nil {
		return nil, errors.Wrap(err, "details grpc service get detail error")
	}
	ct := timestamppb.New(time.Time{})
	resp := &proto.Detail{
		Id:          req.Id,
		Name:        p,
		Price:       0,
		CreatedTime: ct,
	}

	return resp, nil
}
