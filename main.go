/*
@Time : 2020/11/20 下午9:30
@Author : xiukang
@File : double_sort
@Software: GoLand
*/
package main

import (
	"fmt"
)

//var ch chan int

var ch = make(chan int, 1)

func main() {

	//ch := make(chan int, 1)
	// 求1-20 的阶乘
	//for i := 1; i <= 10; i++ {
	//	ch <- i
	//	time.Sleep(time.Second)
	//	go dome()
	//}
	ch <- 1
	x := <-ch
	<-ch
	go name()
	fmt.Println(x)
}

func name() {
	ch <- 1
}
func dome() int {
	var n int
	n = <-ch
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	fmt.Printf("%v的阶乘=%v\n", n, res)
	return res
}
