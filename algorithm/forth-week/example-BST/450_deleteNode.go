package example_BST

//删除二叉搜索树的结点
//过程:	1.搜索该结点
//		2.查到之后如果该节点左子树空,则用右子树代替该节点
//		3.如果右子树空,则用左子树代替该节点
//		4.如果都不空,用这个结点的后继结点代替,并且把后继结点删掉
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		next := root.Right
		for next.Left != nil {
			next = next.Left
		}
		root.Right = deleteNode(root.Right, next.Val)
		root.Val = next.Val
	}
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}
