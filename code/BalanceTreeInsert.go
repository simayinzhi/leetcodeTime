package code

func LeftRotate(root *BalanceTreeNode) *BalanceTreeNode{
	node := root.Right
	root.Right = node.Left
	node.Left = root
	root = node
	return root
}
func RightRotate(root *BalanceTreeNode) *BalanceTreeNode{
	node := root.Left
	root.Left = node.Right
	node.Right = root
	root = node
	return root
}
func LeftBalance(root *BalanceTreeNode) *BalanceTreeNode{
	if root.Left.Bf == 1{
		root.Bf = 0
		root.Left.Bf = 0
		root = RightRotate(root)
	}else if root.Left.Bf == -1{
		rightNode := root.Left.Right
		if rightNode.Bf == 1{
			root.Bf = -1
			root.Left.Bf = 0
		}else if rightNode.Bf == 0{
			root.Bf = 0
			root.Left.Bf = 0
		}else {
			root.Bf = 0
			root.Left.Bf = 1

		}
		rightNode.Bf = 0
		root.Left = LeftRotate(root.Left)
		root = RightRotate(root)
	}
	return root
}
func RightBalance(root *BalanceTreeNode) *BalanceTreeNode{
	if root.Right.Bf == -1{
		root.Bf = 0
		root.Right.Bf = 0
		root = LeftRotate(root)
	}else if root.Right.Bf == 1{
		rightNode := root.Right.Left
		if rightNode.Bf == 1{
			root.Bf = 0
			root.Right.Bf = -1
		}else if rightNode.Bf == 0{
			root.Bf = 0
			root.Right.Bf = 0
		}else {
			root.Bf = 1
			root.Right.Bf = 0
		}
		rightNode.Bf = 0
		root.Right = RightRotate(root.Right)
		root = LeftRotate(root)
	}
	return root
}

func InsertNode(root *BalanceTreeNode, key int, taller *bool) int{
	if root == nil{
		*root = BalanceTreeNode{key, 0, nil, nil}
		*taller = true
		return 1
	}
	if root.Val == key{
		return 0
	}else if root.Val > key{
		result := InsertNode(root.Left, key, taller)
		if result == 0{
			return 0
		}
		if *taller == true && root.Bf == 1{
			RightBalance(root)
		}
	}else {
		result := InsertNode(root.Right, key, taller)
		if result == 0 {
			return 0
		}
		if *taller == true && root.Bf == 1 {
			LeftBalance(root)
		}
	}
	return 1
}
