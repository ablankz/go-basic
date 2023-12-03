package standard

import (
	"bufio"
	"fmt"
	"os"
)

// 標準入力を行単位で読み込む

func Bufio() {
	// 標準入力をソースにしたスキャナの生成
	scanner := bufio.NewScanner(os.Stdin)

	// 入力のスキャナが成功する際、繰り返すループ
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "読み込みエラー", err)
	}

}
