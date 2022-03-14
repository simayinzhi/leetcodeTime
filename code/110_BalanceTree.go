package code

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func IsBalanced_1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftLength := lengthOfTree_1(root.Left)
	rightLength := lengthOfTree_1(root.Right)
	if leftLength - rightLength > 1 ||
		leftLength - rightLength < -1{
		return false
	}
	if IsBalanced_1(root.Left) && IsBalanced_1(root.Right){
		return true
	}
	return false
}

func lengthOfTree_1(root *TreeNode) int {
	if root == nil{
		return 0
	}

	return Max(lengthOfTree(root.Left), lengthOfTree(root.Right)) + 1
}

