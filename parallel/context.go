package parallel

import (
	"context"
	"fmt"
	"time"
)

// 一番の目的はメインゴルーチンからサブゴルーチンを一斉にキャンセルするために使われる
// WithCancel マニュアルでキャンセルを行う
// WithDeadline デッドラインを時刻で指定する
// WithTimeout タイムアウトを指定する

// 第一引数には親のコンテキストを指定
// トップのコンテキストに関しては空のコンテキストを意味するcontext.Backgroundを指定する
// 返り値の(ctx, cancel)はコンテキストとキャンセル用の関数

func Context() {
	// cancel-all
	// var wg sync.WaitGroup
	// ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
	// defer cancel()
	// wg.Add(3)
	// go subTask(ctx, &wg, "a")
	// go subTask(ctx, &wg, "b")
	// go subTask(ctx, &wg, "c")
	// wg.Wait()

	// cascade
	// クリティカルであるサブゴルーチンのみタイムアウトを設定してさらにこのサブゴルーチンがキャンセルされると
	// メインゴルーチンにメッセージを渡して、他の起動しているゴルーチンもキャンセルする

	// var wg sync.WaitGroup
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	v, err := criticalTask(ctx)
	// 	if err != nil {
	// 		fmt.Printf("critical task cancelled due to: %v\n", err)
	// 		cancel()
	// 		return
	// 	}
	// 	fmt.Println("success", v)
	// }()
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	v, err := normalTask(ctx)
	// 	if err != nil {
	// 		fmt.Printf("normal task cancelled due to: %v\n", err)
	// 		return
	// 	}
	// 	fmt.Println("success", v)
	// }()
	// wg.Wait()

	// deadline
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(20*time.Millisecond))
	defer cancel()
	ch := subTask(ctx)
	v, ok := <-ch
	if ok {
		fmt.Println(v)
	}
	fmt.Println("finish")
}

// cancel-all
// func subTask(ctx context.Context, wg *sync.WaitGroup, id string) {
// 	defer wg.Done()
// 	// 指定時間間隔でチャネルへの書き込み信号を生成してくれる
// 	t := time.NewTicker(500 * time.Millisecond)
// 	select {
// 	case <-ctx.Done(): // timeoutするとcontextのDoneチャネルに書き込まれて受信する
// 		fmt.Println(ctx.Err())
// 		return
// 	case <-t.C: // time.NewTickerの読み込み
// 		t.Stop() // 最初の一回でいいため今回はすぐストップ
// 		fmt.Println(id)
// 	}
// }

// cascade
// func criticalTask(ctx context.Context) (string, error) {
// 	ctx, cancel := context.WithTimeout(ctx, 1200*time.Millisecond)
// 	defer cancel()
// 	t := time.NewTicker(1000 * time.Millisecond)
// 	select {
// 	case <-ctx.Done():
// 		return "", ctx.Err()
// 	case <-t.C:
// 		t.Stop()
// 	}
// 	return "A", nil
// }
// func normalTask(ctx context.Context) (string, error) {
// 	t := time.NewTicker(3000 * time.Millisecond)
// 	select {
// 	case <-ctx.Done():
// 		return "", ctx.Err()
// 	case <-t.C:
// 		t.Stop()
// 	}
// 	return "B", nil
// }

// deadline
func subTask(ctx context.Context) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		deadline, ok := ctx.Deadline()
		if ok {
			if deadline.Sub(time.Now().Add(30*time.Millisecond)) < 0 {
				fmt.Println("impossible to meet deadline")
				return
			}
		}
		time.Sleep(30 * time.Millisecond)
		ch <- "hello"
	}()
	return ch
}
