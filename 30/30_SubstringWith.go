package main

import "sort"

func findSubstring(s string, words []string) []int {
	sort.Strings(words)
	path := make([]int, 0, len(words))
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
				path = append(path, i)
				if dfs(temp[index+len(words[0]):], words, path) {
					result = append(result, totalIndex+index)
				}
				path = path[:len(path)-1]
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
func main() {

}
