package parallel

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
	"time"
)

func Goroutine() {
	// ゴルーチンの起動に時間がかかるため呼ばれずに終わる
	// go func() {
	// 	fmt.Println("goroutine invoked")
	// }()
	// // mainもmainゴルーチンというゴルーチン
	// fmt.Printf("num of working goroutines: %d\n", runtime.NumGoroutine())
	// fmt.Println("main func finish")

	// var wg sync.WaitGroup
	// // waitグループに1追加
	// wg.Add(1)
	// go func() {
	// 	// 終了時にwaitグループを1減らす
	// 	defer wg.Done()
	// 	fmt.Println("goroutine invoked")
	// }()
	// // waitグループが0になるまで待つ
	// wg.Wait()
	// fmt.Printf("num of working goroutines: %d\n", runtime.NumGoroutine())
	// fmt.Println("main func finish")

	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalln("Error:", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalln("Error:", err)
		}
	}()
	if err := trace.Start(f); err != nil {
		log.Fatalln("Error:", err)
	}
	defer trace.Stop()
	// go tool trace trace.out で確認可能
	ctx, t := trace.NewTask(context.Background(), "main")
	defer t.End()
	fmt.Println("The number of logical CPU Cores:", runtime.NumCPU())

	// 逐次処理
	// task(ctx, "Task1")
	// task(ctx, "Task2")
	// task(ctx, "Task3")

	// 並行処理
	var wg sync.WaitGroup
	wg.Add(3)
	go cTask(ctx, &wg, "Task1")
	go cTask(ctx, &wg, "Task2")
	go cTask(ctx, &wg, "Task3")
	wg.Wait()

	s := []int{1, 2, 3}
	for _, i := range s {
		wg.Add(1)
		// 次の形だと実行されるタイミングのiの値となる
		// go func() {
		// 	defer wg.Done()
		// 	fmt.Println(i)
		// }()
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()

	fmt.Println("main func finish")
}

//	func task(ctx context.Context, name string) {
//		// 下のdeferで実行されるのはendのみ、チェーンでつなげた場合はチェーンの最後のみクリーンアップで実行される
//		defer trace.StartRegion(ctx, name).End()
//		time.Sleep(time.Second)
//		fmt.Println(name)
//	}

func cTask(ctx context.Context, wg *sync.WaitGroup, name string) {
	defer trace.StartRegion(ctx, name).End()
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Println(name)
}
