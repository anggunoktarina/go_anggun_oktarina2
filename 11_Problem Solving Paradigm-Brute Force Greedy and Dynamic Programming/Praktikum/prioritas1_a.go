package main

import "fmt"

func biner(n int) {
	for i := 0; i <= n; i++ {
		fmt.Printf("%b", i)
	}
}

func main() {
	biner(3)
	biner(2)
}
