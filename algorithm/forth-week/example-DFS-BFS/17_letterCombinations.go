package example_DFS_BFS

import "fmt"

/*
	题目大意: 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
			 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

	示例 1：输入：digits = "23"
		   输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]

	思路:类似求子集,第一个位置选一个,第二个位置选一个,总共选 digits 的长度
*/

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	ans := []string{}
	number := make(map[byte]string)
	number['2'] = "abc"
	number['3'] = "def"
	number['4'] = "ghi"
	number['5'] = "jkl"
	number['6'] = "mno"
	number['7'] = "pqrs"
	number['8'] = "tuv"
	number['9'] = "wxyz"
	var dfs func(index int, str string)
	dfs = func(index int, str string) {
		if index == len(digits) {
			ans = append(ans, str)
			return
		}
		for _, ch := range number[digits[index]] {
			dfs(index+1, str+string(ch))
		}
	}
	dfs(0, "")
	return ans
}

func main() {
	digits := "23"
	fmt.Println(letterCombinations(digits))
}
