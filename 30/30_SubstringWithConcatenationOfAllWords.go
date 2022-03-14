package main

import (
	"fmt"
	"sort"
)

func findSubstring(s string, words []string) []int {
	sort.Strings(words)
	path := make(map[string]int)
	for i := 0; i < len(words); i++ {
		path[words[i]]++
	}
	result := make([]int, 0, len(words))

	for i := 0; i < len(words); i++ {
		if i > 0 && words[i] == words[i-1] {
			continue
		}
		temp := s
		totalIndex := 0
		for len(temp) >= len(words)*len(words[0]) {
			index := strStr(temp, words[i])
			if index >= 0 {
				path[words[i]]--
				if dfs(temp[index+len(words[0]):], words, path) {
					result = append(result, totalIndex+index)
				}
				path[words[i]]++
				temp = temp[index+1:]
				totalIndex += index + 1
			} else {
				break
			}
		}
	}
	// 字符串匹配
	return result

}

func dfs(s string, words []string, path map[string]int) bool {
	//for key, value := range path {
	//	for i := 0; i < len(words); i++ {
	//		//if contains(path, i) {
	//		//	continue
	//		//}
	//		if s[0:len(words[0])] == words[i] {
	//			//path = append(path, i)
	//			if dfs(s[len(words[0]):], words, path) {
	//				return true
	//			} else {
	//				//path = path[:len(path)-1]
	//			}
	//		}
	//	}
	//}
	//if len(path) >= len(words) {
	//	return true
	//}
	if len(s) < (len(words)-len(path))*len(words[0]) {
		return false
	}
	for i := 0; i < len(words); i++ {
		//if contains(path, i) {
		//	continue
		//}
		if s[0:len(words[0])] == words[i] {
			//path = append(path, i)
			if dfs(s[len(words[0]):], words, path) {
				return true
			} else {
				//path = path[:len(path)-1]
			}
		}
	}
	return false

}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

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

func main() {
	fmt.Println(findSubstring("aaaaaaaaaaaaaaaaaaaaaaaaaaaa", []string{"a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a"}))

	//fmt.Println(findSubstring("abcabcabc", []string{"abc", "abc", "abc"}))
	//fmt.Println(findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
	//fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}))
	//fmt.Println(findSubstring("barfoofoobarthefoobarman", []string{"bar", "foo", "the"}))
	//fmt.Println(findSubstring("lingmindraboofooowingdingbarrwingmonkeypoundcake", []string{"fooo", "barr", "wing", "ding", "wing"}))
	//fmt.Println("abc"[1:])
	//fmt.Println("abc"[0:])
	//fmt.Println("abc"[0:2])
	//fmt.Println("abc"[0:2] == "ab")
	//path := make([]int, 0, 10)
	//fmt.Println(path)
	//fmt.Println(len(path))
	//path = append(path, 10)
	//path = append(path, 10)
	//fmt.Println(path)
	//path = path[:1]
	//fmt.Println(path)

}

/**
Example 1:

Input: s = "barfoothefoobarman", words = ["foo", "bar"]
Output: [0, 9]
Explanation: Substrings starting at index 0 and 9 are "barfoo" and "foobar" respectively.
The output order does not matter, returning [9,0] is fine too.
Example 2:

Input: s = "wordgoodgoodgoodbestword", words = ["word","good", "best", "word"]
Output: []
Example 3:

Input: s = "barfoofoobarthefoobarman", words = ["bar", "foo", "the"]
Output: [6, 9, 12]
*/
