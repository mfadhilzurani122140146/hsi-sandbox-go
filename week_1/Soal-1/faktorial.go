package main

import (
	"fmt"
)

func factorial(n int) int {
	if n == 0 || n == 1 { // please check negative number too
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	var n int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&n)

	if n < 0 {
		fmt.Println("Faktorial hanya untuk bilangan non-negatif.")
	} else {
		fmt.Printf("Faktorial dari %d adalah %d\n", n, factorial(n))
	}
}
