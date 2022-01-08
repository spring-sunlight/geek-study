package array

import "fmt"

/*
	题目:
	给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。
	单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

	 示例 1：
	输入：board = [
			['A','B','C','E'],
			['S','F','C','S'],
			['A','D','E','E']], word =
	'ABCCED'
	输出：true

	 示例 2：
	输入：board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word =
	'SEE'
	输出：true

	 示例 3：
	输入：board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word =
	'ABCB'
	输出：false

	思路:dfs
*/

func exist(board [][]byte, word string) bool {
	rows := len(board)
	cols := len(board[0])
	visited := make([]int, rows*cols)
	dx := []int{-1, 0, 0, 1}
	dy := []int{0, -1, 1, 0}
	ans := false
	var dfs func(i, j, index int)
	dfs = func(i, j, index int) {
		if i < 0 || i >= rows || j < 0 || j >= cols || index >= len(word) {
			return
		}
		if visited[i*cols+j] == 1 {
			return
		}

		if board[i][j] != word[index] {
			return
		}

		if index == len(word)-1 {
			ans = true
			return
		}
		//左下上右
		for k := 0; k < 4; k++ {
			ni := dx[k] + i
			nj := dy[k] + j
			visited[i*cols+j] = 1
			dfs(ni, nj, index+1)
			visited[i*cols+j] = 0
		}

	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == word[0] {
				visited = make([]int, rows*cols)
				dfs(i, j, 0)
				if ans {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'E', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	board1 := [][]byte{
		{'A', 'B'},
		{'C', 'D'},
	}

	fmt.Println(exist(board, "ABCESEEEF"))
	fmt.Println(exist(board1, "ACDB"))
}
