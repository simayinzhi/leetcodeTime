package main

import "fmt"

var i = 0

func generateParenthesis(n int) []string {
	var res = make([]string, 0) // create result
	if n <= 0 {
		return res
	}

	var dfs func(string, int, int)                 // inner function
	dfs = func(path string, open int, close int) { // define inner func
		i++
		println(i)
		if len(path) == 2*n { // end condition, append path to result
			res = append(res, path)
			return
		}

		if open < n { // Can go left, go left
			dfs(path+"(", open+1, close)
		}
		if close < open { // Right < left, go right
			dfs(path+")", open, close+1)
		}
	}

	dfs("", 0, 0)
	return res
}

func main() {
	parenthesis := generateParenthesis(8)
	fmt.Println(parenthesis)
}
