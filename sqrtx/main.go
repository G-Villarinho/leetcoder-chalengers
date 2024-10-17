package main

import (
	"fmt"
)

func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	left, right := 2, x/2

	for left <= right {
		mid := left + (right-left)/2
		square := mid * mid

		if square == x {
			return mid
		} else if square < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right
}

func main() {
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(8))
	fmt.Println(mySqrt(16))
	fmt.Println(mySqrt(0))
	fmt.Println(mySqrt(1))
}
