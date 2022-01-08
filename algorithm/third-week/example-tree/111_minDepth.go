package example_tree

import (
	"fmt"
	"geek-study/algorithm/common"
)

//深度优先搜索
func minDepthByDFS(root *common.TreeNode) int {
	return getMinDepthByDFS(root)
}

func getMinDepthByDFS(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil {
		return getMinDepthByDFS(root.Right) + 1
	}
	if root.Right == nil {
		return getMinDepthByDFS(root.Left) + 1
	}
	return min(getMinDepthByDFS(root.Left), getMinDepthByDFS(root.Right)) + 1
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

//广度优先搜索
func minDepthByBFS(root *common.TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*common.TreeNode{root}
	ans := 1
	for len(queue) != 0 {
		ans++
		temp := []*common.TreeNode{}
		for i := 0; i < len(queue); i++ {
			if queue[i].Left == nil && queue[i].Right == nil {
				return ans
			}
			if queue[i].Left != nil {
				temp = append(temp, queue[i].Left)
			}
			if queue[i].Right != nil {
				temp = append(temp, queue[i].Right)
			}
		}
		queue = temp
	}
	return ans
}

func main() {
	tree := common.NewTree(1, 2, nil, 3, nil)
	common.PrintTree(tree)
	fmt.Println(minDepthByDFS(tree))
}
