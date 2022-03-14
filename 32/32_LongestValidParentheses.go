package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func longestValidParentheses(s string) int {
	if s == "" {
		return 0
	}
	beginIndex := strings.Index(s, "(")
	endIndex := strings.LastIndex(s, ")")
	if beginIndex < 0 || endIndex < 0 || beginIndex >= endIndex {
		return 0
	}
	s = s[beginIndex : endIndex+1]
	lastLeft := 0
	lastRight := 0
	var stack []int
	length := 0

	for i := 0; i < len(s); i++ {
		if s[i] == 40 {
			stack = append(stack, i)
		} else {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
				lastRight = i
			} else {
				currentLength := lastRight - lastLeft + 1
				if currentLength > length {
					length = currentLength
				}
				lastLeft = i + 1
			}
		}
		if i == len(s)-1 {
			l := -1
			r := 0
			for j := 0; j < len(stack); j++ {
				if lastRight > stack[j] {
					r = stack[j]
					currentLength := r - l
					if currentLength > length {
						length = currentLength
					}
					l = stack[j]
				} else {
					break
				}
			}
			currentLength := lastRight - l
			if currentLength > length {
				length = currentLength
			}
		}
	}
	return length
}

func generateLeftRight(n int) string {
	var s string
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		j := rd.Int() & 1
		if j%2 == 0 {
			s += "("
		} else {
			s += ")"
		}
	}
	return s
}

func main() {
	//s := generateLeftRight(8)
	//fmt.Println(s)
	//fmt.Println(longestValidParentheses(s))

	fmt.Println(longestValidParentheses("(())()(()(("))
	fmt.Println(longestValidParentheses("))()("))
	fmt.Println(longestValidParentheses(")))(("))
	fmt.Println(longestValidParentheses("(()"))
	fmt.Println(longestValidParentheses(")()())"))
	fmt.Println(longestValidParentheses(""))
}
