package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	// Buat map untuk menyimpan setiap elemen yang muncul
	uniqueElements := make(map[string]bool)

	// Iterate setiap elemen di array A dan B
	for _, elem := range arrayA {
		uniqueElements[elem] = true
	}
	for _, elem := range arrayB {
		uniqueElements[elem] = true
	}

	// Buat slice baru untuk menyimpan elemen yang unik
	mergedArray := make([]string, 0, len(uniqueElements))
	for elem := range uniqueElements {
		mergedArray = append(mergedArray, elem)
	}

	return mergedArray
}

func main() {
	//Test Cases
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	//["king","devil jin","akuma","eddie","steve","geese"]

	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	//["sergei","jin","steve","bryan"]

	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	//["alisa","yoshimitsu","devil jin","law"]

	fmt.Println(ArrayMerge([]string{"devil jin"}, []string{"sergei"}))
	//["devil jin","sergei"]

	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	//["hwoarang"]

	fmt.Println(ArrayMerge([]string{}, []string{}))
	// []
}
