package chapter1

func Rotate90Degree(matrix [][]rune) bool {
	if len(matrix) == 0 || len(matrix) != len(matrix[0]) {
		return false
	}
	n := len(matrix)
	for layer := 0; layer < n/2; layer++ {
		first := layer
		last := n - 1 - layer
		for i := first; i < last; i++ {
			offset := i - first
			top := matrix[first][i]                                // 上端の値を保持
			matrix[first][i] = matrix[last-offset][first]          // 左端　→　上端
			matrix[last-offset][first] = matrix[last][last-offset] // 下端　→　左端
			matrix[last][last-offset] = matrix[i][last]            // 右端　→　下端
			matrix[i][last] = top                                  // 上端　→　右端
		}
	}
	return true
}

// Rotate 4回やるとなぜかもとに戻らない
func Rotate(matrix [][]rune) bool {
	if len(matrix) == 0 || len(matrix) != len(matrix[0]) {
		return false
	}
	start, end := 0, len(matrix)-1
	for start < end {
		for i := start; i < end; i++ {
			matrix[start][i], matrix[i][end] = matrix[i][end], matrix[start][i]
			matrix[start][i], matrix[end][end-i] = matrix[end][end-i], matrix[start][i]
			matrix[start][i], matrix[end-i][start] = matrix[end-i][start], matrix[start][i]
		}
		start++
		end--
	}
	return true
}

/* ---------------- */

//func PrintMatrix(matrix [][]rune) {
//	if matrix == nil {
//		fmt.Println("can not print nil image")
//		return
//	}
//	for i := 0; i < len(matrix); i++ {
//		for j := 0; j < len(matrix); j++ {
//			fmt.Printf("%s ", string(matrix[i][j]))
//		}
//		fmt.Printf("\n")
//	}
//	fmt.Printf("\n")
//}
//
//func GenerateMatrix(n int) [][]rune {
//	if n <= 0 {
//		return nil
//	}
//	matrix := make([][]rune, 0, n)
//	for i := 0; i < n; i++ {
//		matrix = append(matrix, make([]rune, 0, n))
//		for j := 0; j < n; j++ {
//			matrix[i] = append(matrix[i], rune(97+j))
//		}
//	}
//	return matrix
//}
