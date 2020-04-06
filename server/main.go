package main

import (
	"context"
	"net"

	"github.com/dvbnrg/ToyotaCodingChallengeTake2/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
}

func main() {
	listener, err := net.Listen("tcp", ":80")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterHelloServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(err)
	}
}

func (s *server) Hello(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a := request.GetHelloReq()

	return &proto.Response{HelloRes: a}, nil
}
