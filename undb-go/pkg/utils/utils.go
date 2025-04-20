package utils

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"time"
)

// GenerateID 生成唯一ID
func GenerateID(prefix string) string {
	b := make([]byte, 16)
	rand.Read(b)
	return prefix + "_" + base64.URLEncoding.EncodeToString(b)
}

// Now 获取当前时间
func Now() time.Time {
	return time.Now().UTC()
}

// FormatTime 格式化时间
func FormatTime(t time.Time) string {
	return t.Format(time.RFC3339)
}

// ParseTime 解析时间
func ParseTime(s string) (time.Time, error) {
	return time.Parse(time.RFC3339, s)
}

// ToJSON 转换为JSON
func ToJSON(v interface{}) (string, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// FromJSON 从JSON解析
func FromJSON(s string, v interface{}) error {
	return json.Unmarshal([]byte(s), v)
}

// Contains 检查切片是否包含元素
func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

// Unique 去重
func Unique[T comparable](s []T) []T {
	seen := make(map[T]bool)
	result := []T{}
	for _, v := range s {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

// Map 映射
func Map[T, U any](s []T, f func(T) U) []U {
	result := make([]U, len(s))
	for i, v := range s {
		result[i] = f(v)
	}
	return result
}

// Filter 过滤
func Filter[T any](s []T, f func(T) bool) []T {
	result := []T{}
	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}
