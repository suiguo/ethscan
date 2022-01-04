package endpoint

import (
	"context"
	"fmt"
	"tiny/base/grpc/pb"
)

func TEndPoint(ctx context.Context, request interface{}) (response interface{}, err error) {
	switch req := request.(type) {
	case *pb.Info:
		response := &pb.Resp{Resp: req.Msg}
		return response, nil
	}
	return nil, fmt.Errorf("unknow proto")
}
