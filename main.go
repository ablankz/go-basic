package main

import (
	"fmt"
	"go-basic/standard"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	fmt.Println(os.Getenv("GO_ENV"))
	// fmt.Println(calculator.Offset)
	// fmt.Println(calculator.Sum(1, 2))
	// fmt.Println(calculator.Multiply(2, 3))

	// basic.Variables()
	// basic.Pointer()
	// basic.Shadowing()
	// basic.Slice()
	// basic.Map()
	// basic.Struct()
	// basic.Function()
	// basic.Interface()
	// basic.Control()
	// basic.Error()
	// basic.Generics()
	// basic.UnitTestSrc()
	// basic.Logger()
	// parallel.Goroutine()
	// parallel.Channel()
	// parallel.Select()
	// parallel.Mutex()
	// parallel.Context()
	// parallel.ErrorGroup()
	// parallel.Pipeline()
	// parallel.FanOutFanIn()
	// parallel.Heartbeat()

	// standard.Time()
	// standard.Math()
	// standard.Flag()
	// standard.Fmt()
	// standard.Log()
	// standard.Strconv()
	// standard.Strings()
	// standard.Bufio()
	// standard.Regexp()
	// standard.Crypt()
	// standard.Json()
	// standard.Sort()
	// standard.Context()
	standard.NetUrl()
}
