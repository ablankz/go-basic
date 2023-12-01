package parallel

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func FanOutFanIn() {
	cores := runtime.NumCPU()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	nums := []int{1, 2, 3, 4, 5, 7, 8}

	outChs := make([]<-chan string, cores)
	inData := fanGenerator(ctx, nums...)
	for i := 0; i < cores; i++ {
		outChs[i] = fanOut(ctx, inData, i+1)
	}
	// ここからデータを結合させていく(FanIn)
	var i int
	flag := true
	for v := range fanIn(ctx, outChs...) {
		if i == 3 {
			cancel()
			flag = false
		}
		if flag {
			fmt.Println(v)
		}
		i++
	}
	fmt.Println("finish")
}
func fanGenerator(ctx context.Context, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case <-ctx.Done():
				return
			case out <- n:
			}
		}
	}()
	return out
}
func fanOut(ctx context.Context, in <-chan int, id int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		heavyWork := func(i int, id int) string {
			time.Sleep(200 * time.Millisecond)
			return fmt.Sprintf("result:%v (id:%v)", i*i, id)
		}
		for v := range in { // 重い処理が終わるまで次の値を読み込めない->コア数分この関数を繰り返し呼ぶ
			select {
			case <-ctx.Done():
				return
			case out <- heavyWork(v, id): // 時間が長い処理
			}
		}
	}()
	return out
}
func fanIn(ctx context.Context, chs ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	out := make(chan string)
	// コア一個分の処理
	multiplex := func(ch <-chan string) {
		defer wg.Done()
		// 1コアで2つ分処理したコアがあれば下のforでは2回回る
		// (この関数を呼び出した段階では何回回るかはわからない)
		for text := range ch {
			select {
			case <-ctx.Done():
				return
			case out <- text:
			}
		}
	}
	wg.Add(len(chs))
	for _, ch := range chs {
		go multiplex(ch)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
