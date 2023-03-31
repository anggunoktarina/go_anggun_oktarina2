package main

import (
	"fmt"
)

type Student struct {
	Nama  []string
	Score []int
}

func (e Student) averageScore() float64 {
	sum := 0
	for _, score := range e.Score {
		sum += score
	}
	return float64(sum) / float64(len(e.Score))
}

func (e Student) findMinScore() int {
	min := e.Score[0]
	for _, score := range e.Score {
		if score < min {
			min = score
		}
	}
	return min
}

func (e Student) findMaxScore() int {
	max := e.Score[0]
	for _, score := range e.Score {
		if score > max {
			max = score
		}
	}
	return max
}

func main() {
	e := Student{
		Nama:  []string{"Rizky", "Kobar", "Ismail", "Umam", "Yopan"},
		Score: []int{80, 75, 70, 75, 60},
	}

	fmt.Println("Avarage Score:", e.averageScore())
	fmt.Println("Minimun Score", e.findMinScore())
	fmt.Println("Maximun Score", e.findMaxScore())
}
