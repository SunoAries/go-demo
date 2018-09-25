package main

import (
	"fmt"
	"sort"
)

func main() {
	var a map[string]string
	a = make(map[string]string, 5)
	b := make(map[string]string)
	a["name"] = "Nick" // "create"
	fmt.Println(a, b)
	delete(a, "name") // "delete"
	fmt.Println(a, b)
	a["name"] = "Dawn"   // "update"
	name, _ := a["name"] // "read"
	fmt.Println(a, name)

	testMap3()

	testMapSort()
}

func modify(a map[string]map[string]string) {
	_, ok := a["name"]
	if !ok {
		a["name"] = make(map[string]string)
	}
	a["name"]["Nick"] = "suoning"
	a["name"]["Nicky"] = "manyRou"
}

func testMap3() {
	var a map[string]map[string]string
	a = make(map[string]map[string]string, 10) //初始化一维
	a["name"] = make(map[string]string)        //初始化二维

	modify(a)
	fmt.Println(a)
}

func testMapSort() {
	a := make(map[int]int, 5)

	a[8] = 10
	a[5] = 10
	a[2] = 10
	a[1] = 10
	a[9] = 10
	fmt.Println(a)
	var keys []int
	for k := range a {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	fmt.Println(a)
}
