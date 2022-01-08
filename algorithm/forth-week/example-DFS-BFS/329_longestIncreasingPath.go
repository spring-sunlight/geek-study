package example_DFS_BFS

import "fmt"

/*
	题目大意: 给定一个 m x n 整数矩阵 matrix ，找出其中 最长递增路径 的长度。

	对于每个单元格，你可以往上，下，左，右四个方向移动。 你 不能 在 对角线 方向上移动或移动到 边界外（即不允许环绕）。

	示例 1：
	输入：matrix = [[9,-1,-1],
				   [6,-1,-1],
                   [2,1,-1]]
	输出：4
	解释：最长递增路径为 [1, 2, 6, 9]。

	思路:最长适合使用 DFS,每次把当前点的答案算出来后都记录一下
*/

func longestIncreasingPath(matrix [][]int) int {
	rows := len(matrix)
	cols := len(matrix[0])
	res := make([][]int, rows)
	ans := 0
	for i := 0; i < rows; i++ {
		res[i] = make([]int, cols)
	}
	dx := []int{-1, 0, 0, 1}
	dy := []int{0, -1, 1, 0}
	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		if res[x][y] != 0 {
			return res[x][y]
		}
		res[x][y] = 1
		for i := 0; i < 4; i++ {
			nx := x + dx[i]
			ny := y + dy[i]
			if !isValid(nx, ny, rows, cols) {
				continue
			}
			if matrix[nx][ny] <= matrix[x][y] {
				continue
			}
			res[x][y] = max(res[x][y], dfs(nx, ny)+1)
		}
		return res[x][y]
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {

			ans = max(ans, dfs(i, j))
		}
	}
	for i := 0; i < rows; i++ {
		fmt.Println(res[i])
	}
	return ans
}

//func isValid(x, y, rows, cols int) bool {
//	return x >= 0 && x < rows && y >= 0 && y < cols
//}
//
//func max(x, y int) int {
//	if x >= y {
//		return x
//	}
//	return y
//}

func main() {
	matrix := [][]int{
		{9, 9, 4},
		{6, 6, 8},
		{2, 1, 2},
	}

	matrix1 := [][]int{
		{3, 4, 5},
		{3, 2, 6},
		{2, 2, 1},
	}
	fmt.Println(longestIncreasingPath(matrix))
	fmt.Println(longestIncreasingPath(matrix1))
}
