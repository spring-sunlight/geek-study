package array

import "fmt"

/*
	题目大意: 给定一个 n,生成一个 n*n 的螺旋矩阵

	例子: n = 4

    输出:   1, 2, 3,4
          12,13,14,5
		  11,16,15,6
		  10, 9, 8,7

*/

func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}
	count := 1
	up, down, left, right := 0, n-1, 0, n-1
	var recur func(up, down, left, right int)
	recur = func(up, down, left, right int) {
		if up > down || left > right {
			return
		}
		if up == down && left == right {
			ans[up][left] = count
			return
		}
		for i := left; i <= right; i++ {
			ans[up][i] = count
			count++
		}

		for i := up + 1; i < down; i++ {
			ans[i][right] = count
			count++
		}
		for i := right; i >= left; i-- {
			ans[down][i] = count
			count++
		}

		for i := down - 1; i > up; i-- {
			ans[i][left] = count
			count++
		}
		recur(up+1, down-1, left+1, right-1)
	}
	recur(up, down, left, right)
	return ans
}

func main() {
	nums := generateMatrix(3)
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
}
