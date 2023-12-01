package basic

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func funcDefer() {
	// deferをつけるとこの関数を終了時に呼ばれる
	// 呼び出される順番はdeferのついているうちの下から順次
	defer fmt.Println("main func final-finish")
	defer fmt.Println("main func semi-finish")
	fmt.Println("hello world")
}

// 可変長引数
func trimExtension(files ...string) []string {
	// ファイル数分のキャパシティを持つスライス作成
	out := make([]string, 0, len(files))
	for _, f := range files {
		out = append(out, strings.TrimSuffix(f, ".csv"))
	}
	return out
}

func fileChecker(name string) (string, error) {
	f, err := os.Open(name)
	// ファイルがなければ
	if err != nil {
		return "", errors.New("file not found")
	}
	// 関数を出るときにファイルを閉じる
	defer f.Close()
	return name, nil
}

// 無名関数を引数に取る関数
func addExt(f func(string) string, name string) {
	fmt.Println(f(name))
}

// 無名関数を返す関数
func multiply() func(int) int {
	return func(n int) int {
		return n * 1000
	}
}

func countUp() func(int) int {
	// 関数内で閉じた変数
	count := 0
	// 関数を通して操作することで不正なアクセスなどを防ぐ
	return func(n int) int {
		count += n
		return count
	}
}

func Function() {
	funcDefer()
	files := []string{"file1.csv", "file2.csv", "file3.csv"}
	// fileを展開して渡す
	fmt.Println(trimExtension(files...))
	name, err := fileChecker("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(name)

	i := 1
	// 無名関数(即実行)
	func(i int) {
		fmt.Println(i)
	}(i)
	// 無名関数(あとから実行)
	f1 := func(i int) int {
		return i + 1
	}
	fmt.Println(f1(i))

	f2 := func(file string) string {
		return file + ".csv"
	}
	addExt(f2, "file1")

	f3 := multiply()
	fmt.Println(f3(2))

	f4 := countUp()
	for i := 1; i <= 5; i++ {
		// 呼ばれるたびに内部のcountが足されていき、それが返る
		v := f4(2)
		fmt.Printf("%v\n", v)
	}
}
