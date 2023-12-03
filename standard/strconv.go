package standard

import (
	"fmt"
	"strconv"
)

// 文字列変換

func Strconv() {
	bt := true
	fmt.Printf("%[1]T %[1]v\n", strconv.FormatBool(bt))

	// 10進数変換
	i := strconv.FormatInt(-100, 10)
	fmt.Printf("%[1]v, %[1]T\n", i)

	// 上記と同じ
	i2 := strconv.Itoa(100)
	fmt.Printf("%[1]v, %[1]T\n", i2)

	// 第二引数はverbと同じ
	// 第三引数は桁数の制限(-1なら自動で決定)
	fmt.Println(strconv.FormatFloat(123.456, 'E', -1, 64))
	fmt.Println(strconv.FormatFloat(123.456, 'e', 2, 64))
	fmt.Println(strconv.FormatFloat(123.456, 'f', -1, 64))
	fmt.Println(strconv.FormatFloat(123.456, 'f', 2, 64))
	fmt.Println(strconv.FormatFloat(123.456, 'g', -1, 64))
	fmt.Println(strconv.FormatFloat(123456789.123, 'f', -1, 64))
	fmt.Println(strconv.FormatFloat(123.456, 'g', 4, 64))
	fmt.Println(strconv.FormatFloat(123456789.123, 'G', 8, 64))

	bt1, _ := strconv.ParseBool("true")
	fmt.Printf("%v, %T\n", bt1, bt1)
	bt2, _ := strconv.ParseBool("1")
	bt3, _ := strconv.ParseBool("t")
	bt4, _ := strconv.ParseBool("T")
	bt5, _ := strconv.ParseBool("TRUE")
	bt6, _ := strconv.ParseBool("True")
	fmt.Println(bt2, bt3, bt4, bt5, bt6)

	bf1, ok := strconv.ParseBool("false")
	if ok != nil {
		fmt.Println("Convert Error")
	}
	fmt.Printf("%v, %T\n", bf1, bf1)
	bf2, _ := strconv.ParseBool("0")
	bf3, _ := strconv.ParseBool("f")
	bf4, _ := strconv.ParseBool("F")
	bf5, _ := strconv.ParseBool("FALSE")
	bf6, _ := strconv.ParseBool("False")
	fmt.Println(bf2, bf3, bf4, bf5, bf6)

	// 第二引数は基数(10進数など)
	// 第三引数は精度(0の場合、goのintの精度)
	i3, _ := strconv.ParseInt("12345", 10, 0)
	fmt.Printf("%v, %T\n", i3, i3)
	i4, _ := strconv.ParseInt("-1", 10, 0)
	fmt.Printf("%v, %T\n", i4, i4)

	// 10進数のintならこっちが簡単
	i10, _ := strconv.Atoi("123")
	fmt.Println(i10)

	fl1, _ := strconv.ParseFloat("3.14", 64)
	fl2, _ := strconv.ParseFloat(".2", 64)
	fl3, _ := strconv.ParseFloat("-2", 64)
	fl4, _ := strconv.ParseFloat("1.2345e8", 64)
	fl5, _ := strconv.ParseFloat("1.2345E8", 64)
	fmt.Println(fl1, fl2, fl3, fl4, fl5)
}
