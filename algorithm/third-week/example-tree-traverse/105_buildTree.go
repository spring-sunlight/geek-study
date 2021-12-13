package example_tree_traverse

//Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
//Output: [3,9,20,null,null,15,7]

func buildTree(preorder []int, inorder []int) *TreeNode {
	var dfs func(l1, r1, l2, r2 int) *TreeNode
	dfs = func(l1, r1, l2, r2 int) *TreeNode {
		if l1 > r1 {
			return nil
		}
		root := &TreeNode{Val: preorder[l1]}
		mid := 0
		for inorder[mid] != root.Val {
			mid++
		}
		root.Left = dfs(l1+1, l1+mid-l2, l2, mid-1)
		root.Right = dfs(l1+mid-l2+1, r1, mid+1, r2)
		return root
	}
	return dfs(0, len(preorder)-1, 0, len(inorder)-1)
}
