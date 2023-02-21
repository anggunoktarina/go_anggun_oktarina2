package main

import "fmt"

func main() {
	//pendeklarasian tipe data dan variable
	var bilangan int
	//deklarasi input data
	fmt.Print("masukkan bilangan:")
	fmt.Scanln(&bilangan)
	//kondisi if else
	if bilangan%2 == 0 {
		fmt.Println("bilangan genap")
	} else {
		fmt.Println("bilangan ganjil")
	}
}
