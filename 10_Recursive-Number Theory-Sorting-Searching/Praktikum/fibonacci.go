package main

import "fmt"

func fibonacci(angka int) (int, error) {
	if angka < 0 {
		return 0, fmt.Errorf("invalid input value %d: angka must be non-negative", angka)
	}
	if angka <= 1 {
		return angka, nil
	}
	fib1, err1 := fibonacci(angka - 1)
	if err1 != nil {
		return 0, err1
	}
	fib2, err2 := fibonacci(angka - 2)
	if err2 != nil {
		return 0, err2
	}
	return fib1 + fib2, nil
}

func main() {
	for i := -1; i < 12; i++ {
		result, err := fibonacci(i)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
	}

	fmt.Println(fibonacci(0))  // 0
	fmt.Println(fibonacci(2))  // 1
	fmt.Println(fibonacci(9))  // 34
	fmt.Println(fibonacci(10)) // 55
	fmt.Println(fibonacci(12)) // 144
}
