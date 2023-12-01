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
	basic.Shadowing()

}
