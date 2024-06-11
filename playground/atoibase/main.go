package main

import (
	"fmt"
)

func anyBaseToDecimal(number string, base string) int {
	var result int
	var power int
	var digitIndex int
	for i := len(number) - 1; i >= 0; i-- {
		digit := number[i]
		track := false
		for index, char := range base {
			if char == rune(digit) {
				digitIndex = index
				track = true
				break
			}
		}
		if !track {
			fmt.Println(string(number[i]), "not in", base)
			return 0
		}
		result += digitIndex * Power(len(base), power)
		power++
	}
	return result
}

func Power(num, n int) int {
	if n == 0 {
		return 1
	}
	return num * Power(num, n-1)
}

func main() {
	number :="10"
	base := "01"
	fmt.Println("Decimal value:", anyBaseToDecimal(number, base))
}
