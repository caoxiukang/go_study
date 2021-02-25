/*
title  从大数组中找到de和等于100的
@Time : 2020/11/2 下午11:29
@Author : xiukang
@File : sum=1-100
@Software: GoLand
*/
package sum_1_100

import (
	"fmt"
	"sort"
)

func GetSumMum(arr []int, sum int) (x int, y int) {
	var i, j int
	var n int
	n = len(arr)
	sort.Ints(arr)
	i = 0
	j = n - i - 1
	for i = 0; i < j; j = n - i {
		if arr[i]+arr[j] == sum {
			fmt.Print(arr[i], arr[j])
		}
		i++
		j--
	}
	return 0, 0
}
