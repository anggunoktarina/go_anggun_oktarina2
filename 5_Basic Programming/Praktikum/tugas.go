package main

import "fmt"

func main() {
	var nama string
	var alas, tinggi, luas float32

	fmt.Println("====PROGRAM TRAPESIUM====")
	fmt.Printf("masukan nama : ")
	fmt.Scanf("%s\n", &nama)
	fmt.Print("masukan alas trapesium: ")
	fmt.Scanln(&alas)
	fmt.Print("masukan tinggi trapesium: ")
	fmt.Scanln(&tinggi)
	//statements
	luas = 0.5 * (alas + tinggi)
	//output
	fmt.Println("=============================")
	fmt.Printf("nama anda %v:\n", nama)
	fmt.Println("luas trapesium adalah", luas)

}
