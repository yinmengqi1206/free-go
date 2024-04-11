package leetcode

import (
	"fmt"
	"testing"
)

func TestLeetCode(t *testing.T) {
	fmt.Printf("两数之和: %v\n", twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}
