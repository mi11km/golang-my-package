package chapter1

import "strings"

func URLifyWithStringsPackage(str string, length int) string {
	var spaces int
	for _, r := range str {
		if r == ' ' {
			spaces++
		}
	}
	replaces := spaces - len(str) + length
	return strings.Replace(strings.Replace(str, " ", "%20", replaces), " ", "", -1)
}

func URLify(str string, length int) string {
	var spaces int
	for _, r := range str {
		if r == ' ' {
			spaces++
		}
	}
	replaces := spaces - len(str) + length
	runes := make([]rune, length+replaces*2)
	i := 0
	for _, r := range str {
		if r == ' ' {
			if replaces > 0 {
				runes[i] = '%'
				runes[i+1] = '2'
				runes[i+2] = '0'
				i += 3
				replaces--
			}
		} else {
			runes[i] = r
			i++
		}
	}
	return string(runes)
}

func URLifyArray(runes []rune, length int) string {
	var spaces, index, i int
	for i = 0; i < length; i++ {
		if runes[i] == ' ' {
			spaces++
		}
	}
	index = length + spaces*2 - 1
	str := make([]rune, index+1)
	index = 0
	for i = 0; i < length; i++ {
		if runes[i] == ' ' {
			str[index] = '%'
			str[index+1] = '2'
			str[index+2] = '0'
			index += 3
		} else {
			str[index] = runes[i]
			index++
		}
	}
	return string(str)
}
