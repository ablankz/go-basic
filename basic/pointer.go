package basic

import (
	"fmt"
	"unsafe"
)

func Pointer() {
	var ui1 uint16
	fmt.Printf("memory address of ui1: %p\n", &ui1)
	var ui2 uint16
	fmt.Printf("memory address of ui2: %p\n", &ui2)

	// ポインタ変数(初期値nil(0x0))
	var p1 *uint16
	fmt.Printf("value of p1: %p\n", p1)
	// ui1のアドレスをp1に代入
	p1 = &ui1
	fmt.Printf("value of p1: %p\n", p1)
	// 変数のサイズを表示(どんな型のポインタでも8バイト固定)
	fmt.Printf("size of p1: %d[bytes]\n", unsafe.Sizeof(p1))
	// ポインタ変数のポインタ(アドレス)
	fmt.Printf("memory address of p1: %p\n", &p1)
	// デリファレンス(逆参照)
	fmt.Printf("value of ui1(dereference): %v\n", *p1)
	// p1の逆参照からui1を書き換える
	*p1 = 1
	fmt.Printf("value of ui1: %v\n", ui1)

	// ダブルポインタ(初期値nil(0x0))
	var pp1 **uint16
	fmt.Printf("value of pp1: %p\n", pp1)
	// p1のポインタをpp1に代入
	pp1 = &p1
	// 変数のサイズを表示(どんな型のポインタでも8バイト固定)
	fmt.Printf("size of pp1: %d[bytes]\n", unsafe.Sizeof(pp1))
	// ポインタ変数のポインタ(アドレス)
	fmt.Printf("memory address of pp1: %p\n", &pp1)
	// 1回のデリファレンス(逆参照)(ui1のポインタ)
	fmt.Printf("value of p1(dereference): %v\n", *pp1)
	// 2回のデリファレンス(逆参照)
	fmt.Printf("value of ui1(dereference): %v\n", **pp1)
}
