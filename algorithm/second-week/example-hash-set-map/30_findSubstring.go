package example_hash_set_map

func findSubstring(s string, words []string) []int {
	//1.将 words 以单词为维度构建 map,与顺序无关就是与个数有关,key:单词,value:单词个数
	wordsMap := make(map[string]int)
	for i := 0; i < len(words); i++ {
		wordsMap[words[i]]++
	}
	ans := []int{}
	//2.遍历 s,步长为 1,查看每个子字符串,以 words 中单词分割,查看每个单词个数是否符合 wordsMap
	for i := 0; i+len(words[0])*len(words) <= len(s); i++ {
		flag := true
		str := s[i : i+len(words[0])*len(words)]
		subStrMap := make(map[string]int)
		for j := 0; j+len(words[0]) <= len(str); j += len(words[0]) {
			subStrMap[str[j:j+len(words[0])]]++
			if wordsMap[str[j:j+len(words[0])]] == 0 || wordsMap[str[j:j+len(words[0])]] < subStrMap[str[j:j+len(words[0])]] {
				flag = false
				break
			}
		}
		if flag {
			ans = append(ans, i)
		}
	}
	return ans
}
