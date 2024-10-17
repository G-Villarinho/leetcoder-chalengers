package main

import "fmt"

func romanToInt(s string) int {
	romanValues := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	sum := 0
	i := 0

	for i < len(s) {
		if i < len(s)-1 && romanValues[rune(s[i])] < romanValues[rune(s[i+1])] {
			sum += romanValues[rune(s[i+1])] - romanValues[rune(s[i])]
			i += 2
		} else {
			sum += romanValues[rune(s[i])]
			i++
		}
	}

	return sum
}

func main() {
	fmt.Println(romanToInt("1994"))
}
