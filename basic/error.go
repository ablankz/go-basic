package basic

import (
	"errors"
	"fmt"
	"os"
)

// センチネルルエラーの作成
// 慣習としてErrから始まる
var ErrCutom = errors.New("not found")

func Error() {
	err01 := errors.New("something wrong")
	err02 := errors.New("something wrong")
	fmt.Printf("%[1]p %[1]T %[1]v\n", err01)
	// 以下のエラーメソッドのある構造体はerrorインタフェースを実装していることになる
	fmt.Println(err01.Error())
	// errはポインタであるため
	fmt.Println(err01 == err02)

	// エラーオブジェクトのラップ
	// これはmessageとerrorをプロパティに持つ
	err0 := fmt.Errorf("add info: %w", errors.New("original error"))
	fmt.Printf("%[1]p %[1]T %[1]v\n", err0)
	// 以下のメソッドでerrorプロパティをとってくる
	fmt.Println(errors.Unwrap(err0))
	fmt.Printf("%T\n", errors.Unwrap(err0))
	// %vを使ってラップエラーを作成(unwrapメソッドを実装していない)
	err1 := fmt.Errorf("add info: %v", errors.New("original error"))
	fmt.Println(err1)
	fmt.Printf("%T\n", err1)
	// unwrapメソッドを実装していないためnilを返す
	fmt.Println(errors.Unwrap(err1))

	err2 := fmt.Errorf("in repository layer: %w", ErrCutom)
	fmt.Println(err2)
	// さらにラップすると付加情報が累積
	err2 = fmt.Errorf("in repository layer: %w", err2)
	fmt.Println(err2)

	// wrapを取りながらその都度第２引数に一致するか見てくれる
	if errors.Is(err2, ErrCutom) {
		fmt.Println("matched")
	}

	file := "dummy.txt"
	err3 := fileCheck(file)
	if err3 != nil {
		if errors.Is(err3, os.ErrNotExist) {
			fmt.Printf("%v file not found\n", file)
		} else {
			fmt.Printf("unknown error")
		}
	}
}

func fileCheck(name string) error {
	f, err := os.Open(name)
	// ファイルがなければ
	if err != nil {
		return fmt.Errorf("in checker: %w", err)
	}
	// 関数を出るときにファイルを閉じる
	defer f.Close()
	return nil
}
