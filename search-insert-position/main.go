package main

import "fmt"

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 5
	result := searchInsert(nums, target)
	fmt.Println(result)

	target = 2
	result = searchInsert(nums, target)
	fmt.Println(result)

	target = 7
	result = searchInsert(nums, target)
	fmt.Println(result)

	target = 0
	result = searchInsert(nums, target)
	fmt.Println(result)
}
