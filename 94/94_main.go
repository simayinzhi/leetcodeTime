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
	fmt.Println(code.Preorder(&root))
	fmt.Println(code.Preorder2(&root))
	fmt.Println(code.Preorder3(&root))
	fmt.Println(code.Inorder(&root))
	fmt.Println(code.Inorder2(&root))
	fmt.Println(code.Inorder3(&root))
	fmt.Println(code.Postorder(&root))
	fmt.Println(code.Postorder2(&root))
	fmt.Println(code.Postorder3(&root))
	slice1 := []int{1,2,3,4,5,6}
	fmt.Println(slice1[0+1:])
	fmt.Println(slice1[:5])
	fmt.Println(slice1[5+1:])



}
