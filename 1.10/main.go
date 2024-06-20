package main

import (
	"fmt"
)

func main() {
	var n, c, d int
	_, _ = fmt.Scan(&n, &c, &d)
	findFirst(n, c, d)
}

func findFirst(n int, c int, d int) {
	firstNumber := 0
	for i := 1; i <= n; i++ {
		if i%c == 0 && i%d != 0 {
			firstNumber = i
			break
		} else {
			continue
		}
	}
	if firstNumber == 0 {
		return
	} else {
		fmt.Println(firstNumber)
	}
}
