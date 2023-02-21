package main

import "fmt"

func main() {
	var rows int
	fmt.Print("masukan jumlah baris:")
	fmt.Scan(&rows)

	for i := 1; i <= rows; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
