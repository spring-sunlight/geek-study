package example_DFS_BFS

import "fmt"

/*

	start: "AACCGGTT"
	end:   "AAACGGTA"
	bank: ["AACCGGTA", "AACCGCTA", "AAACGGTA"]
	返回值: 2

	思路: 最少次数使用 BFS,每次生成一个新的基因序列,如果在基因库中,就放到队列里,每次遍历层数+1
		 注意访问过的结点就不需要再访问了可以减少运算时间
*/
func minMutation(start string, end string, bank []string) int {
	ans := 0
	bankMap := make(map[string]struct{})
	visited := make(map[string]struct{})
	for i := 0; i < len(bank); i++ {
		bankMap[bank[i]] = struct{}{}
	}
	count := 0
	queue := []string{start}

	for len(queue) != 0 {
		count++
		temp := []string{}
		for i := 0; i < len(queue); i++ {
			for j := 0; j < 8; j++ {
				change := changeOnce(queue[i], j)
				for k := 0; k < len(change); k++ {

					if _, ok := bankMap[change[k]]; !ok {
						continue
					}
					if _, ok := visited[change[k]]; ok {
						continue
					}
					if change[k] == end {
						ans = count
						return ans
					}
					visited[change[k]] = struct{}{}

					temp = append(temp, change[k])
				}
			}
		}
		queue = temp
	}
	return -1
}

func changeOnce(s string, index int) []string {
	res := []string{}

	gene := []byte{'A', 'G', 'C', 'T'}
	for i := 0; i < 4; i++ {
		if gene[i] == s[index] {
			continue
		}
		byteS := []byte(s)
		byteS[index] = gene[i]
		res = append(res, string(byteS))
	}
	return res
}

func main() {
	start := "AAAAAAAA"
	end := "CCCCCCCC"
	bank := []string{"AAAAAAAA", "AAAAAAAC", "AAAAAACC", "AAAAACCC", "AAAACCCC", "AACACCCC", "ACCACCCC", "ACCCCCCC", "CCCCCCCA"}

	fmt.Println(changeOnce(start, 0))
	fmt.Println(minMutation(start, end, bank))

}
