package main

import (
	"fmt"
)

func twoSumBruteForceSolution(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSumHashTableSolution(nums []int, target int) []int {
	hash := make(map[int]int)
	for i, n := range nums {
		key := target - n
		// hash[key] will return there is the value in the existed map or not
		if index, ok := hash[key]; ok {
			return []int{index, i}
		}
		hash[n] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(nums[0:])
	target := 17
	output := twoSumHashTableSolution(nums, target)
	fmt.Println(output)
}
