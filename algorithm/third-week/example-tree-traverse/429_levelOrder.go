package example_tree_traverse

func levelOrder(root *Node) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	q := []*Node{root}
	for len(q) != 0 {
		temp := []*Node{}
		val := []int{}
		for i := 0; i < len(q); i++ {
			temp = append(temp, q[i].Children...)
			val = append(val, q[i].Val)
		}
		ans = append(ans, val)
		q = temp
	}
	return ans
}
