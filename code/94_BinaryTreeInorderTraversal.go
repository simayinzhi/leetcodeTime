package code

func Preorder(root *TreeNode) (res []int) {
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil{
			return
		}
		res = append(res, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return
}


func Preorder2(root *TreeNode) []int {
	if root == nil{
		return []int{}
	}
	var result []int
	result = append(result, root.Val)
	result = append(result, Preorder2(root.Left)...)
	result = append(result, Preorder2(root.Right)...)
	return result
}
func Inorder(root *TreeNode) (res []int) {
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil{
			return
		}
		preorder(node.Left)
		res = append(res, node.Val)
		preorder(node.Right)
	}
	preorder(root)
	return
}

func Inorder2(root *TreeNode) []int {
	if root == nil{
		return []int{}
	}
	var result []int
	result = append(result, Inorder2(root.Left)...)
	result = append(result, root.Val)
	result = append(result, Inorder2(root.Right)...)
	return result
}

func Postorder(root *TreeNode) (res []int) {
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil{
			return
		}
		preorder(node.Left)
		preorder(node.Right)
		res = append(res, node.Val)
	}
	preorder(root)
	return
}

func Postorder2(root *TreeNode) []int {
	if root == nil{
		return []int{}
	}
	var result []int
	result = append(result, Postorder2(root.Left)...)
	result = append(result, Postorder2(root.Right)...)
	result = append(result, root.Val)
	return result
}

func Inorder3(root *TreeNode) []int {
	var result []int
	var stack []*TreeNode
	temp := root
	for temp != nil || len(stack) > 0{
		for temp != nil {
			stack = append(stack, temp)
			temp = temp.Left
		}
		temp = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, temp.Val)
		temp = temp.Right
	}
	return result
}
func Preorder3(root *TreeNode) []int {
	var result []int
	var stack []*TreeNode
	temp := root
	for temp != nil || len(stack) > 0{
		for temp != nil {
			result = append(result, temp.Val)
			stack = append(stack, temp)
			temp = temp.Left
		}
		temp = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		temp = temp.Right
	}
	return result
}
func Postorder3(root *TreeNode) []int {
	var result []int
	var stack []*TreeNode
	stack = append(stack, root)
	for root != nil && len(stack) > 0{
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append([]int{node.Val}, result...)
		if node.Left != nil{
			stack = append(stack, node.Left)
		}
		if node.Right != nil{
			stack = append(stack, node.Right)
		}
	}
	return result
}
