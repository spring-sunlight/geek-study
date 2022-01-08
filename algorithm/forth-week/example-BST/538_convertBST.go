package example_BST

//根节点的值等于根节点的值加上右子树的所有和
//根节点的右子树的值同理
//根节点的左子树等于当前值加上父亲结点算出的和

func convertBST(root *TreeNode) *TreeNode {
	sum := 0

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Right)
		sum += root.Val
		root.Val = sum
		dfs(root.Left)
	}
	dfs(root)
	return root
}
