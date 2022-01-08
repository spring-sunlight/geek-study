package tree

func levelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	ans = append(ans, []int{root.Val})
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		childrens := []*TreeNode{}
		temp := []int{}
		for _, node := range queue {
			if node.Left != nil {
				childrens = append(childrens, node.Left)
				temp = append(temp, node.Left.Val)
			}
			if node.Right != nil {
				childrens = append(childrens, node.Right)
				temp = append(temp, node.Right.Val)
			}
		}
		if len(temp) != 0 {
			ans = append(ans, temp)
		}
		queue = childrens
	}
	return ans
}
