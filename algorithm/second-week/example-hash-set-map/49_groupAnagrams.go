package example_hash_set_map

import "sort"

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for i := 0; i < len(strs); i++ {
		strToByte := []byte(strs[i])
		sort.Slice(strToByte, func(i, j int) bool {
			return strToByte[i] < strToByte[j]
		})
		str := string(strToByte)
		m[str] = append(m[str], strs[i])
	}
	ans := [][]string{}
	for _, v := range m {
		ans = append(ans, v)
	}
	return ans
}
