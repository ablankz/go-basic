package basic

import (
	"fmt"
	"unsafe"
)

// これらを実装した(レシーバに持つ)構造体はこのインタフェースを実装したと自動的にみなされる
type controller interface {
	speedUp() int
	speedDown() int
}

type vehicle struct {
	speed       int
	enginePower int
}

type bycycle struct {
	speed      int
	humanPower int
}

func (v *vehicle) speedUp() int {
	v.speed += 10 * v.enginePower
	return v.speed
}

func (v *vehicle) speedDown() int {
	v.speed -= 5 * v.enginePower
	return v.speed
}

func (b *bycycle) speedUp() int {
	b.speed += 2 * b.humanPower
	return b.speed
}

func (b *bycycle) speedDown() int {
	b.speed -= 1 * b.humanPower
	return b.speed
}

func speedUpAndDown(c controller) {
	fmt.Printf("current speed: %v\n", c.speedUp())
	fmt.Printf("current speed: %v\n", c.speedDown())
}

// 標準パッケージのStringerインタフェースを実装する
// これを実装していればPrintfなどの引数に渡せる
func (v vehicle) String() string {
	return fmt.Sprintf("Vehicle current speed is %v (enginePower %v)", v.speed, v.enginePower)
}

// 型アサーション
// v, ok = x.(T)
// 上記の構文で型アサーションできる
// xがTの型である場合はokにtrueが入り、以降xはその型とみなせる

func Interface() {
	// ポインタレシーバを呼び出すため
	// v := &vehicle{
	// 	speed:       0,
	// 	enginePower: 5,
	// }
	// 以下の記述でも可能
	v := &vehicle{0, 5}
	speedUpAndDown(v)
	b := &bycycle{0, 5}
	speedUpAndDown(b)
	// Stringerインタフェースを実装しているため呼び出せる
	fmt.Println(v)

	// 以下の２つは同じ意味
	// 空のインタフェースは0個のメソッドを実装する型とみなされるため、実質全ての型がこのインタフェースを実装している
	var i1 interface{}
	var i2 any
	// ともに値、型はnil
	fmt.Printf("%[1]v %[1]T %v\n", i1, unsafe.Sizeof(i1))
	fmt.Printf("%[1]v %[1]T %v\n", i2, unsafe.Sizeof(i2))
	checkType(i2)
	i2 = 1
	checkType(i2)
	i2 = 1.9
	checkType(i2)
	i2 = "hello"
	checkType(i2)
	i2 = &vehicle{1, 2}
	checkType(i2)
	// 以下のようにすることで型を指定して扱える
	// .()がなければanyで渡されるためコンパイルエラー
	speedUpAndDown(i2.(*vehicle))
}

func checkType(i any) {
	// i.(type)で型を取得できる
	switch i.(type) {
	case nil:
		fmt.Println("nil")
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown")
	}
}
