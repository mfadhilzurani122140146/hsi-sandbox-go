package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Jumlah elemen harus lebih dari 0.")
		return
	}

	numbers := make([]int, n)

	fmt.Println("Masukkan elemen-elemen slice:")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemen ke-%d: ", i+1)
		fmt.Scan(&numbers[i])
	}

	max := numbers[0]
	for _, val := range numbers {
		if val > max {
			max = val
		}
	}

	fmt.Printf("Output: %d (karena %d adalah bilangan terbesar dalam slice tersebut)\n", max, max)
}
