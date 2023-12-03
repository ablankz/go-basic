package standard

import (
	"fmt"
	"os"
)

func Fmt() {
	fmt.Println("表示")

	// ビルトイン関数
	print("Hello")
	println("Hello!")

	fmt.Print("Hello")
	fmt.Println("Hello!")

	// 半角スペースで区切られる
	fmt.Println("Hello", "world!")
	fmt.Println("Hello", "world!")

	// sは文字列
	fmt.Printf("%s\n", "Hello")
	// #で値をそのまま "Hello"
	fmt.Printf("%#v\n", "Hello")

	s1 := fmt.Sprint(12)
	s2 := fmt.Sprintf("Hello %v", 1)
	s3 := fmt.Sprintln("Hello")

	fmt.Print(s1 + s2 + s3)

	// 書き込み先指定, 任意のIoWriter型への出力関数
	fmt.Fprint(os.Stdout, "Hello")
	fmt.Fprintf(os.Stdout, "Hello")
	fmt.Fprintln(os.Stdout, "Hello")

	f, _ := os.Create("test.txt")
	defer f.Close()

	fmt.Fprintln(f, "FPrintln")

}
