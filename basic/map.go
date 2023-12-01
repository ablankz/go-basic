package basic

import "fmt"

func Map() {
	// mapの初期化(キーがstringで値がint)
	var m1 map[string]int
	m2 := map[string]int{}
	// var宣言の場合はnilで初期化
	fmt.Printf("%v %v \n", m1, m1 == nil)
	fmt.Printf("%v %v \n", m2, m2 == nil)
	// 値を代入
	m2["A"] = 10
	m2["B"] = 20
	m2["C"] = 0
	fmt.Printf("%v %v %v\n", m2, len(m2), m2["A"])
	// Aがキーのものを削除
	delete(m2, "A")
	// 存在しないキーの値を取ると0が返ってくる
	fmt.Printf("%v %v %v\n", m2, len(m2), m2["A"])
	// 存在するかどうかを識別するには以下のようにする
	// 存在しない場合はokにfalseが入る
	v, ok := m2["A"]
	fmt.Printf("%v %v\n", v, ok)
	v, ok = m2["C"]
	fmt.Printf("%v %v\n", v, ok)

	// for分では以下のように値を取得する
	// ただし順番は保証されない
	for k, v := range m2 {
		fmt.Printf("%v %v\n", k, v)
	}
}
