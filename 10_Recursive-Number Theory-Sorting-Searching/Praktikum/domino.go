package main

import "fmt"

func playingDomino(cards [][]int, deck []int) interface{} {
	for _, card := range cards {
		for _, val := range card {
			for _, val2 := range deck {
				if val == val2 {
					return card
				}
			}
		}
	}

	return "tutup kartu"
}

func main() {
	fmt.Println(playingDomino([][]int{{6, 5}, {3, 2}, {2, 1}}, []int{3, 4}))
}
