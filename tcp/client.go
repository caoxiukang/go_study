/*
@Time : 2020/5/3 下午8:23
@Author : xiukang
@File : client
@Software: GoLand
*/
package tcp

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

var ClientAddr *net.TCPAddr

func Client() {
	ClientAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	conn, err := net.DialTCP("tcp", nil, ClientAddr)
	if err != nil {
		fmt.Println("Client connect error ! " + err.Error())
		return
	}

	fmt.Printf("Client connect ok\n")
	onMessageReceived(conn)
	defer conn.Close()
}
func onMessageReceived(conn *net.TCPConn) {
	// 获取来自服务器的信息
	readerFromServer := bufio.NewReader(conn)
	// 获取来自控制台的输入
	readerFromConsole := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("请输入：")
		sendMessage, _ := readerFromConsole.ReadString('\n')
		sendMessage = strings.Trim(sendMessage, "\r\n")

		if sendMessage == "q!" {
			break
		}
		_, err := conn.Write([]byte(sendMessage + "\n"))
		if err != nil {
			fmt.Println(err)
			break
		}

		msg, err := readerFromServer.ReadString('\n')
		fmt.Println("服务器：" + msg + "\n")
		if err != nil || err == io.EOF {
			fmt.Println(err)
			break
		}
	}
}
