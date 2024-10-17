package main

import "fmt"

func StrStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	haystackLen := len(haystack)
	needleLen := len(needle)

	if needleLen > haystackLen {
		return -1
	}

	for i := 0; i <= haystackLen-needleLen; i++ {

		if haystack[i:i+needleLen] == needle {
			return i
		}
	}

	return -1
}

func main() {
	haystack := "hello"
	needle := "ll"
	fmt.Println(haystack[:2+len(needle)])
}
