package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"reflect"
)

func EncryptPassword(password string) string {
	h := md5.New()
	if _, err := io.WriteString(h, password); err != nil {
		return ""
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

// struct -> map
func Struct2map(obj any) (data map[string]any, err error) {
	// 通过反射将结构体转换成map
	data = make(map[string]any)
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)
	for i := 0; i < objT.NumField(); i++ {
		fileName, ok := objT.Field(i).Tag.Lookup("json")
		if ok {
			data[fileName] = objV.Field(i).Interface()
		} else {
			data[objT.Field(i).Name] = objV.Field(i).Interface()
		}
	}
	return data, nil
}

func GetCommonElems(ids1, ids2 []int64) []int64 {
	m := make(map[int64]struct{}, len(ids1))
	for _, v := range ids1 {
		m[v] = struct{}{}
	}

	ans := make([]int64, 0, len(ids1)/2)
	for _, v := range ids2 {
		if _, ok := m[v]; ok {
			ans = append(ans, v)
		}
	}

	return ans
}
