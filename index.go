package lib

import (
	"strconv"
)

func IntToStr(i int) string {
	return strconv.Itoa(i)
}
func StrToInt(data string) (int, error) {
	return strconv.Atoi(data)
}
