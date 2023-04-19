package utils

import "strconv"

func Int64ToStr(v int64) string {
	return strconv.FormatInt(v, 10)
}

func IntToStr(v int) string {
	return strconv.Itoa(v)
}

func StrToInt(v string) (int, error) {
	return strconv.Atoi(v)
}
