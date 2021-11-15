package third_week

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	m := map[int]int{}
	for i, v := range inorder {
		m[v] = i
	}
	var build func(int, int) *TreeNode
	build = func(inorderLeft, inorderRight int) *TreeNode {
		if inorderLeft > inorderRight {
			return nil
		}
		val := postorder[len(postorder)-1]
		postorder = postorder[:len(postorder)-1]
		root := &TreeNode{Val: val}
		inorderRootIndex := m[val]
		root.Right = build(inorderRootIndex+1, inorderRight)
		root.Left = build(inorderLeft, inorderRootIndex-1)
		return root
	}
	return build(0, len(inorder)-1)
}
