package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	k := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}

func main() {
	nums := []int{1, 1, 2}
	newLength := removeDuplicates(nums)
	fmt.Println(newLength)
	fmt.Println(nums[:newLength])
}
