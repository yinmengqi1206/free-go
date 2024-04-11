package main

import (
	"fmt"
	"testing"
)

// go语言fmt.Printf格式输出
//https://juejin.cn/post/7127305094273433613

func TestPrint(t *testing.T) {
	multiplication()
	// var ids = []int{1, 9, 2, 4, 3, 7} //定义一个数组有5个int元素
	// count := bubbleSort(ids)
	// fmt.Println(count)
	//fmt.Println(expRecur(3))
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'F'; j++ {
			if j == 'B' {
				goto OVER
			}
			fmt.Printf("%v_%c \n", i, j)
		}
	}
OVER:
	fmt.Println("over")
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

/* 平方阶（冒泡排序） */
func bubbleSort(nums []int) int {
	count := 0 // 计数器
	// 外循环：待排序元素数量为 n-1, n-2, ..., 1
	for i := len(nums) - 1; i > 0; i-- {
		// 内循环：冒泡操作
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				// 交换 nums[j] 与 nums[j + 1]
				tmp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = tmp
				count += 3 // 元素交换包含 3 个单元操作
			}
		}
	}
	return count
}

// 指数阶
func exponential(n int) int {
	count, base := 0, 1
	for i := 0; i < n; i++ {
		for j := 0; j < base; j++ {
			count++
		}
		base *= 2
	}
	return count
}

/* 指数阶（递归实现）*/
func expRecur(n int) int {
	if n == 1 {
		return 1
	}
	return expRecur(n-1) + expRecur(n-1) + 1
}
