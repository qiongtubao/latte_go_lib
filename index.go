package latte_go_lib

import (
	"encoding/json"
	"errors"
	"reflect"
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

func ToInt64(data interface{}) (int64, error) {
	switch data.(type) {
	case int64:
		return data.(int64), nil
	case int:
		string := strconv.Itoa(data.(int))
		result, _ := strconv.ParseInt(string, 10, 64)
		return result, nil
	case string:
		return strconv.ParseInt(data.(string), 10, 64)
	default:
		return 0, errors.New("change error")
	}
}
func ToInt32(data interface{}) (int32, error) {
	return 0, nil
}
func ToInt16(data interface{}) (int32, error) {
	return 0, nil
}
func ToInt8(data interface{}) (int8, error) {
	return 0, nil
}
func ToUint8(data interface{}) (uint8, error) {
	return 0, nil
}
func ToUint16(data interface{}) (uint16, error) {
	return 0, nil
}
func ToUint32(data interface{}) (uint32, error) {
	return 0, nil
}

func ToUint64(data interface{}) (uint64, error) {
	return 0, nil
}
func ToInt(data interface{}) (int, error) {
	switch data.(type) {
	case int, int64:
		return data.(int), nil
	default:
		return 0, errors.New("change error")
	}
}
func ToString(data interface{}) (string, error) {
	switch data.(type) {
	case string:
		return data.(string), nil
	case int:
		return strconv.Itoa(data.(int)), nil
	case int64:
		return strconv.FormatInt(data.(int64), 10), nil
	default:
		return "", errors.New("暂不支持其他转换")
	}
}
func TypeChange(data interface{}, kind reflect.Kind) (interface{}, error) {
	switch kind {
	case reflect.Int:
		return ToInt(data)
	case reflect.Int64:
		return ToInt64(data)
	case reflect.String:
		return ToString(data)
	}
	return data, nil
	// switch kind {

	// case reflect.Int64:
	// 	return ToInt64(data)
	// case reflect.Int32:
	// 	return ToInt32(data)
	// case reflect.Int16:
	// 	return ToInt16(data)
	// case reflect.Int8:
	// 	return ToInt8(data)
	// case reflect.Uint8:
	// 	return ToUint8(data)
	// case reflect.Uint16:
	// 	return ToUint16(data)
	// case reflect.Uint32:
	// 	return ToUint32(data)
	// case reflect.Uint64:
	// 	return ToUint64(data)
	// case reflect.Float32:
	// 	return ToFloat32(data)
	// case reflect.Float64:
	// 	return ToFloat64(data)
	// case reflect.Slice:
	// 	return ToSlice(data)
	// case reflect.Struct:
	// 	return
	// }
}
