package code

func LevelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil{
		return result
	}

	var stack []*TreeNode
	var tempStack []*TreeNode
	stack = append(stack, root)
	var level int
	for len(stack)>0{
		result = append(result, []int{})
		for i:=0; i < len(stack) ; i++{
			temp := stack[i]
			result[level] = append(result[level], temp.Val)
			if temp.Left != nil{
				tempStack = append(tempStack, temp.Left)
			}
			if temp.Right != nil{
				tempStack = append(tempStack, temp.Right)
			}
		}
		stack = make([]*TreeNode, len(tempStack))
		copy(stack, tempStack)
		tempStack = tempStack[0:0]

		level ++
	}
	return result
}

func LevelOrderBottom(root *TreeNode) [][]int {
	var result [][]int
	if root == nil{
		return result
	}

	var stack []*TreeNode
	var tempStack []*TreeNode
	stack = append(stack, root)
	var level int
	for len(stack)>0{
		var levelOrder []int
		for i:=0; i < len(stack) ; i++{
			temp := stack[i]
			levelOrder = append(levelOrder, temp.Val)
			if temp.Left != nil{
				tempStack = append(tempStack, temp.Left)
			}
			if temp.Right != nil{
				tempStack = append(tempStack, temp.Right)
			}
		}
		result = append([][]int{levelOrder}, result...)
		stack = make([]*TreeNode, len(tempStack))
		copy(stack, tempStack)
		tempStack = tempStack[0:0]
		level ++
	}
	return result
}

func NaryPreorder(root *Node) []int {
	if root == nil{
		return []int{}
	}
	var result []int
	result = append(result, root.Val)
	for i:=0; i < len(root.Children); i++{
		result = append(result, NaryPreorder(root.Children[i])...)
	}
	return result
}

func NaryLevelOrder(root *Node) [][]int {
	var result [][]int
	if root == nil{
		return result
	}

	var stack []*Node
	var tempStack []*Node
	stack = append(stack, root)
	var level int
	for len(stack)>0{
		result = append(result, []int{})
		for i:=0; i < len(stack) ; i++{
			temp := stack[i]
			result[level] = append(result[level], temp.Val)
			tempStack = append(tempStack, temp.Children...)
		}
		stack = make([]*Node, len(tempStack))
		copy(stack, tempStack)
		tempStack = tempStack[0:0]
		level ++
	}
	return result
}


func ZigzagLevelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil{
		return result
	}

	var stack []*TreeNode
	var tempStack []*TreeNode
	stack = append(stack, root)
	var level int
	for len(stack)>0{
		var zigzagSlice []int

		for i:=0; i < len(stack) ; i++{
			temp := stack[i]
			if level % 2 == 0{
				zigzagSlice = append(zigzagSlice, temp.Val)
			}else {
				zigzagSlice = append([]int{temp.Val}, zigzagSlice...)
			}
			if temp.Left != nil{
				tempStack = append(tempStack, temp.Left)
			}
			if temp.Right != nil{
				tempStack = append(tempStack, temp.Right)
			}
		}
		result = append(result, zigzagSlice)
		//zigzagSlice = zigzagSlice[0:0]
		stack = make([]*TreeNode, len(tempStack))
		copy(stack, tempStack)
		tempStack = tempStack[0:0]
		level ++
	}
	return result
}
