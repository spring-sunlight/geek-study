package tree

//二叉树锯齿形遍历
//BFS即可

func zigzagLevelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	queue := []*TreeNode{root}
	ans = append(ans, []int{queue[0].Val})
	depth := 0
	for len(queue) != 0 {
		depth++
		temp := []*TreeNode{}
		value := []int{}
		for _, q := range queue {
			if q.Left != nil {
				temp = append(temp, q.Left)
			}
			if q.Right != nil {
				temp = append(temp, q.Right)
			}
		}
		if depth%2 == 1 {
			for i := len(temp) - 1; i >= 0; i-- {
				value = append(value, temp[i].Val)
			}
		} else {
			for i := 0; i < len(temp); i++ {
				value = append(value, temp[i].Val)
			}
		}
		if len(value) != 0 {
			ans = append(ans, value)
		}
		queue = temp
	}
	return ans
}
