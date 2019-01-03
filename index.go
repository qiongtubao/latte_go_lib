package latte_go_lib

import (
	"strconv"
)

func IntToStr(i int) string {
	return strconv.Itoa(i)
}
func StrToInt(data string) (int, error) {
	return strconv.Atoi(data)
}
func IntToInt64(data int) int64 {
	string := strconv.Itoa(data)
	result, _ := strconv.ParseInt(string, 10, 64)
	return result
}

func Int64ToInt(data int64) int {
	string := strconv.FormatInt(data, 10)
	intstr, _ := strconv.Atoi(string)
	return intstr
}
