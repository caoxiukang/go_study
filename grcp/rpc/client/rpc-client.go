package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"os"
	pd "grpc/pd"
	"time"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// 1.建立链接
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// 2.获取相应服务的连接
	greeter := pd.NewGreeterClient(conn)
	xiukang := pd.NewXiukangrClient(conn)
	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	// 1秒的上下文
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := greeter.SayHello(ctx, &pd.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)

	r, err = xiukang.SayXiukang(ctx, &pd.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)

}
