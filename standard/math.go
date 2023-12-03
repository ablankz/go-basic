package standard

import (
	"fmt"
	"math"
)

func Math() {
	fmt.Println(math.Pi)

	fmt.Println(math.Sqrt2)

	fmt.Println(math.MaxFloat32)
	fmt.Println(math.SmallestNonzeroFloat32)
	fmt.Println(math.MaxFloat64)
	fmt.Println(math.SmallestNonzeroFloat64)
	fmt.Println(math.MaxInt64)
	fmt.Println(math.MinInt64)

	fmt.Println(math.Abs(100))
	fmt.Println(math.Abs(-100))

	fmt.Println(math.Pow(0, 2))
	fmt.Println(math.Pow(2, 2))

	fmt.Println(math.Sqrt(2))
	// 立方根
	fmt.Println(math.Cbrt(8))

	fmt.Println(math.Max(1, 2))
	fmt.Println(math.Min(1, 2))

	// 切り捨て
	fmt.Println(math.Trunc(3.14))
	fmt.Println(math.Trunc(-3.14))
	// 引数の数値より小さい最大の整数
	fmt.Println(math.Floor(3.5))  // 3
	fmt.Println(math.Floor(-3.5)) // -4
	// 引数の数値より大きい最小の整数
	fmt.Println(math.Ceil(3.5))  // 4
	fmt.Println(math.Ceil(-3.5)) // -3

	n := math.Sqrt(-1)
	// nanはいかなる値と比較してもfalseになるため以下のような関数がある
	fmt.Println(math.IsNaN(n))

	fmt.Println(math.Inf(0)) // +Inf

	fmt.Println(math.Inf(-1)) // -Inf

	fmt.Println(math.NaN()) //NaN
}
