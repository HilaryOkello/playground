package main

import (
	"fmt"
	"os"
)

func main() {
	numString := os.Args[1]
	num := atoi(numString)
	var result string
	digits := "01"
	base := 2
	if num == 0 {
		result = string(digits[0])
	}
	for num > 0 {
		remainder := num % base
		result = string(digits[remainder]) + result
		num /= base
	}

	fmt.Println(result)
}

func atoi(s string) int {
	var result int
	sign := 1
	if s[0] == '-' || s[0] == '+' {
		if s[0] == '-' {
			sign = -1
		}
		s = s[1:]
	}
	for _, char := range s {
		if char < '0' || char > '9' {
			return 0
		} else {
			result = result*10 + int(char-'0')
		}
	}
	return result * sign
}
