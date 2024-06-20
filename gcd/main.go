package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 || len(os.Args) > 3 {
		return
	}
	a, _ := strconv.Atoi(os.Args[1])
	b, _ := strconv.Atoi(os.Args[2])

	if a < b {
		a, b = b, a
	}
	for b != 0 {
		a, b = b, a%b
	}
	fmt.Println(a)
}
