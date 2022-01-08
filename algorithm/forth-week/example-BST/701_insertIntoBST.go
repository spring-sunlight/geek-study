package example_BST

//比根节点大放在右子树,比根节点小放在左子树
type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Left: nil, Right: nil, Val: val}
	}
	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	}
	return root
}
