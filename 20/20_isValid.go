package main

import (
	"fmt"
)

// 左边，入栈
// 右边，出栈，看是否成对，成对则删掉栈顶元素，不成对则返回false
// 匹配完成后栈不为空，则返回false
func isValid(s string) bool {
	var stack []uint8
	for i := 0; i < len(s); i++ {
		if isLeftBracket(s[i]) {
			stack = append(stack, s[i])
		}
		if isRightBracket(s[i]) {
			if len(stack) == 0 {
				return false
			}
			if isPariBracket(stack[len(stack)-1], s[i]) {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}

func isRightBracket(s2 uint8) bool {
	if s2 == 41 ||
		s2 == 93 ||
		s2 == 125 {
		return true
	}
	return false
}
func isLeftBracket(s1 uint8) bool {
	if s1 == 40 ||
		s1 == 91 ||
		s1 == 123 {
		return true
	}
	return false
}

func isPariBracket(s1 uint8, s2 uint8) bool {
	if s1 == 40 && s2 == 41 ||
		s1 == 91 && s2 == 93 ||
		s1 == 123 && s2 == 125 {
		return true
	}
	return false
}

func main() {
	sl := "[}"
	fmt.Println(isPariBracket(sl[0], sl[1]))
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()()"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid(""))
	fmt.Println(isValid("{}[]()"))
	fmt.Println(isValid("[)]"))
}
