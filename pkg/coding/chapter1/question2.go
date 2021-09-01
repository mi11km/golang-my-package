package chapter1

// IsPermutation ２つの文字列が与えられたとき、片方がもう片方の並び替えになっているかどうかを決定する関数
func IsPermutation(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	m := make(map[rune]int)
	for _, r := range str1 {
		m[r]++
	}
	for _, r := range str2 {
		m[r]--
	}
	for _, r := range str1 {
		if m[r] != 0 {
			return false
		}
	}
	return true
}
