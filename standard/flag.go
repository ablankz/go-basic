package standard

import (
	"flag"
	"fmt"
)

// コマンドラインのオプション

func Flag() {
	var (
		max int
		msg string
		x   bool
	)

	flag.IntVar(&max, "n", 32, "処理数の最大値")
	flag.StringVar(&msg, "m", "", "処理メッセージ")
	// 与えられたらtrue, なければfalse
	flag.BoolVar(&x, "x", false, "拡張オプション")
	flag.Parse()

	fmt.Println("処理数の最大値 =", max)
	fmt.Println("処理メッセージ =", msg)
	fmt.Println("拡張オプション =", x)

}
