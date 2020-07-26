package main

import "fmt"

// stack implementation approach

type Stack string

func (s *Stack) IsEmpty() bool {
	return len(s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str)
}
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", ""
	} else {
		index := len(*s) - 1
		elem := (*s)[index]
		*s = (*s)[:index]
		return "", elem
	}
}

func ValidParenth(s string) bool {
	if len(s) == 0 {
		return true
	}
	var st Stack
	braces := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	for i, ch := range s {
		if len(braces[ch]) > 0 {
			//st = append(st, s[i])
			st.Push(s[i])
		} else {

			if !st.IsEmpty() || braces[st.Pop()] != s[i] {
				return false
			}
		}
	}
	return len(st) == 0

}

func main() {

	fmt.Println(ValidParenth("()[]{}"))
	fmt.Println(ValidParenth("([)]"))
}
