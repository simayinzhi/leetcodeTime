package code

func NaryTreePostorder(root *Node) []int {
	var result []int
	var stack []*Node
	stack = append(stack, root)
	for root != nil && len(stack) > 0{
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append([]int{node.Val}, result...)
		stack = append(stack, node.Children...)
	}
	return result
}
