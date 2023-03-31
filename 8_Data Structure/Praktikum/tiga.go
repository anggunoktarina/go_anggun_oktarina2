package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {
	counts := make(map[int]int)
	for _, digitStr := range angka {
		digit, err := strconv.Atoi(string(digitStr))
		if err != nil {
			panic("invalid angka string")
		}
		counts[digit]++
	}
	var munculSekali []int
	for digit, count := range counts {
		if count == 1 {
			munculSekali = append(munculSekali, digit)
		}
	}
	return munculSekali
}

func main() {
	fmt.Println(munculSekali("1234123"))    // [4]
	fmt.Println(munculSekali("76523752"))   // [6, 3]
	fmt.Println(munculSekali("12345"))      // [1,2,3,4,5]
	fmt.Println(munculSekali("1122334455")) // []
	fmt.Println(munculSekali("0872504"))    // [8,7,2,5,4]
}
