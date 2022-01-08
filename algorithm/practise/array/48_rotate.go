package array

import "fmt"

//题目:把二维数组顺时针旋转 90°
//思路:
//1.先将数组沿着左上-右下对角线对折
//2.再将数组按照竖直对称轴翻转

//例如:	1, 2, 3,4	       10,11,12,1
//     12,13,14,5    =>     9,16,13,2
//	   11,16,15,6           8,15,14,3
//	   10, 9, 8,7           7, 6,5, 4

//1.沿着左上-右下对角线对折
//		1, 2, 3,4	       1,12,11,10
//     12,13,14,5    =>    2,13,16, 9
//	   11,16,15,6          3,14,15, 8
//	   10, 9, 8,7          4, 5, 6, 7

//2.沿着竖直对称轴翻转对折
//	   1,12,13,10	       10,11,12,1
//     2,13,16, 9    =>     9,16,13,2
//	   3,14,15, 8           8,15,14,3
//	   4, 5, 6, 7           7, 6,5, 4

//由此可知:
//1.把二维数组旋转 180°: 先按照水平对称轴翻转,再按照竖直对称轴翻转
//2.把二维数组旋转 270°: 先按照左上-右下角对角线对折,再按照水平对称轴翻转

func rotate(matrix [][]int) {
	rows := len(matrix)
	cols := len(matrix[0])
	for i := 0; i < rows; i++ {
		for j := i; j < cols; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for i := 0; i < rows; i++ {
		for left, right := 0, cols-1; left < right; left, right = left+1, right-1 {
			matrix[i][left], matrix[i][right] = matrix[i][right], matrix[i][left]
		}
	}
	for i := 0; i < rows; i++ {
		fmt.Println(matrix[i])
	}
}

func main() {
	matrix := [][]int{{1, 2, 3, 4}, {12, 13, 14, 5}, {11, 16, 5, 6}, {10, 9, 8, 7}}
	rotate(matrix)
}
