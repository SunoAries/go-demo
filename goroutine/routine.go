package main

import (
	"fmt"
	"sync"
	"time"
)

func calcz(w *sync.WaitGroup, i int) {
	fmt.Println("calc: ", i)
	time.Sleep(time.Second)
	w.Done()
}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go calcz(&wg, i)
	}
	wg.Wait()
	fmt.Println("all goroutine finish")
}
