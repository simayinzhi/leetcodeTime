package code

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func numTrees(n int) int {
	return len(GenerateTrees(n))
}

func GenerateTrees(n int) []*TreeNode {
	if n < 1 || n > 8 {
		return nil
	}
	return generateTreesOpt(1, n)
}

func generateTreesOpt(start int,end int) []*TreeNode{
	if start > end{
		return []*TreeNode{nil}
	}
	if start == end{
		nodes := make([]*TreeNode, 0, 1)
		return append(nodes, &TreeNode{start,nil, nil})
	}

	nodes := []*TreeNode{}
	for i:=start; i<= end;i++{
		leftNodes := generateTreesOpt(start,i-1)
		rightNodes := generateTreesOpt(i+1,end)
		for j:=0; j<len(leftNodes); j++ {
			for k:=0; k<len(rightNodes); k++ {
				nodes = append(nodes, &TreeNode{i,leftNodes[j],rightNodes[k]})
			}
		}
	}
	return nodes
}

