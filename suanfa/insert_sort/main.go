/*
@Time : 2021/1/27 下午11:35
@Author : xiukang
@File : main.go
@Software: GoLand
*/
package main

import "fmt"

func InsertSort(arr []int) {
	length := len(arr)
	for j := 1; j < length; j++ {
		insertVal := arr[j]
		insertIndex := j - 1
		for insertIndex >= 0 && arr[insertIndex] < insertVal {

			arr[insertIndex+1] = arr[insertIndex]
			insertIndex--
		}
		if insertIndex+1 != j {
			arr[insertIndex+1] = insertVal
		}
	}
}
// 9 7 6  10   j=3   insertIndex=2    insertVal=10
// 9 7 6  6    j=3   insertIndex=1    insertVal=10
// 9 7 7  6    j=3   insertIndex=0
// 9 9 7  6    j=3   insertIndex=0
// 10 9 7 6    j=3


func main() {
	x := []int{4, 1, 33, 5000,4, 7, 8, 9, 50}
	InsertSort(x)
	fmt.Println(x)
}
