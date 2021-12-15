package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/safeie/grpc-flatbuffers-example/api/models"
	"google.golang.org/grpc"

	flatbuffers "github.com/google/flatbuffers/go"
)

var (
	greetings = [...]string{"Hi", "Hallo", "Ciao"}
)

type greeterServer struct {
	models.UnimplementedGreeterServer
}

func (s *greeterServer) SayHello(ctx context.Context, request *models.HelloRequest) (*flatbuffers.Builder, error) {
	v := request.Name()
	var m string
	if v == nil {
		m = "Unknown"
	} else {
		m = string(v)
	}
	b := flatbuffers.NewBuilder(0)
	idx := b.CreateString("Welcome " + m)
	models.HelloReplyStart(b)
	models.HelloReplyAddMessage(b, idx)
	b.Finish(models.HelloReplyEnd(b))
	return b, nil
}

func (s *greeterServer) SayManyHellos(request *models.ManyHellosRequest, stream models.Greeter_SayManyHellosServer) error {
	v := request.Name()
	var m string
	if v == nil {
		m = "Unknown"
	} else {
		m = string(v)
	}
	num := request.NumGreetings()
	b := flatbuffers.NewBuilder(0)

	for _, greeting := range greetings {
		idx := b.CreateString(fmt.Sprintf("%s %s ,num %d", greeting, m, num))
		models.HelloReplyStart(b)
		models.HelloReplyAddMessage(b, idx)
		b.Finish(models.HelloReplyEnd(b))
		if err := stream.Send(b); err != nil {
			return err
		}
	}

	return nil
}

func newServer() *greeterServer {
	return &greeterServer{}
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 3000))
	if err != nil {
		log.Fatalf("Falied to listen: %v", err)
	}

	codec := &flatbuffers.FlatbuffersCodec{}
	grpcServer := grpc.NewServer(grpc.ForceServerCodec(codec))
	models.RegisterGreeterServer(grpcServer, newServer())
	if err := grpcServer.Serve(lis); err != nil {
		fmt.Println(err)
		panic(err)
	}
}
