package service

import (
	"context"
	"tiny/base/grpc/pb"

	kit_grpc "github.com/go-kit/kit/transport/grpc"
)

type BaseServerImpl struct {
	Kit *kit_grpc.Server
	pb.BaseServer
}

func (b *BaseServerImpl) Hello(ctx context.Context, req *pb.Info) (*pb.Resp, error) {
	_, resp, err := b.Kit.ServeGRPC(ctx, req)
	return resp.(*pb.Resp), err
}
