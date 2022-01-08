package example_DFS_BFS

import "fmt"

func numIslands(grid [][]byte) int {
	ans := 0
	rows := len(grid)
	cols := len(grid[0])
	visited := make([][]int, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]int, cols)
	}
	dx := []int{-1, 0, 0, 1}
	dy := []int{0, -1, 1, 0}

	res := 0
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if isValid(x, y, rows, cols) && grid[x][y] == '1' && visited[x][y] == 0 {
			res++
			//fmt.Println(res)
			visited[x][y] = 1
			for i := 0; i < 4; i++ {
				nx := x + dx[i]
				ny := y + dy[i]
				dfs(nx, ny)
			}
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			res = 0
			dfs(i, j)
			if res != 0 {
				ans++
			}
			//fmt.Println("i:", i, "j:", j, "res: ", res)
			//ans = max(res, ans)
		}
	}
	return ans
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func isValid(x, y int, rows, cols int) bool {
	return x >= 0 && x < rows && y >= 0 && y < cols
}

func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}

	grid1 := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Println("main:", numIslands(grid))
	fmt.Println("main1:", numIslands(grid1))
}
