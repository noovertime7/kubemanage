package utils

import "strconv"

// ParseInt64 将字符串转换为 int64
func ParseInt64(s string) (int64, error) {
	if len(s) == 0 {
		return 0, nil
	}
	return strconv.ParseInt(s, 10, 64)
}

// ParseInt 将字符串转换为 int64
func ParseInt(s string) (int, error) {
	if len(s) == 0 {
		return 0, nil
	}
	return strconv.Atoi(s)
}
