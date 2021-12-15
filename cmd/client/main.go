package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"time"

	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/safeie/grpc-flatbuffers-example/api/models"
	"google.golang.org/grpc"
)

var (
	addr = "3000"
	name = flag.String("name", "Flatbuffers", "name to be sent go server :D")
)

func printSayHello(client models.GreeterClient, name string) {
	log.Printf("Name to be sent (%s)", name)
	b := flatbuffers.NewBuilder(0)
	i := b.CreateString(name)
	models.HelloRequestStart(b)
	models.HelloRequestAddName(b, i)
	b.Finish(models.HelloRequestEnd(b))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	request, err := client.SayHello(ctx, b, grpc.CallContentSubtype("flatbuffers"))
	if err != nil {
		log.Fatalf("%v.SayHello(_) = _, %v: ", client, err)
	}
	log.Printf("server said %q", request.Message())
}

func printSayManyHello(client models.GreeterClient, name string, num int32) {
	log.Printf("Name to be sent (%s), num to be sent (%d)", name, num)
	b := flatbuffers.NewBuilder(0)
	i := b.CreateString(name)
	models.ManyHellosRequestStart(b)
	models.ManyHellosRequestAddName(b, i)
	models.ManyHellosRequestAddNumGreetings(b, num)
	b.Finish(models.ManyHellosRequestEnd(b))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.SayManyHellos(ctx, b, grpc.CallContentSubtype("flatbuffers"))
	if err != nil {
		log.Fatalf("%v.SayManyHellos(_) = _, %v", client, err)
	}
	for {
		request, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.SayManyHellos(_) = _, %v", client, err)
		}
		log.Printf("server said %q", request.Message())
	}
}

func main() {
	flag.Parse()
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%s", addr), grpc.WithInsecure(), grpc.WithCodec(flatbuffers.FlatbuffersCodec{}))
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()
	client := models.NewGreeterClient(conn)
	printSayHello(client, *name)

	num := rand.Int31()
	printSayManyHello(client, *name, num)
}
