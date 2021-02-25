/*
@Time : 2020/5/3 下午8:25
@Author : xiukang
@File : dup
@Software: GoLand
*/
package dup

import (
	"bufio"
	"fmt"
	"os"
)

func TestDup() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "end" {
			break
		}
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
}
