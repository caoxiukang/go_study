/*
@Time : 2020/11/21 上午12:08
@Author : xiukang
@File : bubble_sort
@Software: GoLand
*/
package bubble_sort

import "fmt"

type Student struct {
	name1 string
}

func (s *Student) Name() string {
	return (*s).name1
}


func ArrSort(arr [6]int) [6]int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}

	}
	fmt.Print(arr)
	return arr
}
