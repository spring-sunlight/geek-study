package example_BST

//二叉树搜索树后继
//过程:	1.先查找当前结点在不在
//		2.如果存在并且有右子树,则从右子树一路向左到最下面,就是大于结点的最小的值
//		3.如果不存在或者没有右子树,则结果在查询的结点的过程中,记录一下大于节点的值中最小的点

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var ans *TreeNode
	for root != nil {
		if root.Val == p.Val {
			if root.Right != nil {
				root = root.Right
				for root.Left != nil {
					root = root.Left
				}
				ans = root
			}
			break
		}
		if root.Val > p.Val && (ans == nil || ans.Val > root.Val) {
			ans = root
		}
		if root.Val < p.Val {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	return ans
}
