package basic

func AddU(x, y int) int {
	return x + y
}

func DivideU(x, y int) float32 {
	if y == 0 {
		return 0.
	}
	return float32(x) / float32(y)
}

func UnitTestSrc() {
	// ↓テスト時は呼ばれないようにコメントアウトする
	// x, y := 3, 5
	// fmt.Printf("%v %v\n", AddU(x, y), DivideU(x, y))
}
