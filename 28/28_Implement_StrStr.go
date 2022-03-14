package main

import "fmt"

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) == 0 {
		return -1
	}

	i := 1
	j := 1
	var next = make([]int, len(needle)+1)
	getNext(needle, next)

	for i <= len(haystack) && j <= len(needle) {
		if j == 0 || haystack[i-1] == needle[j-1] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j > len(needle) {
		return i - len(needle) - 1
	} else {
		return -1
	}
}

func getNext(str string, next []int) {
	next[1] = 0
	i := 1
	j := 0
	for i < len(str) {
		if j == 0 || str[i-1] == str[j-1] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}
	}
}
func getNext1(str string, next []int) {
	next[1] = 0
	i := 1
	j := 0
	for i < len(str) {
		if j == 0 || str[i-1] == str[j-1] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}
	}
}

func main() {
	//hay := "hello"
	//needle := "ll"
	//fmt.Println(strStr(hay, needle))
	//hay = "aaaaa"
	//needle = "bba"
	//fmt.Println(strStr(hay, needle))
	//hay = ""
	//needle = ""
	//fmt.Println(strStr(hay, needle))
	//hay = "a"
	//needle = ""
	//fmt.Println(strStr(hay, needle))
	hay := "mississippi"
	needle := "ababac"
	fmt.Println(strStr(hay, needle))
}
