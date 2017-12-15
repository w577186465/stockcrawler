package crawler

import(
	"regexp"
)

func Alllink (s string) ([]string) {
	reg := regexp.MustCompile(`href="(.*?)"`)
	matchs := reg.FindAllStringSubmatch(s, -1)
	var links []string
	for _, v := range (matchs) {
		links = append(links, v[1])
	}
	return links
}