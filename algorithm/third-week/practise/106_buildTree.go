package practise

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//中序遍历 inorder = [9,3,15,20,7]
//后序遍历 postorder = [9,15,7,20,3]

func buildTree(inorder []int, postorder []int) *TreeNode {
	inorderIndexMap := make(map[int]int)
	for index, value := range inorder {
		inorderIndexMap[value] = index
	}
	var dfs func(l1, r1, l2, r2 int) *TreeNode
	dfs = func(l1, r1, l2, r2 int) *TreeNode {
		if l1 > r1 {
			return nil
		}
		fmt.Println(l1, r1, l2, r2)

		root := &TreeNode{Val: postorder[r2]}
		inorderRootIndex := inorderIndexMap[postorder[r2]]
		leftNum := inorderRootIndex - l1

		root.Left = dfs(l1, inorderRootIndex-1, l2, l2+leftNum-1)

		root.Right = dfs(inorderRootIndex+1, r1, l2+leftNum, r2-1)

		return root
	}
	return dfs(0, len(inorder)-1, 0, len(postorder)-1)
}

func buildTree1(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	rootIndex := 0
	for ; rootIndex < len(inorder); rootIndex++ {
		if inorder[rootIndex] == postorder[len(postorder)-1] {
			break
		}
	}
	root.Left = buildTree1(inorder[:rootIndex], postorder[:rootIndex])
	root.Right = buildTree1(inorder[rootIndex+1:], postorder[rootIndex:len(postorder)-1])
	return root

}
