package calculator

import (
	"fmt"
	// . "fmt" // インポートの名前省略(非推奨)
	// f "fmt" // インポートのエイリアス(非推奨)
)

var offset float64 = 10
var Offset float64 = 1

func Sum(a float64, b float64) float64 {
	// パッケージ内ならprivateでも呼び出せる
	fmt.Println("multiply: ", multiply(a, b))
	// f.Println("インポートのエイリアス(非推奨)")
	// Println("インポートの名前省略(非推奨)")
	return a + b + offset
}
