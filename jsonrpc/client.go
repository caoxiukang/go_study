/*
@Time : 2020/8/16 下午5:14
@Author : xiukang
@File : client
@Software: GoLand
*/
package jsonrpc

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"
)

var Client *rpc.Client

func start() {
	time.Sleep(time.Second * 2)
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Print("start jsonrpc client err")
		return;
	}
	Client = jsonrpc.NewClient(conn)
	fmt.Print("jsonrpc client start at 0.0.0.0:8080")
}

// @Title 向servic发起请求
// @Param serviceName string 服务名字
// @Param args interface 参数
// @Param reply interface
// @return list []Catalog

func GetJsonService(serviceName string, args interface{}, reply interface{}) {

	err := Client.Call(serviceName, args, reply)
	if err != nil {
		fmt.Print("GetService err")
	}
}

// explames
func StartJsonRpcClient() {
	user := User{}
	GetJsonService("UserService.GetUser", 1, &user)
	fmt.Print(user)
}
