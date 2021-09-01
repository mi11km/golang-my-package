package chapter1

// ToZero M*N行列について０の要素がある行と列をすべてゼロにする関数
func ToZero(matrix [][]int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	rows := make([]int, 0, m)
	columns := make([]int, 0, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				rows = append(rows, i)
				columns = append(columns, j)
			}
		}
	}
	for _, row := range rows {
		nullifyRow(matrix, row)
	}
	for _, col := range columns {
		nullifyColumn(matrix, col)
	}
	return true
}

// SetZeros １行目と１列目にその行や列が 0 を保持してるかどうか記憶することで空間計算量を O(1) にできる
func SetZeros(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	rowHasZero, colHasZero := false, false
	for i := 0; i < len(matrix[0]); i++ {
		if matrix[0][i] == 0 {
			rowHasZero = true
			break
		}
	}
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			colHasZero = true
			break
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			nullifyRow(matrix, i)
		}
	}
	for i := 1; i < len(matrix[0]); i++ {
		if matrix[0][i] == 0 {
			nullifyColumn(matrix, i)
		}
	}
	if rowHasZero {
		nullifyRow(matrix, 0)
	}
	if colHasZero {
		nullifyColumn(matrix, 0)
	}
}

func nullifyRow(matrix [][]int, row int) {
	for i := 0; i < len(matrix[0]); i++ {
		matrix[row][i] = 0
	}
}

func nullifyColumn(matrix [][]int, col int) {
	for i := 0; i < len(matrix); i++ {
		matrix[i][col] = 0
	}
}
