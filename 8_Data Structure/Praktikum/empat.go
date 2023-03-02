package main

import "fmt"

func PairSum(arr []int, target int) []int {
	// Inisialisasi indeks awal dan akhir
	i, j := 0, len(arr)-1

	// Melakukan iterasi sampai ditemukan pasangan angka
	for i < j {
		sum := arr[i] + arr[j]
		if sum == target {
			// Jika ditemukan pasangan, kembalikan indeksnya
			return []int{i, j}
		} else if sum < target {
			// Jika jumlah pasangan kurang dari target, naikkan indeks pertama
			i++
		} else {
			// Jika jumlah pasangan lebih dari target, turunkan indeks kedua
			j--
		}
	}

	// Jika tidak ditemukan pasangan, kembalikan array kosong
	return []int{}
}

func main() {
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6)) // [1,3]
	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11))  // [0, 2]
	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12))   // [2, 3]
	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10))   // [1, 2]
	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6))    // [0, 1]
}
