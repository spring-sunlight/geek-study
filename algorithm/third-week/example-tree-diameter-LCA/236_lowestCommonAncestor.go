package example_tree_diameter_LCA

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var dfs func(root *TreeNode) *TreeNode
	dfs = func(root *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}
		if root.Val == p.Val || root.Val == q.Val {
			return root
		}
		leftTree := dfs(root.Left)
		rightTree := dfs(root.Right)
		if leftTree != nil && rightTree != nil {
			return root
		}
		if leftTree == nil {
			return rightTree
		}
		return leftTree
	}
	return dfs(root)
}
