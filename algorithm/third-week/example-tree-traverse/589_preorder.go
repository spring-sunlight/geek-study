package example_tree_traverse

import "geek-study/algorithm/common"

type Node struct {
	Val      int
	Children []*Node
}

//递归
func preorder(root *Node) []int {
	ans := []int{}

	var dfs func(root *Node)
	dfs = func(root *Node) {
		if root == nil {
			return
		}
		ans = append(ans, root.Val)
		for _, child := range root.Children {
			dfs(child)
		}
	}
	dfs(root)
	return ans
}

//迭代
func preorder1(root *Node) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	stack := common.NewStack()
	stack.Push(root)
	for !stack.IsEmpty() {
		top := stack.Top()
		ans = append(ans, top.(*Node).Val)
		stack.Pop()
		for i := len(top.(*Node).Children) - 1; i >= 0; i-- {
			stack.Push(top.(*Node).Children[i])
		}
	}
	return ans
}
