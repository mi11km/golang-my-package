package chapter1

// IsTransformedByOneStep 文字の挿入or削除or置き換えのどれか(もしくは操作なし)で一方の文字列をもう一方に変換できるかどうか
func IsTransformedByOneStep(str1, str2 string) bool {
	if str1 == str2 {
		return true
	}
	diffLen := len(str1) - len(str2)
	if diffLen > 1 || diffLen < -1 {
		return false
	} else if diffLen == -1 {
		str1, str2 = str2, str1 // str1 > str2
	}
	shorterStr := make([]rune, len(str2)+1)
	for i, r := range str2 {
		shorterStr[i] = r
	}
	diffCounts := 0
	for i, r := range str1 {
		if diffLen != 0 {
			i -= diffCounts
		}
		if shorterStr[i] != r {
			if diffCounts == 0 {
				diffCounts++
			} else {
				return false
			}
		}
	}
	return true
}
