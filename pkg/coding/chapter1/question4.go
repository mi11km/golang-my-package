package chapter1

import (
	"strings"
)

// IsPalindromeOfPermutation 与えられた文字列の順列が回文かどうか判定する 空白文字は考慮せず、大文字小文字も区別しない
func IsPalindromeOfPermutation(str string) bool {
	str = strings.Replace(str, " ", "", -1)
	str = strings.ToLower(str)

	charCountMap := make(map[rune]int)
	for _, r := range str {
		charCountMap[r]++
	}

	isOddLen := len(str)%2 == 1
	isContainedOdd := false
	for _, count := range charCountMap {
		if count%2 == 1 {
			if !isContainedOdd && isOddLen {
				isContainedOdd = true
			} else {
				return false
			}
		}
	}

	return true
}

/* 順列回文判定：アルファベット限定(26文字)
アルファベットを26ビットの２進数にマップし、偶数個あれば0、奇数個は1として記録。
記録したものから１引いたビットと元のものの論理積が1ならok。
ex) 00010000 - 1 = 00001111, 00010000 & 00001111 = 0
    01010100 - 1 = 01010011, 01010100 & 01010011 = 01010000
*/

func IsPalindromeOfPermutationWithBitVector(str string) bool {
	str = strings.Replace(str, " ", "", -1)
	str = strings.ToLower(str)

	var charBitVector uint32
	for _, r := range str {
		charBitVector ^= 1 << (r - 97)
	}
	return charBitVector&(charBitVector-1) == 0
}
