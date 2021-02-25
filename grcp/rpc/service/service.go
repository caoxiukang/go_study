package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pd "grpc/pd"
	"log"
	"net"
)

const (
	port = ":50051"
)

// 1.定义服务
type server struct{}        //服务对象
type xiukangServer struct{} //服务对象

// 2.实现服务的方法
// SayHello 实现服务的接口 在proto中定义的所有服务都是接口
func (s *server) SayHello(ctx context.Context, in *pd.HelloRequest) (*pd.HelloReply, error) {
	fmt.Println("request: ", in.Name)
	return &pd.HelloReply{Message: "Hello " + in.Name}, nil
}

func (xiu *xiukangServer) SayXiukang(ctx context.Context, in *pd.HelloRequest) (*pd.HelloReply, error) {
	fmt.Println("request: ", in.Name)
	return &pd.HelloReply{Message: "Hello xiukang：" + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer() //起一个服务

	// 3.注册服务
	pd.RegisterGreeterServer(s, &server{})
	pd.RegisterXiukangrServer(s, &xiukangServer{})
	// 注册反射服务 这个服务是CLI使用的 跟服务本身没有关系
	reflection.Register(s)
	// 服务监听阻塞中若需要结合其他的服务（如  gin,beego）则需要使用协程
	log.Println("rpc服务已经开启 0.0.0.0:50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
