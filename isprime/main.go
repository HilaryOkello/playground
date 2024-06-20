package main

import "fmt"

func main() {
	fmt.Println(isprime(0))
	fmt.Println(nextprime(24))
}

func isprime(n int) bool {
	if n < 2 {
		return false
	} else if n == 2 || n == 3 {
		return true
	} else if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func prevprime(num int) int {
	if num < 2 {
		return 0
	}
	for {
		if isprime(num) {
			return num
		}
		num = num - 1
	}
}

func nextprime(num int) int {
	if num < 2 {
		return 0
	}
	for {
		if isprime(num) {
			return num
		}
		num = num + 1
	}
}
