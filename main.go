package main

import (
	"fmt"
	"net"
	"tiny/base/grpc/pb"
	"tiny/base/grpc/service"
	t_endpoint "tiny/base/kit/endpoint"
	t_transport "tiny/base/kit/transport"

	kit_grpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

func main() {
	// var p interface{}
	// p = map[string]interface{}{"1": 1, "2": "456", "3": nil}
	// fmt.Println("start[", time.Now().UnixMilli())
	// for k := 0; k < 10000000; k++ {
	// 	switch t := p.(type) {
	// 	case map[string]interface{}:
	// 		t["1"] = 2
	// 		continue
	// 	}
	// }
	// fmt.Println("end  [", time.Now().UnixMilli())
	srv := kit_grpc.NewServer(t_endpoint.TEndPoint, t_transport.Decode, t_transport.Encode)

	lis, err := net.Listen("tcp", ":6001")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	s := grpc.NewServer()
	service := &service.BaseServerImpl{Kit: srv}
	pb.RegisterBaseServer(s, service)
	s.Serve(lis)
}
