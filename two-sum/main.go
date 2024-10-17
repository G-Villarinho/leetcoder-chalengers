package main

import "fmt"

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if j, found := numMap[complement]; found {
			return []int{j, i}
		}

		numMap[num] = i
	}

	return nil
}

func main() {
	nums := []int{3, 2, 4}
	target := 6
	result := twoSum(nums, target)
	fmt.Println(result)
}
