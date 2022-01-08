package array

import "fmt"

/*
	题目:编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
	 	每行中的整数从左到右按升序排列。
	 	每行的第一个整数大于前一行的最后一个整数。

	示例 1：
	输入：matrix = [
			[1,3,5,7],
			[10,11,16,20],
			[23,30,34,60]
		], target = 3
	输出：true

	示例 2：
	输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
	输出：false

	思路:双指针,先找出哪行,再找出哪列
*/

func searchMatrix(matrix [][]int, target int) bool {
	rows := len(matrix)
	cols := len(matrix[0])
	row, col := 0, 0
	up, down, left, right := 0, rows-1, 0, cols-1
	for up <= down {
		if target < matrix[up][left] || target > matrix[down][right] {
			return false
		}
		if target > matrix[up][right] && target < matrix[down][left] {
			up++
			down--
		} else if target <= matrix[up][right] {
			row = up
			break
		} else if target >= matrix[down][left] {
			row = down
			break
		}
	}
	fmt.Println(row)
	for left <= right {
		if target < matrix[row][left] || target > matrix[row][right] {
			return false
		}
		if target > matrix[row][left] && target < matrix[row][right] {
			left++
			right--
		} else if target == matrix[row][left] {
			col = left
			fmt.Println(row, col)
			return true
		} else if target == matrix[row][right] {
			col = right
			fmt.Println(row, col)
			return true
		}
	}
	return false
}

func main() {
	matrix := [][]int{
		{1, 3, 5, 7},
		{9, 11, 13, 15},
		{17, 19, 21, 23},
		{25, 27, 29, 31},
	}

	matrix1 := [][]int{{1}}

	fmt.Println(searchMatrix(matrix, 1))
	fmt.Println(searchMatrix(matrix, 3))
	fmt.Println(searchMatrix(matrix, 7))
	fmt.Println(searchMatrix(matrix, 25))
	fmt.Println(searchMatrix(matrix, 31))
	fmt.Println(searchMatrix(matrix, -1))
	fmt.Println(searchMatrix(matrix, 8))
	fmt.Println(searchMatrix(matrix, 16))
	fmt.Println(searchMatrix(matrix, 24))
	fmt.Println(searchMatrix(matrix, 32))
	fmt.Println()
	fmt.Println(searchMatrix(matrix1, 1))

}
