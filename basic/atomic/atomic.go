package main

import (
	"fmt"
	"sync"
	"time"
)

type atomicInt struct {
	value int
	lock  sync.Mutex
}

func (a *atomicInt) increment() {
	fmt.Printf("safe increment: %d\n", a.value)
	func() {
		a.lock.Lock()
		defer a.lock.Unlock()

		a.value++
		println(a.value)
	}()
}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()

	return a.value
}

func main() {
	var a atomicInt
	a.increment()
	go func() {
		fmt.Println("并发 1 ：")
		a.increment()
		fmt.Println("并发结束1")
	}()
	go func() {
		fmt.Println("并发 2 ：")
		a.increment()
		fmt.Println("并发结束2")
	}()
	go func() {
		fmt.Println("并发 3 ：")
		a.increment()
		fmt.Println("并发结束3")
	}()
	go func() {
		fmt.Println("并发 4 ：")
		a.increment()
		fmt.Println("并发结束4")
	}()

	// for i := 0; true; i++ {
	// 	println(i)
	// 	time.Sleep(1000 * time.Millisecond)
	// }
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(a.get())
}
