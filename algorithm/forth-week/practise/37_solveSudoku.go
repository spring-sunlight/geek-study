package practise

import "fmt"

// @solution-sync:begin
func solveSudoku(board [][]byte) {
	var rows, cols [9][9]bool
	var cube [3][3][9]bool
	var space [][2]int
	//首先把数组中没填的数字的位置记录下来
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '.' {
				space = append(space, [2]int{i, j})
			} else {
				rows[i][board[i][j]-'1'] = true
				cols[j][board[i][j]-'1'] = true
				cube[i/3][j/3][board[i][j]-'1'] = true
			}
		}
	}
	var dfs func(pos int) bool
	dfs = func(pos int) bool {
		if pos == len(space) {
			return true
		}
		//对于每个位置,使用 1-9 位置挨个试
		i, j := space[pos][0], space[pos][1]
		for b := 0; b < 9; b++ {
			digit := byte(b)
			if !rows[i][digit] && !cols[j][digit] && !cube[i/3][j/3][digit] {
				rows[i][digit] = true
				cols[j][digit] = true
				cube[i/3][j/3][digit] = true
				board[i][j] = digit + '1'
				if dfs(pos + 1) {
					return true
				}
				rows[i][digit] = false
				cols[j][digit] = false
				cube[i/3][j/3][digit] = false
			}

		}
		return false
	}
	dfs(0)
}

// @solution-sync:end

func main() {
	var board = [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	solveSudoku(board)
	for i := 0; i < len(board); i++ {
		fmt.Println(string(board[i]))
	}
}
