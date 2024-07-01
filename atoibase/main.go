package main

import "fmt"

func main() {
	fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("1111101", "01"))
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", "-ab"))
}

func AtoiBase(s string, base string) int {
	if r := validateBase(base); r < 1 {
		return r
	}
	var result int
	var digitIndex int
	var p int
	for i := len(s) - 1; i >= 0; i-- {
		digit := s[i]
		track := false
		for index, char := range base {
			if char == rune(digit) {
				track = true
				digitIndex = index
				break
			}
		}
		if !track {
			return 0
		}
		result += digitIndex * power(len(base), p)
		p++
	}
	return result
}

func power(num, p int) int {
	if p == 0 {
		return 1
	}
	return num * power(num, p-1)
}

func validateBase(base string) int {
	if len(base) < 2 {
		return 0
	}
	baseMap := make(map[rune]bool)
	for _, char := range base {
		if char == '-' || char == '+' {
			return 0
		}
		if baseMap[char] {
			return 0
		}
		baseMap[char] = true
	}
	return 1
}
