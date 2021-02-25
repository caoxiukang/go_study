/*
@Time : 2020/5/3 下午8:45
@Author : xiukang
@File : dupFile
@Software: GoLand
*/
package dup

import (
	"bufio"
	"fmt"
	"os"
)

func DupFile() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	//for input.Scan() {
	//	fmt.Print("...........\n")
	//	fmt.Println(input.Text()) // scanner.Bytes()
	//}

	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
