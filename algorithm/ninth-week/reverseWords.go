package ninth_week

import "strings"

func reverseWords(s string) string {
	str := strings.Fields(strings.Trim(s, " "))
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}
	return strings.Join(str, " ")
}
