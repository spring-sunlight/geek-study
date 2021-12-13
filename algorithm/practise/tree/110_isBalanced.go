package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftDepth := getDepth(root.Left)
	rightDepth := getDepth(root.Left)
	if leftDepth-rightDepth > 1 || rightDepth-leftDepth > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func getDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(getDepth(root.Left), getDepth(root.Right)) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
