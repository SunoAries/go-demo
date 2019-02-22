package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println("input:")
		if input.Text() == "exit" {
			break
		}
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 0 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
