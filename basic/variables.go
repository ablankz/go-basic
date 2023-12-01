package basic

import "fmt"

// 定数定義
const secret = "abc"

// 型を定義
type Os int

const (
	// 1, 2, 3が振られる
	Mac Os = iota + 1
	Windows
	Linux
)

// 関数の外で一括定義
var (
	gi int
	gf float64
	gb bool
)

func Variables() {
	// 宣言のみの場合それぞれの型のデフォルト値で初期化
	// int 0, bool false, string "", pointe nil,...
	var i int
	fmt.Println(i)

	i += 1
	fmt.Println(i)

	fmt.Println(gi)
	fmt.Println(gf)
	fmt.Println(gb)

	// 初期値の設定も可能
	var i1 int = 1
	fmt.Println(i1)

	// 初期値を設定する場合、自動的に予測される
	var i2 = 2
	fmt.Println(i2)

	// 初期値の設定が必須
	i3 := 3
	fmt.Println(i3)

	// 型を指定できる
	ui := uint(4)
	fmt.Println(ui)
	// 書式設定可能(%v: 型に合わせた書式, %T: 型, %p: 変数ポインタ)
	// fmt.Printf("ui: %v %T %p\n", ui, ui, &ui)
	// 引数を指定できる
	fmt.Printf("ui: %[1]v %[1]T\n", ui)

	// 複数の型
	f := 1.23
	s := "hello"
	b := true

	fmt.Printf("f: %[1]v %[1]T\n", f)
	fmt.Printf("s: %[1]v %[1]T\n", s)
	fmt.Printf("b: %[1]v %[1]T\n", b)

	// 分割代入可能
	pi, title := 3.14, "Go"
	fmt.Printf("pi: %v title: %v\n", pi, title)

	// 異なる型同士では計算不可
	z := float64(i) + f
	fmt.Printf("z: %[1]v %[1]T\n", z)

	// 定数は変更不可
	// secret = 2
	fmt.Printf("secret: %[1]v %[1]T\n", secret)

	fmt.Printf("Mac:%v Windows:%v Linux:%v\n", Mac, Windows, Linux)
}
