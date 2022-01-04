package main

import (
	"context"
	"fmt"
	"tiny/base/grpc/pb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:6001", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()
	c := pb.NewBaseClient(conn)
	r, err := c.Hello(context.Background(), &pb.Info{Msg: "hello world"})
	if err == nil {
		fmt.Println(r.Resp)
	}
}
