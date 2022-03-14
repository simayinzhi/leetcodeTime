package code




func Rob(root *TreeNode) int {
	if root == nil{
		return 0
	}
	layer := 1
	blackRed := make([]int, 2)
	steal(root, blackRed, layer)
	if blackRed[0] > blackRed[1]{
		return blackRed[0]
	}else {
		return blackRed[1]
	}
}

func steal(node *TreeNode, blackRed []int, layer int) {
	blackRed[layer%2] += node.Val
	layer ++
	if node.Left != nil{
		steal(node.Left, blackRed, layer)
	}
	if node.Right != nil{
		steal(node.Right, blackRed, layer)
	}
}
