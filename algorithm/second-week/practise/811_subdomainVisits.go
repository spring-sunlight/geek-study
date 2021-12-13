package practise

import (
	"fmt"
	"strconv"
	"strings"
)

//输入:
//["900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"]
//输出:
//["901 mail.com","50 yahoo.com","900 google.mail.com","5 wiki.org","5 org","1
//intel.mail.com","951 com"]
//说明:
//按照假设，会访问"google.mail.com" 900次，"yahoo.com" 50次，"intel.mail.com" 1次，"wiki.org"
//5次。
//而对于父域名，会访问"mail.com" 900+1 = 901次，"com" 900 + 50 + 1 = 951次，和 "org" 5 次。
func subdomainVisits(cpdomains []string) []string {
	domainMap := make(map[string]int)
	for i := 0; i < len(cpdomains); i++ {
		domain := strings.Split(cpdomains[i], " ")
		count, _ := strconv.Atoi(domain[0])
		str := domain[1]
		for strings.Index(str, ".") != -1 {
			domainMap[str] += count
			str = str[strings.Index(str, ".")+1:]
		}
		domainMap[str] += count
	}
	fmt.Println(domainMap)
	ans := []string{}
	for k, v := range domainMap {
		ans = append(ans, strconv.Itoa(v)+" "+k)
	}
	return ans
}

func main() {
	cpdomains := []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}
	fmt.Println(subdomainVisits(cpdomains))
}
