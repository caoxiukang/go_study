/*
@Time : 2020/8/16 下午5:14
@Author : xiukang
@File : service
@Software: GoLand
*/
package jsonrpc

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type UserService struct {
}

type User struct {
	UserId int
	Name   string
}

func (UserService) GetUser(userId int, user *User) error {
	user.UserId = userId
	// this is your get user info service
	user.Name = "cao"
	fmt.Print("this is in service\n")
	return nil
}

func StartJsonRpcService() {
	rpc.Register(UserService{})
	listener, err := net.Listen("tcp", ":8080")
	fmt.Print("jsonrpc service start at 0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}
