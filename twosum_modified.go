// ./main.go
package main

import (
	"fmt"
	"sort"
)

func twoSum(nums []int, target int) []int {
	fmt.Println("Array:", nums)
	fmt.Println("Target:", target)
	sort.Ints(nums)
	fmt.Println("Sorted Array:", nums)
	arrlength := len(nums)
	m := 0
	for i := 0; i < arrlength; i++ {
		if nums[i] < target {
			fmt.Println("Entered here:", nums[i])
			m = m + 1
			//newarr[i] = nums[i]
		}
	}
	newarr := nums[:m]
	fmt.Println("Entered here:", newarr)
	arrlength1 := len(newarr)
	fmt.Println("New Array:", newarr)
	for i := 0; i < arrlength1; i++ {
		for j := i + 1; j < arrlength1; j++ {
			if newarr[i]+newarr[j] == target {
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
