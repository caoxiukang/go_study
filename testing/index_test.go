/*
@Time : 2020/11/4 下午10:38
@Author : xiukang
@File : indexText
@Software: GoLand
*/
package testings

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	var x, y int
	x = 1
	y = 3
	fmt.Print(x + y)
	t.Log("this is test ")
}

func TestSe(t *testing.T)  {
	t.Log(" T 1 WE")
}
