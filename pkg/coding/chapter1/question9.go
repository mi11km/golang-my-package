package chapter1

import "strings"

// IsRotateString s2がs1を回転させたものかどうか判定する関数
func IsRotateString(s1, s2 string) bool {
	if len(s1) == len(s2) && len(s1) > 0 {
		s1s1 := strings.Builder{}
		s1s1.WriteString(s1)
		s1s1.WriteString(s1)
		return strings.Contains(s1s1.String(), s2)
	}
	return false
}
