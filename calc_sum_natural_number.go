package main

import "fmt"

func main() {
	var n, sum int
	fmt.Println("Enter a positive number:")

	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		sum += i
	}

	fmt.Println("Sum: ", sum)
}
