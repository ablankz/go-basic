package basic

import (
	"fmt"
	"unsafe"
)

type Task struct {
	Title    string
	Estimate int
}

func Struct() {
	// 構造体の初期化
	task1 := Task{
		Title:    "Learn GoLang",
		Estimate: 3,
	}
	// 構造体のフィールド値変更
	task1.Title = "Learning Go"
	// %に+をつけることでフィールド名も記述してくれる
	fmt.Printf("%[1]T %+[1]v %v\n", task1, task1.Title)

	var task2 = task1
	task2.Title = "new"
	// 実体は別のメモリ領域に生成されることがわかった
	fmt.Printf("task1: %v task2: %v\n", task1.Title, task2.Title)

	task1p := &Task{
		Title:    "Learn concurrency",
		Estimate: 2,
	}
	fmt.Printf("task1p: %T %+v %v\n", task1p, *task1p, unsafe.Sizeof(task1p))
	// (*task1p).Title = "Changed"
	// ↑を以下のようにかける
	task1p.Title = "Changed"
	fmt.Printf("task1p: %+v\n", *task1p)
	var task2p *Task = task1p
	task2p.Title = "Chaged by Task2"
	// 参照越しに変更すると元の方も書き換わる
	fmt.Printf("task1: %+v\n", *task1p)
	fmt.Printf("task2: %+v\n", *task2p)

	// (&task1).extendEstimatePointer()
	// ↑を以下のようにかける
	task1.extendEstimatePointer()
	fmt.Printf("task2 value receiver: %+v\n", task1.Estimate)
}

// Taskのreceiver作成
// 以下の値レシーバだとインスタンスに変更を加えられない
// func (task Task) extendEstimate() {
// 	task.Estimate += 10
// }

// 以下のポインタレシーバだとインスタンスに変更を加えられない
func (task *Task) extendEstimatePointer() {
	task.Estimate += 10
}
