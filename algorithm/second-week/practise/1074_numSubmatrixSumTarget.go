package practise

import "fmt"

/*
	输入：matrix = [[0,1,0],[1,1,1],[0,1,0]], target = 0
	输出：4
	解释：四个只含 0 的 1x1 子矩阵。

	输入：matrix = [[1,-1],[-1,1]], target = 0
	输出：5
	解释：两个 1x2 子矩阵，加上两个 2x1 子矩阵，再加上一个 2x2 子矩阵。

	思路:二维数组前缀和,暴力求解

*/

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	ans := 0
	prefixSum := make([][]int, len(matrix)+1)
	for i := 0; i < len(prefixSum); i++ {
		prefixSum[i] = make([]int, len(matrix[0])+1)
	}
	for i := 1; i < len(prefixSum); i++ {
		for j := 1; j < len(prefixSum[0]); j++ {
			prefixSum[i][j] = prefixSum[i-1][j] + prefixSum[i][j-1] - prefixSum[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	fmt.Println(prefixSum)
	for row1 := 1; row1 < len(prefixSum); row1++ {
		for col1 := 1; col1 < len(prefixSum[0]); col1++ {
			for row2 := row1; row2 < len(prefixSum); row2++ {
				for col2 := col1; col2 < len(prefixSum[0]); col2++ {
					sum := prefixSum[row2][col2] - prefixSum[row2][col1-1] - prefixSum[row1-1][col2] + prefixSum[row1-1][col1-1]
					//fmt.Println(row1, col1, row2, col2, sum)
					//fmt.Println("prefixSum[row2][col2]:", prefixSum[row2][col2])
					//fmt.Println("prefixSum[row1-1][col1]:", prefixSum[row1-1][col1])
					//fmt.Println("prefixSum[row1][col1-1]:", prefixSum[row1][col1-1])
					//fmt.Println("prefixSum[row1-1][col1-1]:", prefixSum[row1-1][col1-1])
					//fmt.Println()
					if sum == target {
						fmt.Println(row1, col1, row2, col2, sum)
						ans++
					}
				}
			}
		}
	}
	return ans
}

func main() {
	matrix := [][]int{
		{0, 1, 1, 1, 0, 1},
		{0, 0, 0, 0, 0, 1},
		{0, 0, 1, 0, 0, 1},
		{1, 1, 0, 1, 1, 0},
		{1, 0, 0, 1, 0, 0},
	}
	fmt.Println(numSubmatrixSumTarget(matrix, 0))
}
