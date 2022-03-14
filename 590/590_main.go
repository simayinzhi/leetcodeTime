package main

import (
	"computerBase/code"
	"fmt"
)

func main() {
	root := code.Node{1,
		[]*code.Node{
			&code.Node{3,
				[]*code.Node{
					&code.Node{5,nil},
					&code.Node{6,nil}}},
			&code.Node{2,nil},
			&code.Node{4,nil}}}
	fmt.Println(code.NaryTreePostorder(&root))
}
