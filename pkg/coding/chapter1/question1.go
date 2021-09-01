package chapter1

// IsDuplicated 与えられた文字列に重複する文字があるかないか判定する関数
func IsDuplicated(str string) bool {
	m := make(map[rune]struct{})
	for _, r := range str {
		if _, ok := m[r]; ok {
			return false
		} else {
			m[r] = struct{}{}
		}
	}
	return true
}

// IsUniqueChars 引数の文字列を ASCII(128文字) に限定した場合
func IsUniqueChars(str string) bool {
	if len(str) > 128 {
		return false
	}
	var charSet [128]bool
	for _, c := range str {
		if charSet[c] {
			return false
		}
		charSet[c] = true
	}
	return true
}
