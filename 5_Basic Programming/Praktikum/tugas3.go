package main

import "fmt"

func main() {
	//deklarasi variabel nilai
	var tugas, uts, uas, nilai int32

	//meminta input nilai dari user
	fmt.Printf("masukan nilai tugas :")
	fmt.Scanf("%d\n", &tugas)
	fmt.Printf("masukan nilai uts :")
	fmt.Scanf("%d\n", &uts)
	fmt.Printf("masukan nilai uas :")
	fmt.Scanf("%d\n", &uas)
	nilai = (tugas + uts + uas) / 3
	//menentukan grade berdasarkan nilai yang dimasukkan
	if nilai >= 90 {
		fmt.Printf("nilai keseluruhan anda =%d\n", nilai)
		fmt.Println("grade anda adalah A")
	} else if nilai >= 80 {
		fmt.Printf("nilai keseluruhan anda =%d\n", nilai)
		fmt.Println("grade anda adalah B")
	} else if nilai >= 70 {
		fmt.Printf("nilai keseluruhan anda =%d\n", nilai)
		fmt.Println("grade anda adalah C")
	} else if nilai >= 60 {
		fmt.Print("nilai keseluruhan anda =%d\n", nilai)
		fmt.Println("grade anda adalah D")
	} else {
		fmt.Printf("nilai keseluruhan anda =%d\n", nilai)
		fmt.Println("grade anda adalah E")
	}
}
