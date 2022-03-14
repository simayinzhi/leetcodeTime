package code

func DeleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil{
		return root
	}
	if key == root.Val{
		if root.Right != nil{
			root.Val= successor(root)
			root.Right = DeleteNode(root.Right, root.Val)
		}else {
			if root.Left != nil{
				root.Val= predecessor(root)
				root.Left = DeleteNode(root.Left, root.Val)
			}else {
				return nil
			}
		}
	}else if key > root.Val && root.Right != nil{
		root.Right = DeleteNode(root.Right, key)
	}else if key < root.Val && root.Left != nil{
		root.Left = DeleteNode(root.Left, key)
	}
	return root
}

func predecessor(node *TreeNode) int{
	target := node.Left
	for target.Right != nil{
		target = target.Right
	}
	return target.Val
}

func successor(node *TreeNode) int{
	target := node.Right
	for target.Left != nil{
		target = target.Left
	}
	return target.Val
}


func FindNodeInBST(root TreeNode, key int) *TreeNode{
	if key == root.Val{
		return &root
	}else if key > root.Val && root.Right != nil{
		return FindNodeInBST(*root.Right, key)
	}else if key < root.Val && root.Left != nil{
		return FindNodeInBST(*root.Left, key)
	}
	return nil
}


