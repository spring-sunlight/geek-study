package example_prefixSum

import "fmt"

type NumMatrix struct {
	matrix          [][]int
	matrixPrefixSum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	return NumMatrix{
		matrix:          matrix,
		matrixPrefixSum: matrixPrefixSum(matrix),
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	row1++
	row2++
	col1++
	col2++
	return this.matrixPrefixSum[row2][col2] - this.matrixPrefixSum[row1-1][col2] - this.matrixPrefixSum[row2][col1-1] + this.matrixPrefixSum[row1-1][col1-1]
}

func matrixPrefixSum(matrix [][]int) [][]int {
	m := len(matrix)
	n := len(matrix[0])
	ans := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		ans[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			ans[i][j] = ans[i-1][j] + ans[i][j-1] - ans[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	return ans
}

func testNumMatrix() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	numMatrix := Constructor(matrix)
	fmt.Println(numMatrix.matrixPrefixSum)
	fmt.Println(numMatrix.SumRegion(1, 1, 2, 2))
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
