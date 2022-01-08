package example_graph

import "fmt"

//输入: edges = [[1,2], [2,3], [3,4], [1,4], [1,5]]
//输出: [1,4]

func findRedundantConnection(edges [][]int) []int {
	n := 0

	//找到图中值最大的点
	for i := 0; i < len(edges); i++ {
		if n < edges[i][0] {
			n = edges[i][0]
		}
		if n < edges[i][1] {
			n = edges[i][1]
		}
	}

	//初始化出边数组
	// [1, 6]
	// [6, 3]
	// [3, 4]   =>   [[] [6 4 5] [] [6 4] [3 1] [1] [1 3]]
	// [1, 4]
	// [1, 5]

	edge := make([][]int, n+1)
	visit := make([]bool, n+1)
	hasCycle := false

	//加边
	addEdge := func(i, j int) {
		edge[i] = append(edge[i], j)
	}

	//dfs 无向图找环
	var dfs func(int, int)

	dfs = func(x int, father int) {
		visit[x] = true
		for _, y := range edge[x] {
			if y == father {
				continue
			}
			if visit[y] {
				hasCycle = true
			} else {
				dfs(y, x)
			}
		}
	}
	for _, e := range edges {
		u, v := e[0], e[1]
		addEdge(u, v)
		addEdge(v, u)
		visit = make([]bool, n+1)
		dfs(u, -1)
		if hasCycle {
			return e
		}
	}

	return nil
}

func findRedundantConnection1(edges [][]int) []int {

	edge := make(map[int][]int)
	visit := make(map[int]bool)
	hasCycle := false
	//加边
	addEdge := func(u, v int) {
		edge[u] = append(edge[u], v)
	}
	//dfs 无向图找环
	var dfs func(u, father int)
	dfs = func(u, father int) {
		visit[u] = true
		for _, e := range edge[u] {
			if e == father {
				continue
			}
			if visit[e] == true {
				hasCycle = true
			} else {
				dfs(e, u)
			}
		}
	}

	for i := 0; i < len(edges); i++ {
		u := edges[i][0]
		v := edges[i][1]
		addEdge(u, v)
		addEdge(v, u)
		//判断是否出现环
		visit = make(map[int]bool)
		dfs(u, -1)
		if hasCycle == true {
			return edges[i]
		}
	}
	return nil
}

func main() {
	nums := [][]int{
		{1, 6},
		{6, 3},
		{3, 4},
		{1, 4},
		{1, 5},
	}
	fmt.Println(findRedundantConnection1(nums))
	fmt.Println(findRedundantConnection(nums))
}
