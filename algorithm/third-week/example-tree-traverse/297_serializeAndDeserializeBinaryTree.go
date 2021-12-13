package example_tree_traverse

import (
	"strconv"
	"strings"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
//1,2,3,null, 4,5

//每个结点多遍历一层,加 nil
func (this *Codec) serialize(root *TreeNode) string {
	ans := []string{}

	var preorder func(root *TreeNode)
	preorder = func(root *TreeNode) {
		if root == nil {
			ans = append(ans, "nil")
			return
		}
		ans = append(ans, strconv.Itoa(root.Val))
		preorder(root.Left)
		preorder(root.Right)
	}
	preorder(root)
	return strings.Join(ans, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	str := strings.Split(data, ",")

	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		if str[0] == "nil" {
			str = str[1:]
			return nil
		}
		val, _ := strconv.Atoi(str[0])
		root := &TreeNode{Val: val}
		str = str[1:]
		root.Left = dfs()
		root.Right = dfs()
		return root
	}
	return dfs()
}
