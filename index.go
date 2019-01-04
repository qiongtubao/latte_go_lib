package latte_go_lib

import (
	"encoding/json"
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

func ToJson(data interface{}) (string, error) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

func JSONToMap(data string) (map[string]interface{}, error) {
	var m map[string]interface{}
	err := json.Unmarshal([]byte(data), &m)
	return m, err
}
