package main

import (
	"fmt"
)

func main() {
	multiplication()
}

// 99乘法表
func multiplication() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d |", i, j, j*i)
		}
		fmt.Println()
	}
}
