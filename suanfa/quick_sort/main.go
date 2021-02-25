/*
@Time : 2021/2/3 下午10:37
@Author : xiukang
@File : main
@Software: GoLand
*/
package main

import "fmt"

// 0 1 2 3 4 5 6 7
// 2 3 1 4 5 0 6 8


//       l   r
// 2 3 1 0 5 4 6 8
//         l
// 思想：
//1.每次以数组的中间为基准，然后从两端开始边往中间找数据
//2。找到比中间数小的，放左边，找到比中间数大的放右边
//3.
func QuickSort(arr []int, left int, right int) {
	l := left
	r := right
	middle := arr[(right+left)/2]
	// 从小到大排序
	for ; l < r; {
		// 在左边找到比middle大的
		for arr[l] < middle {
			l++
		}
		// 在右边找到比middle小的
		for arr[r] > middle {
			r--
		}

		if l >= r {
			break
		}

		arr[l], arr[r] = arr[r], arr[l]

		// 两边同时向中间搜索，
		// 左边叔
		if arr[l] == middle {
			r--
		}
		if arr[r] == middle {
			l++
		}

		if r == l {
			r--
			l++
		}

		if left < r {
			QuickSort(arr, left, r)
		}

		if right > l {
			QuickSort(arr, l, right)
		}
	}
}
func main() {
	list := []int{2, 3, 1, 4, 5, 0, 6, 8}
	fmt.Println(list)
	QuickSort(list, 0, len(list)-1)
	fmt.Println(list)
}
