package practise

/*
	题目大意:给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' ，找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充

	示例 1：

	输入：board = [
			['X','O','X','X'],
			['X','O','O','X'],
			['X','X','O','X'],
			['X','O','X','X']]

	输出：[	['X','X','X','X'],
			['X','X','X','X'],
			['X','X','X','X'],
			['X','O','X','X']]
	解释：被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O' 最终都
	会被填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。

	思路: 如果一个区域没有被围绕,则这个区域必定和边界为 O 的值相连,从边界为 O 的值开始,把能到达的点都放到标记数组中
*/

func solve(board [][]byte) {
	rows := len(board)
	cols := len(board[0])
	flag := make([]int, rows*cols)

	dx := []int{-1, 0, 0, 1}
	dy := []int{0, -1, 1, 0}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= rows || j < 0 || j >= cols {
			return
		}
		if board[i][j] == 'X' {
			return
		}

		if flag[i*cols+j] == 1 {
			return
		}
		flag[i*cols+j] = 1
		for k := 0; k < 4; k++ {
			ni := dx[k] + i
			nj := dy[k] + j
			dfs(ni, nj)
		}

	}

	for i := 0; i < rows; i++ {
		if board[i][0] == 'O' {
			dfs(i, 0)
		}
		if board[i][cols-1] == 'O' {
			dfs(i, cols-1)
		}
	}

	for i := 0; i < cols; i++ {
		if board[0][i] == 'O' {
			dfs(0, i)
		}
		if board[rows-1][i] == 'O' {
			dfs(rows-1, i)
		}
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if flag[i*cols+j] == 0 {
				board[i][j] = 'X'
			} else {
				board[i][j] = 'O'
			}
		}
		//fmt.Println(board[i])
	}

}

func main() {
	board := [][]byte{
		{'X', 'O', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	solve(board)
}
