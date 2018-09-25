package main

import (
	"fmt"
	"time"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

var (
	domainSyncChan = make(chan int, 10)
)

func domainPut(num int) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("error to chan put.")
		}
	}()
	fmt.Println(domainSyncChan)
	domainSyncChan <- num
	fmt.Println(domainSyncChan)

}

func main() {
	for i := 0; i < 10; i++ {
		domainName := i
		go domainPut(domainName)
	}
	time.Sleep(time.Second * 2)
}
