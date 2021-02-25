/*

给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。
在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0) 。
出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。


@Time : 2021/2/23 下午11:09
@Author : xiukang
@File : main
@Software: GoLand
*/
package main

import "fmt"

func maxArea(height []int) (area int) {
	var l int = 0
	var r int = len(height) - 1

	a := 0
	for l < r {
		if height[l]>=height[r]{
			a = (r-l) * height[r]
		}else{
			a = (r-l) * height[l]
		}

		if a >= area {
			area = a
		}
		if height[l] <= height[r] {
			l++
		} else{
			r--
		}
	}

	return area

}

func main() {
	x := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	y := maxArea(x)
	fmt.Println(y)
}
