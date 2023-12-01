package basic

import (
	"fmt"
	"time"
)

func Control() {
	fmt.Println("---------------------if---------------------")
	a := 1
	if a == 0 {
		fmt.Println("zero")
	} else if a > 0 {
		fmt.Println("positive")
	} else {
		fmt.Println("negative")
	}
	fmt.Println("---------------------for--------------------")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	// 省略することで無限ループ
	// for {
	// 	fmt.Println("working")
	// 	time.Sleep(2 * time.Second)
	// }
	var i int
	for {
		if i > 3 {
			break
		}
		fmt.Println(i)
		i += 1
		time.Sleep(300 * time.Millisecond)
	}
	// ラベルをつける
loop:
	for i := 0; i < 10; i++ {
		switch i {
		case 2:
			continue
		case 3:
			continue
		case 8:
			// ただのbreakだとswitch文から抜けるためラベルを指定してbreak
			break loop
		default:
			fmt.Printf("%v ", i)
		}
	}
	var items = []item{{price: 10}, {price: 20}, {price: 30}}

	// 以下の方法だと値が書き換わらない
	for _, i := range items {
		i.price *= 1.1
	}
	fmt.Printf("\n%+v\n", items)
	// 以下の方法だと値が書き換わる
	for i := range items {
		items[i].price *= 1.1
	}
	fmt.Printf("%+v\n", items)
}

type item struct {
	price float32
}
