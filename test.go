package basic

import (
	"fmt"
)

func main() {
	//fmt.Println("My favorite number is", rand.Intn(10))
	//var arr = [5]int{1,
	//	23, 234, 2123, 0}
	//slice(arr)
	rangetest()
}

func slice(arr [5]int) {
	var shit = arr[2:]
	println(shit[0])
}

func rangetest() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, _ := range nums {

		fmt.Println("index", i)

	}
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
