package array

/*
	题目大意:给定一个二维矩阵,把元素为 0 所在的行,列上的元素都置为 0,要求原地修改

	例子: 1,1,1
		 1,0,1
         1,1,1

    输出: 1,0,1
		 0,0,0
	     1,0,1

	思路:把 0 元素所在的行列记录下来,最后统一改变
*/

func setZeroes(matrix [][]int) {
	rowsMap := make(map[int]struct{})
	colsMap := make(map[int]struct{})
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				rowsMap[i] = struct{}{}
				colsMap[j] = struct{}{}
			}
		}
	}
	for row, _ := range rowsMap {
		matrix[row] = make([]int, len(matrix[0]))
	}
	for col, _ := range colsMap {
		for i := 0; i < len(matrix); i++ {
			matrix[i][col] = 0
		}
	}
}
