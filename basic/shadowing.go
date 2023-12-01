package basic

import "fmt"

func Shadowing() {
	ok, result := true, "A"
	if ok {
		// このブロック内でシャドウィングされる
		// アドレスも別のものになる
		result := "B"
		fmt.Printf("value: %v, pointer: %p\n", result, &result)
	} else {
		// このブロック内でシャドウィングされる
		// アドレスも別のものになる
		result := "C"
		fmt.Printf("value: %v, pointer: %p\n", result, &result)
	}
	// ブロックを出た時点で隠れていたresultが現れる
	fmt.Printf("value: %v, pointer: %p\n", result, &result)

	if ok {
		// 以下の定義だと同じ参照
		result = "B"
		fmt.Printf("value: %v, pointer: %p\n", result, &result)
	} else {
		// 以下の定義だと同じ参照
		result = "C"
		fmt.Printf("value: %v, pointer: %p\n", result, &result)
	}
	// ブロックを出た時点で隠れていたresultが現れる
	fmt.Printf("value: %v, pointer: %p\n", result, &result)
}
