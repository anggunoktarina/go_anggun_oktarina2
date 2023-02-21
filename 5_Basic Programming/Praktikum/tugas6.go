package main

import "fmt"

func main() {
	var num int
	fmt.Print("masukkan bilangan:")
	fmt.Scan(&num)

	fmt.Print("faktor dari", num, "adalah:\n")

	for i := 1; i <= num; i++ {
		if num%i == 0 {
			fmt.Printf("%v\n", i)
		}
	}
}
