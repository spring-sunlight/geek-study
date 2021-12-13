package tree

func rightSideView(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	} else {
		ans = append(ans, root.Val)
	}
	nodes := []*TreeNode{root}
	for len(nodes) != 0 {
		temp := []*TreeNode{}
		for i := 0; i < len(nodes); i++ {
			if nodes[i].Left != nil {
				temp = append(temp, nodes[i].Left)
			}
			if nodes[i].Right != nil {
				temp = append(temp, nodes[i].Right)
			}
		}
		if len(temp) != 0 {
			ans = append(ans, temp[len(temp)-1].Val)
		}
		nodes = temp
	}
	return ans
}
