package main

import (
	"fmt"

)

//时间复杂度：O(n²)
//空间复杂度：O(1)
func twoSum(nums []int, target int) []int {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i] + nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

//时间复杂度：O(n)
//空间复杂度：O(n)
func twoSumMap(nums []int, target int) []int {
	length := len(nums)
	intM := make(map[int]int, length)
	for k, v := range nums {
		intM[v] = k
	}
	for i:= 0; i < length ; i++ {
		complement := target - nums[i]
		if n, ok := intM[complement]; ok && n != i {
			return []int{i, n}
		}
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{1, 2, 3, 4, 5, 6, 7, 8}, 2))
	fmt.Println(twoSumMap([]int{1, 2, 3, 4, 5, 6, 7, 8}, 2))
}
