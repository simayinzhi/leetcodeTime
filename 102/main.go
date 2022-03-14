package main

import (
	"computerBase/code"
	"fmt"
)

func main() {
	root := code.TreeNode{1,
		&code.TreeNode{2,
			&code.TreeNode{4,
				nil,&code.TreeNode{6,
					&code.TreeNode{7,nil,nil},
				&code.TreeNode{8,nil,nil}}},nil},
		&code.TreeNode{3,nil,&code.TreeNode{5,nil,nil}}}
	fmt.Println(code.LevelOrder(&root))
	fmt.Println(code.ZigzagLevelOrder(&root))
	fmt.Println(code.LevelOrderBottom(&root))

	//tree := code.BuildTreeByPreorderInorder([]int{1,2,4,5,3,6,7},[]int{4,2,5,1,6,3,7})
	tree := code.BuildTreeByPreorderInorderIterator([]int{1,2,4,5,3,6,7},[]int{4,2,5,1,6,3,7})
	//tree := code.BuildTreeByInorderPostorder([]int{4,2,5,1,6,3,7},[]int{4,5,2,6,7,3,1})
	//tree := code.BuildTreeByPreorderInorder([]int{1,2,3},[]int{3,2,1})
	fmt.Println(tree)
	fmt.Println(code.LevelOrder(tree))
	fmt.Println(code.ZigzagLevelOrder(tree))
	fmt.Println(code.LevelOrderBottom(tree))
}
