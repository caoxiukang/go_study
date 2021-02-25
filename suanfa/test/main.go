/*
@Time : 2020/11/20 下午9:30
@Author : xiukang
@File : double_sort
@Software: GoLand
*/
package main

import (
	"encoding/json"
	"fmt"
	"study/suanfa/bubble_sort"
)

type inster int

type B struct {
	bubble_sort.Student
}

func main() {
	var mapx map[string]string
	var x []map[string]string

	mapx = make(map[string]string)
	mapx["1"] = "1"
	mapx["2"] = "2"

	x = append(x, mapx)
	mapx["2"] = "4"
	x = append(x, mapx)
	str, _ := json.Marshal(x)
	v := string(str)
	//fmt.Println(string(str))

	json.Unmarshal([]byte(v),&x)
	fmt.Println(x)
}

func (i *inster) add() {
	fmt.Print(*i)
}
