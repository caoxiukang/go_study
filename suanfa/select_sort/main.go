/*
@Time : 2021/1/27 下午11:35
@Author : xiukang
@File : main.go
@Software: GoLand
*/
package main

import "fmt"

func SelectSort(arr []int) {
	maxIndex := 0
	length := len(arr)
	for j := 0; j < length-1; j++ {
		maxNum := (arr)[j]
		maxIndex = j
		for i := j + 1; i < length; i++ {
			if maxNum < (arr)[i] {
				maxNum = (arr)[i]
				maxIndex = i
			}
		}

		if maxIndex != j {
			(arr)[j], (arr)[maxIndex] = (arr)[maxIndex], (arr)[j]
		}
	}
}

func main() {
	x := []int{4, 1, 33, 4, 7, 8, 9, 0}
	SelectSort(x)
	fmt.Println(x)
}
