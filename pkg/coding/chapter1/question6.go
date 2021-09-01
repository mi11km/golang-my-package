package chapter1

import (
	"strings"
)

func CompressStr(str string) string {
	var (
		countConsecutive int
		charConsecutive  rune
	)
	compressed := make([]rune, 0, len(str))
	for _, r := range str {
		if countConsecutive == 0 {
			charConsecutive = r
			countConsecutive++
		} else {
			if charConsecutive != r {
				compressed = append(compressed, charConsecutive, rune(countConsecutive+48))
				countConsecutive = 1
				charConsecutive = r
			} else {
				countConsecutive++
			}
		}
	}
	compressed = append(compressed, charConsecutive, rune(countConsecutive+48))

	if len(compressed) >= len(str)-1 {
		return str
	}
	return string(compressed)
}

func CompressStrWithStringsBuilder(str string) string {
	var (
		countConsecutive int
		charConsecutive  rune
	)
	compressed := &strings.Builder{}
	compressed.Grow(len(str))
	for _, r := range str {
		if countConsecutive == 0 {
			charConsecutive = r
			countConsecutive++
		} else {
			if charConsecutive != r {
				compressed.WriteByte(byte(charConsecutive))
				compressed.WriteByte(byte(countConsecutive + 48))
				countConsecutive = 1
				charConsecutive = r
			} else {
				countConsecutive++
			}
		}
	}
	compressed.WriteByte(byte(charConsecutive))
	compressed.WriteByte(byte(countConsecutive + 48))

	if compressed.Len() >= len(str)-1 {
		return str
	}
	return compressed.String()
}

// CompressStrWithStringsBuilder2 一番早い アルファベット限定
func CompressStrWithStringsBuilder2(str string) string {
	if len(str) < 2 {
		return str
	}
	countConsecutive := 1
	compressed := &strings.Builder{}
	compressed.Grow(len(str))
	for i := 1; i < len(str); i++ {
		if str[i-1] == str[i] {
			countConsecutive++
		} else {
			compressed.WriteByte(str[i-1])
			compressed.WriteByte(byte(countConsecutive + 48))
			countConsecutive = 1
		}
	}
	compressed.WriteByte(str[len(str)-1])
	compressed.WriteByte(byte(countConsecutive + 48))
	if compressed.Len() >= len(str) {
		return str
	}
	return compressed.String()
}
