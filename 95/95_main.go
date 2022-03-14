package main

import (
	"computerBase/code"
	"fmt"
)

func main() {
	var nodes []*code.TreeNode
	nodes =code.GenerateTrees(1)
	fmt.Println(nodes)
	nodes =code.GenerateTrees(2)
	fmt.Println(nodes)
	nodes =code.GenerateTrees(3)
	fmt.Println(nodes)
	nodes =code.GenerateTrees(4)
	fmt.Println(nodes)
}
