package main

import (
	"net"
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main() {
	conn, err:= net.Dial("tcp","localhost:3000")
	if err!=nil {
		fmt.Println("err dialing:",err.Error())
		return
	}
	status, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(conn,"adsf",status)

	defer conn.Close()

	inputReader := bufio.NewReader(os.Stdin)

	for  {
		input,_:=inputReader.ReadString('\n')
		trimedIput := strings.Trim(input,"\r\n")
		if trimedIput=="Q" {
			return
		}
		_,err:=conn.Write([]byte(trimedIput))
		if err!=nil {
			fmt.Println("err conn.write",err)
			return
		}
	}
}
