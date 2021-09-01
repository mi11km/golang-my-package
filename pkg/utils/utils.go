package utils

func NotRepeat(s []int) []int {
	pp := s[0]
	p := s[1]
	isRepeat := pp == p
	for i := range s[2:] {
		i += 2
		if isRepeat && s[i] == p {
			// 3連続していたら
			for j := range s[i+1:] {
				j += i + 1
				if s[i] != s[j] {
					s = insert(s, i, s[j])
					s = delete(s, j+1)
					break
				}
			}
		}
		pp, p = p, s[i]
		isRepeat = pp == p
	}
	return s
}

func insert(slice []int, pos, val int) []int {
	slice = append(slice[:pos+1], slice[pos:]...)
	slice[pos] = val
	return slice
}

func delete(slice []int, pos int) []int {
	slice = append(slice[:pos], slice[pos+1:]...)
	return slice
}
