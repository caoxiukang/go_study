/*
@Time : 2020/11/21 上午12:06
@Author : xiukang
@File : quickFind
@Software: GoLand
*/
package main

import "fmt"

// 1.先将数组排序，从两边开始找数据，，如果找的数据小于数组最中间的数，那就去数组前半段找，如果大于中间的数字，就去后半段查找， 然后迭代

func QuickFind(arr [6]int, leftIndex int, rightIndex int, findVal int) {
	middle := (rightIndex + leftIndex) / 2

	if rightIndex < leftIndex {
		fmt.Println("没找到")
		return
	}

	// 数组中间的数大于要查找的数据，则需要在左侧查找
	if arr[middle] > findVal {
		QuickFind(arr, leftIndex, middle-1, findVal)
	}

	// 数组中间的数小于要查找的数据，则需要在右侧查找
	if arr[middle] < findVal {
		QuickFind(arr, middle+1, rightIndex, findVal)
	}

	if arr[middle] == findVal {
		fmt.Printf("要找的数据在第%d位置", middle)
		return
	}
}

func main() {
	list := [6]int{2, 3, 1, 0, 1, 5}
	QuickFind(list, 0, 5, 5)
}
