package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	trimmed := strings.TrimSpace(s)

	lastWord := ""
	for i := len(trimmed) - 1; i >= 0; i-- {
		if trimmed[i] == ' ' {
			if lastWord != "" {
				break
			}
		} else {
			lastWord = string(trimmed[i]) + lastWord
		}
	}

	return len(lastWord)
}

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
	fmt.Println(lengthOfLastWord("luffy is still joyboy"))
	fmt.Println(lengthOfLastWord(""))
}
