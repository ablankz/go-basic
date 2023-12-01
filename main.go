package main

import (
	"fmt"
	"go-basic/basic"
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
	basic.UnitTestSrc()
}
