package main

import (
	"context"
	"fmt"
	"time"

	"github.com/dvbnrg/ToyotaCodingChallengeTake2/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:80", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := proto.NewHelloServiceClient(conn)

	req := &proto.Request{HelloReq: "Hello World"}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.Hello(ctx, req)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
