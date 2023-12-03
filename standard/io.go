package standard

import (
	"fmt"
	"io"
	"os"
)

// osなどと組み合わせて入出力を行う

func Io() {
	f, _ := os.Open("foo.txt")
	bs1, _ := io.ReadAll(f)
	fmt.Println(string(bs1))
}
