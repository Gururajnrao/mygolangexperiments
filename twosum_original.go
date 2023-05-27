// ./main.go
package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	fmt.Println("Array:", nums)
	fmt.Println("Target:", target)
	arrlength := len(nums)
	for i := 0; i < arrlength; i++ {
		for j := i + 1; j < arrlength; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}

	}

	return []int{}

}

func main() {

	nums := []int{2, 7, 11, 15}
	target := 9
	val := twoSum(nums, target)
	fmt.Println("Result:", val)

}
