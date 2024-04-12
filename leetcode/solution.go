package leetcode

// 首字母大写，代表公开方法
func TwoSum(nums []int, target int) TwoSumStruct {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return TwoSumStruct{begin: p, end: i}
		}
		hashTable[x] = i
	}
	return TwoSumStruct{}
}

type TwoSumStruct struct {
	begin int
	end   int
}
