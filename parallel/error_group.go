package parallel

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func ErrorGroup() {
	// manager
	// eg := new(errgroup.Group)
	// s := []string{"task1", "fake1", "task2", "fake2"}
	// for _, v := range s {
	// 	task := v
	// 	eg.Go(func() error {
	// 		return doTask(task)
	// 	})
	// }
	// // nilでないエラーが一つでもある場合、そのうち、最も早くエラーを投げたものがerrに入る
	// if err := eg.Wait(); err != nil {
	// 	fmt.Printf("error :%v\n", err)
	// }
	// fmt.Println("finish")

	// context
	// eg, ctx := errgroup.WithContext(context.Background())
	// s := []string{"task1", "fake1", "task2", "fake2", "task3"}
	// for _, v := range s {
	// 	task := v
	// 	eg.Go(func() error {
	// 		return doTask(ctx, task)
	// 	})
	// }
	// if err := eg.Wait(); err != nil {
	// 	fmt.Printf("error :%v\n", err)
	// }
	// fmt.Println("finish")

	// timeout
	ctx, cancel := context.WithTimeout(context.Background(), 800*time.Millisecond)
	defer cancel()
	eg, ctx := errgroup.WithContext(ctx)
	s := []string{"task1", "task2", "task3", "task4"}
	for _, v := range s {
		task := v
		eg.Go(func() error {
			return doTask(ctx, task)
		})
	}
	if err := eg.Wait(); err != nil {
		fmt.Printf("error :%v\n", err)
	}
	fmt.Println("finish")

}

// manager
// func doTask(task string) error {
// 	if task == "fake1" || task == "fake2" {
// 		return fmt.Errorf("%v failed", task)
// 	}
// 	fmt.Printf("task %v completed\n", task)
// 	return nil
// }

// context
// func doTask(ctx context.Context, task string) error {
// 	var t *time.Ticker
// 	switch task {
// 	case "fake1":
// 		t = time.NewTicker(500 * time.Millisecond)
// 	case "fake2":
// 		t = time.NewTicker(700 * time.Millisecond)
// 	default:
// 		t = time.NewTicker(200 * time.Millisecond)
// 	}
// 	select {
// 	case <-ctx.Done():
// 		fmt.Printf("%v cancelled : %v\n", task, ctx.Err())
// 		return ctx.Err()
// 	case <-t.C:
// 		t.Stop()
// 		if task == "fake1" || task == "fake2" {
// 			return fmt.Errorf("%v process failed", task)
// 		}
// 		fmt.Printf("task %v completed\n", task)
// 	}
// 	return nil
// }

// timeout
func doTask(ctx context.Context, task string) error {
	var t *time.Ticker
	switch task {
	case "task1":
		t = time.NewTicker(500 * time.Millisecond)
	case "task2":
		t = time.NewTicker(700 * time.Millisecond)
	default:
		t = time.NewTicker(1000 * time.Millisecond)
	}
	select {
	case <-ctx.Done():
		fmt.Printf("%v cancelled : %v\n", task, ctx.Err())
		return ctx.Err()
	case <-t.C:
		t.Stop()
		fmt.Printf("task %v completed\n", task)
	}
	return nil
}
