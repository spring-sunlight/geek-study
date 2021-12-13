package example_recursion

import "fmt"

func combine(n int, k int) [][]int {
	ans := [][]int{}
	share := []int{}

	var recur func(i int)

	recur = func(i int) {
		//剪枝,选的超过 k 个或者选的加剩下的不够 k 个就退出
		if len(share) > k || len(share)+n-i+1 < k {
			return
		}
		if i == n+1 {
			if len(share) == k {
				ans = append(ans, append([]int{}, share...))
			}
			return
		}

		recur(i + 1)

		share = append(share, i)
		recur(i + 1)

		share = share[:len(share)-1]
	}
	recur(1)
	return ans
}

func main() {
	fmt.Println(combine(4, 1))
}
