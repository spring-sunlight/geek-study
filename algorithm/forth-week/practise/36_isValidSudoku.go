package practise

import "fmt"

/*
	题目:
	请你判断一个 9 x 9 的数独是否有效。只需要 根据以下规则 ，验证已经填入的数字是否有效即可。
	数字 1-9 在每一行只能出现一次。
	数字 1-9 在每一列只能出现一次。
	数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）

	思路: 判断是否出现在一行,是否出现在一列,行列分别除以 3 之后是否相等,用 map 存储

*/

func isValidSudoku(board [][]byte) bool {
	rows := len(board)
	cols := len(board[0])
	rowMap := make(map[int]map[byte]struct{})
	colMap := make(map[int]map[byte]struct{})
	squareMap := make(map[[2]int]map[byte]struct{})
	for i := 0; i < rows; i++ {
		rowMap[i] = make(map[byte]struct{})
	}
	for i := 0; i < cols; i++ {
		colMap[i] = make(map[byte]struct{})
	}

	for i := 0; i < rows; i += 3 {
		for j := 0; j < cols; j += 3 {
			squareMap[[2]int{i / 3, j / 3}] = make(map[byte]struct{})
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] != '.' {
				if _, ok := rowMap[i][board[i][j]]; ok {
					//fmt.Println("err:", i, j)
					return false
				} else {
					rowMap[i][board[i][j]] = struct{}{}
				}
				if _, ok := colMap[j][board[i][j]]; ok {
					//fmt.Println("err:", i, j)
					return false
				} else {
					colMap[j][board[i][j]] = struct{}{}
				}
				if _, ok := squareMap[[2]int{i / 3, j / 3}][board[i][j]]; ok {
					//fmt.Println("err:", i, j)
					return false
				} else {
					squareMap[[2]int{i / 3, j / 3}][board[i][j]] = struct{}{}
				}
				//fmt.Println(i, j)
				//fmt.Println(rowMap)
				//fmt.Println(colMap)
				//fmt.Println(squareMap)
				//fmt.Println("======================")
			}
		}
	}
	return true
}

func main() {
	board := [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '1', '.', '.', '.', '.', '6', '.'},
		{'2', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '1', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(isValidSudoku(board))

}
