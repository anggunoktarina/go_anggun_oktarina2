package main

import (
	"fmt"
)

func main() {
	var A, B, C int
	fmt.Print("Enter A, B, and C:")
	fmt.Scan(&A, &B, &C)

	for z := 1; z <= 10000; z++ {
		for x := 1; x <= 10000; x++ {
			y := A - x - z
			if x*x+y*y+z*z == C && x*y*z == B {
				fmt.Printf("x = %d, y = %d, z = %d\n", x, y, z)
				return
			}
		}
	}
	fmt.Println("No solution found.")
}
