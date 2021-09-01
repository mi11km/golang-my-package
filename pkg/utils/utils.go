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

func f(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return f(n-1) + f(n-2)
}

func fNew(n int) int {
	memo := make([]int, 1000)
	memo[0], memo[1] = 0, 1
	return fm(n, memo)
}

func fm(n int, memo []int) int {
	if n == 0 || n == 1 {
		return n
	}
	if memo[n] == 0 {
		memo[n] = fm(n-1, memo) + fm(n-2, memo)
	}
	return memo[n]
}

func fd(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	a := 0
	b := 1
	c := 1
	for i := 2; i <= n; i++ {
		c = a + b
		a = b
		b = c
	}
	return c
}