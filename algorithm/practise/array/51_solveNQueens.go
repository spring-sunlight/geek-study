package array

import "fmt"

/*
	题目大意:N皇后的解法
	思路:类似全排列问题,多了一些条件,深度优先搜索,从第 0 行出发,查看每列是否符合条件,
		对角线上的值 row+col 相等或者 row-col 相等,所以判断三个条件是否被访问
		row+col 和 row-col 可能会越界,因此不能用数组存,用 map 存

*/

func solveNQueens(n int) [][]string {
	ans := [][]string{}
	used := make([]int, n)
	visitedDiff := make(map[int]int)
	visitedSum := make(map[int]int)
	share := []int{}

	var dfs func(row int)
	dfs = func(row int) {
		if row == n {
			ans = append(ans, append([]string{}, generateStr(share)...))
			return
		}

		for col := 0; col < n; col++ {
			if used[col] == 0 && visitedSum[row+col] == 0 && visitedDiff[row-col] == 0 {
				used[col] = 1
				share = append(share, col)
				visitedSum[row+col] = 1
				visitedDiff[row-col] = 1
				dfs(row + 1)
				visitedDiff[row-col] = 0
				visitedSum[row+col] = 0
				share = share[:len(share)-1]
				used[col] = 0
			}
		}
	}
	dfs(0)
	return ans
}

func generateStr(nums []int) []string {
	ans := []string{}
	for i := 0; i < len(nums); i++ {
		b := make([]byte, len(nums))
		for j := 0; j < len(nums); j++ {
			b[j] = '.'
		}
		b[nums[i]] = 'Q'
		ans = append(ans, string(b))
	}
	return ans
}

func main() {
	ans := solveNQueens(1)
	for i := 0; i < len(ans); i++ {
		fmt.Println(ans[i])
	}
}
