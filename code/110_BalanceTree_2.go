package code

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func IsBalanced(root *TreeNode) bool {
	return lengthOfTree(root) > -1
}

func lengthOfTree(root *TreeNode) int {
	if root == nil{
		return 0
	}

	leftLength:= lengthOfTree(root.Left)
	if leftLength == -1{
		return -1
	}
	rightLength:= lengthOfTree(root.Right)
	if rightLength == -1 || Abs(leftLength - rightLength) > 1{
		return -1
	}

	return Max(leftLength, rightLength) + 1
}

