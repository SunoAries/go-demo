package basic

import (
	"fmt"
	"unsafe"
)

var s int

func cal() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str" //多重赋值

	area = LENGTH * WIDTH
	fmt.Printf("面积为 : %d", area)
	println()
	println(a, b, c)
}

func main() {
	forrange()
}

func forrange()  {
	for pos, char := range "日本\x80語" { // \x80 是个非法的UTF-8编码
		fmt.Printf("字符 %#U 始于字节位置 %d\n", char, pos)
	}
}

//func readFile(name string) error {
//	f, err := os.Open(name)
//	if err != nil {
//		return err
//	}
//	d, err := f.Stat()
//	if err != nil {
//		f.Close()
//		return err
//	}
//}

func unsafetest() {
	const (
		a = "abc"
		b = len(a)
		c = unsafe.Sizeof(a)
	)
	println(a, b, c)
}

func assignmenttest() {
	var a int = 10
	var b = 10
	c := 10
	g, h := 123, "hello"
	iota_test()
	fmt.Println(s, a, b, c, g, h)
}

func iota_test() {
	const (
		a = iota
		b
		c
	)
	println(a, b, c)
}

func wfor(bool2 bool) {
	var i = 0
	for bool2 {
		i++
		fmt.Println(i)
	}
}

func add(num1, num2 uint) uint {
	return num1 + num2
}
