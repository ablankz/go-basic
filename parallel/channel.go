package parallel

import (
	"fmt"
	"sync"
	"time"
)

func Channel() {
	// チャネルの挙動
	// チャネルの書き込みは読み込み待機状態になるまで待機する
	// チャネルの読み込みは書き込みが完了するまで待機する
	// 順番としては先に読み込みのコードまでくる->書き込まれるまで待機->書き込まれる...
	// または先に書き込みのコードまでくる->読み込みのコードがくるまで(読み込み待機中になるまで待つ)->書き込まれるまで待機->書き込まれる...
	// ch := make(chan int)
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	ch <- 10
	// 	time.Sleep(500 * time.Millisecond)
	// }()
	// fmt.Println(<-ch)
	// wg.Wait()

	// goroutine leak
	// // ゴルーチン内で読み込み待機がずっと完了しない状態
	// ch1 := make(chan int)
	// go func() {
	// 	fmt.Println(<-ch1)
	// }()
	// // 以下を追記することでゴルーチンリークが発生しない
	// // ch1 <- 10
	// fmt.Printf("num of working goroutines: %d\n", runtime.NumGoroutine())

	// バッファがある場合はバッファがいっぱい出ない限り、読み込み待機状態になるまで待たなくて良い
	// ただし読み込み側は書き込みがあるまで待機する
	// 更に以下のようにバッファサイズが1のチャネルに対して読み取りがない状態で2つ書き込もうとするとバッファがいっぱいであるため、読み込まれるまでの待機状態となりデッドロックになる
	// ch2 := make(chan int, 1)
	// ch2 <- 2
	// //ch2 <- 3
	// fmt.Println(<-ch2)

	/////////////////////////////////////

	ch1 := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(<-ch1)
	}()
	ch1 <- 10
	// チャネルを閉じる
	close(ch1)
	// okにはチャネルが閉じているかboolで返る
	v, ok := <-ch1
	fmt.Printf("%v %v\n", v, ok)
	wg.Wait()

	ch2 := make(chan int, 2)
	ch2 <- 1
	ch2 <- 2
	// バッファありの場合はcloseを呼び出しても、バッファにデータが溜まっている状態ではチャネルが閉じられない
	// closeをしてから、バッファの値が全て読み込まれたところでチャネルは閉じる
	close(ch2)
	v, ok = <-ch2
	fmt.Printf("%v %v\n", v, ok)
	v, ok = <-ch2
	fmt.Printf("%v %v\n", v, ok)
	v, ok = <-ch2
	fmt.Printf("%v %v\n", v, ok)

	// カプセル化
	ch3 := generateCountStream()
	// 下では一つずつ<-ch3が行われている
	// generateCountStreamのゴルーチン内のdeferでcloseされるまで(0-5まで書き込みが完了するまで)
	// 書き込みが行われると同時にイテレータが一つ進み、読み込まれるとまたgenerateCountStreamの書き込みが行われ...
	for v := range ch3 {
		fmt.Println(v)
	}

	// 通知専用チャネル
	// プロパティのない構造体はサイズが0であるため、通知を行うためだけのチャネルに適している
	nCh := make(chan struct{})
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("goroutine %v started\n", i)
			// 読み込み(通知待ち)
			v, ok := <-nCh
			fmt.Printf("%v: v: %v ok: %v\n", i, v, ok)
			fmt.Println(i)
		}(i)
	}
	time.Sleep(2 * time.Second)
	// バッファなしチャネルなのでcloseした時点でチャネルが閉じられる->一斉に待機中の読み込みが進む(渡る値は ({}, false) となる)
	close(nCh)
	fmt.Println("unblocked by manual close")

	wg.Wait()
	fmt.Println("finish")
}

// <-chanにすることで読み取り専用のチャネルを返す
func generateCountStream() <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i <= 5; i++ {
			ch <- i
		}
	}()
	return ch
}
