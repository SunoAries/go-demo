package main

import "fmt"

func main() {
	var a []int
	b := make([]int, 5)
	c := b[0:1]
	d := c[1:4]
	fmt.Println(a, b, c, d)

	arr := [5]int{1, 2, 3}
	arr2 := arr
	arr2[0] = 10
	fmt.Println(arr, arr2)
	feibo(40)
	fmt.Println("2:")
	fmt.Println(febo2(40))
}

func feibo(n int) {
	var a []int
	a = make([]int, n)

	a[0] = 0
	a[1] = 1
	for i := 2; i < n; i++ {
		a[i] = a[i-1] + a[i-2]
	}
	for _, v := range a {
		fmt.Println(v)

	}
}

func febo2(n int) int {
	if n<0 {
		return 0
	}
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return febo2(n-1) + febo2(n-2)
}

