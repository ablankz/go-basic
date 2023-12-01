package calculator

import "fmt"

var offset float64 = 10
var Offset float64 = 1

func Sum(a float64, b float64) float64 {
	// パッケージ内ならprivateでも呼び出せる
	fmt.Println("multiply: ", multiply(a, b))
	return a + b + offset
}
