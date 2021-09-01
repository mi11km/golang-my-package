package utils

func Sum(arr []float64) float64 {
	sum := 0.0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func Average(arr []float64) float64 {
	return Sum(arr) / float64(len(arr))
}

// LeastSquaresMethod 最小2乗法: 1次関数 y = ax + b で近似する場合の a と b を標準出力
func LeastSquaresMethod(x, y []float64) (float64, float64) {
	xy := make([]float64, len(x))
	x_2 := make([]float64, len(x))
	for i := 0; i < len(x); i++ {
		xy[i] = x[i] * y[i]
		x_2[i] = x[i] * x[i]
	}
	x_av, y_av, xy_av, x_2_av := Average(x), Average(y), Average(xy), Average(x_2)

	// y = ax + b
	a := (xy_av - x_av*y_av) / (x_2_av - x_av*x_av)
	b := -1*a*x_av + y_av

	return a, b
}

// LagrangeInterpolation ラグランジュ補間多項式で x = a のときの y を求める
func LagrangeInterpolation(x, y []float64, a float64) float64 {
	result := 0.0
	tmp1, tmp2 := 1.0, 1.0
	for i := 0; i < len(x); i++ {
		tmp1, tmp2 = 1.0, 1.0
		tmp1 *= y[i]
		for j := 0; j < len(x); j++ {
			if i == j {
				continue
			} else {
				tmp1 *= a - x[j]
				tmp2 *= x[i] - x[j]
			}
		}
		result += tmp1 / tmp2
	}
	return result
}
