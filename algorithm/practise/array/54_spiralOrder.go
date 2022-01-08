package array

import "fmt"

/*
	题目大意:按照螺旋顺序输出矩阵

	例: 1,  2, 3, 4
        5,  6, 7, 8
        9, 10,11,12
        13,14,15,16
	输出:1,2,3,4,8,12,16,15,14,13,9,5,6,7,11,10

		1,2,3,4
        8,12,16,15
		14,13,9,5
  		6,7,11,10

	思路: 设定好数组上下左右边界,每一圈按照顺序打印,递归打印下一圈即可
*/

func spiralOrder(matrix [][]int) []int {
	rows := len(matrix)
	cols := len(matrix[0])
	ans := []int{}
	up, down, left, right := 0, rows-1, 0, cols-1
	var recur func(up, down, left, right int)
	recur = func(up, down, left, right int) {
		if up > down || left > right {
			return
		}
		for i := left; i <= right; i++ {
			ans = append(ans, matrix[up][i])
		}
		for i := up + 1; i < down; i++ {
			ans = append(ans, matrix[i][right])
		}
		for i := right; i >= left && up != down; i-- {
			ans = append(ans, matrix[down][i])
		}
		for i := down - 1; i > up && left != right; i-- {
			ans = append(ans, matrix[i][up])
		}
		recur(up+1, down-1, left+1, right-1)
	}
	recur(up, down, left, right)
	return ans
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{5, 6, 7},
		{9, 10, 11},
	}
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
	fmt.Println(spiralOrder(matrix))
}
