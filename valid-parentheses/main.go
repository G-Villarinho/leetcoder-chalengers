package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	items []rune
}

func (s *Stack) Push(item rune) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (rune, error) {
	if len(s.items) == 0 {
		return 0, errors.New("a pilha está vazia")
	}
	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item, nil
}

func (s *Stack) Top() (rune, error) {
	if len(s.items) == 0 {
		return 0, errors.New("a pilha está vazia")
	}
	return s.items[len(s.items)-1], nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}

func isValid(s string) bool {
	p := Stack{}

	for _, char := range s {
		switch char {
		case '(', '[', '{':
			p.Push(char)

		case ')':
			if p.IsEmpty() {
				return false
			}
			top, _ := p.Top()
			if top != '(' {
				return false
			}
			p.Pop()
		case ']':
			if p.IsEmpty() {
				return false
			}
			top, _ := p.Top()
			if top != '[' {
				return false
			}
			p.Pop()
		case '}':
			if p.IsEmpty() {
				return false
			}
			top, _ := p.Top()
			if top != '{' {
				return false
			}
			p.Pop()
		}
	}

	return p.IsEmpty()
}

func main() {
	s := "]"
	fmt.Println(isValid(s))
}
