package code

func BuildTreeByPreorderInorder(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 || len(preorder) != len(inorder){
		return nil
	}
	node := TreeNode{preorder[0], nil, nil}
	// do something
	inorderRootPosition := indexOf(preorder[0], inorder)
	node.Left = BuildTreeByPreorderInorder(preorder[1:inorderRootPosition+1], inorder[:inorderRootPosition])
	node.Right = BuildTreeByPreorderInorder(preorder[inorderRootPosition+1:], inorder[inorderRootPosition+1:])
	return &node
}

func BuildTreeByPreorderInorderIterator(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0{
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	stack := []*TreeNode{}
	stack = append(stack, root)
	var inorderIndex int
	for i := 1; i < len(preorder); i++ {
		preorderVal := preorder[i]
		node := stack[len(stack)-1]
		if node.Val != inorder[inorderIndex]{
			node.Left = &TreeNode{preorderVal, nil, nil}
			stack = append(stack, node.Left)
		} else {
			for len(stack) != 0 && stack[len(stack)-1].Val == inorder[inorderIndex]{
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				inorderIndex++
			}
			node.Right = &TreeNode{preorderVal, nil, nil}
			stack = append(stack, node.Right)
		}
	}
	return root
}


func BuildTreeByInorderPostorder(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 || len(inorder) == 0 || len(postorder) != len(inorder){
		return nil
	}
	node := TreeNode{postorder[len(postorder)-1], nil, nil}
	// do something
	inorderRootPosition := indexOf(postorder[len(postorder)-1], inorder)
	node.Left = BuildTreeByInorderPostorder(inorder[:inorderRootPosition],postorder[:inorderRootPosition])
	node.Right = BuildTreeByInorderPostorder(inorder[inorderRootPosition+1:], postorder[inorderRootPosition:len(postorder)-1])
	return &node
}

func indexOf(element int, data []int) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1    //not found.
}

