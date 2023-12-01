package basic

import "fmt"

func Slice() {
	fmt.Println("-----------------------配列-----------------------")
	// 0が３つの要素で初期化
	var a1 [3]int
	// 以下で明示的な代入を行う
	var a2 = [3]int{10, 20, 30}

	// 以下の場合、要素の明記が必要
	// 以下のようにすることで要素数を自動で算出してくれる
	a3 := [...]int{40, 50}

	fmt.Printf("%v %v %v\n", a1, a2, a3)
	// キャパシティとレングスの表示
	fmt.Printf("%v %v\n", len(a3), cap(a3))
	// 以下のように要素数が異なるだけで別の型と認識されるためa2=3などの代入は不可
	fmt.Printf("%T %T\n", a2, a3)

	fmt.Println("---------------------スライス---------------------")
	// 要素数を空にすることでスライスを宣言
	var s1 []int
	// :=で空で宣言する場合はカーリブラケット{}のみつける
	s2 := []int{}
	fmt.Printf("s1: %[1]T %[1]v %v %v\n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %[1]T %[1]v %v %v\n", s2, len(s2), cap(s2))

	// var宣言では初期値がnilとして認識される
	fmt.Println(s1 == nil, s2 == nil)

	// スライスへの値の追加
	s1 = append(s1, 1, 2, 3)
	// 型の変更なしでcap, lenが増加
	fmt.Printf("s1: %[1]T %[1]v %v %v\n", s1, len(s1), cap(s1))
	// 初期値ありで宣言
	s3 := []int{10, 20, 30}
	// 3ドットをつけることで可変長引数の形に変換して渡せる
	s1 = append(s1, s3...)
	fmt.Printf("s1: %[1]T %[1]v %v %v\n", s1, len(s1), cap(s1))

	// make: 指定したキャパシティの値に応じてメモリ領域を確保する
	// 以下では要素数が0でキャパシティが2のスライスを作成
	s4 := make([]int, 0, 2)
	fmt.Printf("s4: %[1]T %[1]v %v %v\n", s4, len(s4), cap(s4))
	// 指定のキャパシティを超えても要素追加可能
	s4 = append(s4, 1, 2, 3, 4)
	// cap, lenが増加
	fmt.Printf("s4: %[1]T %[1]v %v %v\n", s4, len(s4), cap(s4))

	// 以下では要素数が4(すべて0)でキャパシティが6のスライスを作成
	s5 := make([]int, 4, 6)
	fmt.Printf("s5: %[1]T %[1]v %v %v\n", s5, len(s5), cap(s5))
	// s5から要素番号1以上3未満を取得してs6に代入
	s6 := s5[1:3]
	s6[1] = 10
	// s6の変更によりs5も書き換わる
	fmt.Printf("s5: %[1]T %[1]v %v %v\n", s5, len(s5), cap(s5))
	fmt.Printf("s6: %[1]T %[1]v %v %v\n", s6, len(s6), cap(s6))
	s6 = append(s6, 2)
	// s5の要素[3]も書き換わっている
	fmt.Printf("s5: %[1]T %[1]v %v %v\n", s5, len(s5), cap(s5))
	fmt.Printf("s6: %[1]T %[1]v %v %v\n", s6, len(s6), cap(s6))

	// メモリを共有させたくない場合, コピーもとの長さを計算
	sc6 := make([]int, len(s5[1:3]))
	fmt.Printf("s5 source of copy: %[1]T %[1]v %v %v\n", s5, len(s5), cap(s5))
	fmt.Printf("sc6 dst copy before: %[1]T %[1]v %v %v\n", sc6, len(sc6), cap(sc6))
	// 引数はコピー元, コピー先
	copy(sc6, s5[1:3])
	fmt.Printf("sc6 dst copy after: %[1]T %[1]v %v %v\n", sc6, len(sc6), cap(sc6))
	sc6[1] = 12
	// コピー先でのみの変更になっている
	fmt.Printf("s5: %[1]T %[1]v %v %v\n", s5, len(s5), cap(s5))
	fmt.Printf("sc6: %[1]T %[1]v %v %v\n", sc6, len(sc6), cap(sc6))

	// メモリを部分的に共有させる
	s5 = make([]int, 4, 6)
	// 一番最後のx:y:zのz未満の要素までメモリを共有する
	// 以下の場合だと、1, 2を共有(fs6でいうと0, 1)
	fs6 := s5[1:3:3]
	fmt.Printf("s5: %[1]T %[1]v %v %v\n", s5, len(s5), cap(s5))
	fmt.Printf("fs6: %[1]T %[1]v %v %v\n", fs6, len(fs6), cap(fs6))
	fs6[0] = 6
	fs6[1] = 7
	fs6 = append(fs6, 8)
	// fs6の要素2(s5の要素3)はメモリを共有してないため、s5は要素番号3を8に上書きされない
	fmt.Printf("s5: %[1]T %[1]v %v %v\n", s5, len(s5), cap(s5))
	fmt.Printf("fs6: %[1]T %[1]v %v %v\n", fs6, len(fs6), cap(fs6))
	s5[3] = 9
	// fs6の要素2(s5の要素3)はメモリを共有してないため、s6は要素番号2を9に上書きされない
	fmt.Printf("s5: %[1]T %[1]v %v %v\n", s5, len(s5), cap(s5))
	fmt.Printf("fs6: %[1]T %[1]v %v %v\n", fs6, len(fs6), cap(fs6))
}
