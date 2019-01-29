package string

import (
	"strings"
	"time"
)

func IndexOf(str string, char string) int {
	return strings.Index(str, char)
}
func Join(array []string, j string) string {
	return strings.Join(array, j)
}
func Split(str string, char string) []string {
	return strings.Split(str, char)
}
func Trim(str string) string {
	return strings.TrimSpace(str)
}

/**
以后需要支持map[string]interface{}

*/
func Template(str string, m map[string]string) string {
	return Replace(str, m, "{", "}")
}
func Replace(str string, m map[string]string, prefix string, suffix string) string {
	for key, value := range m {
		str = strings.Replace(str, prefix+key+suffix, value, -1)
	}
	return str
}

func TimeFormat(str string, t time.Time) string {
	return t.Format(Replace(str, map[string]string{
		"YYYY": "2006",
		"MM":   "01",
		"dd":   "02",
		"hh":   "15",
		"mm":   "04",
		"ss":   "05",
		"SSS":  "000",
	}, "", ""))
}
