package main

import "fmt"

func caesar(offset int, input string) string {
	var output string
	for _, r := range input {
		if r >= 'a' && r <= 'z' {
			// shift the character by offset positions
			r = 'a' + (r-'a'+rune(offset))%26
		} else if r >= 'A' && r <= 'Z' {
			// shift the character by offset positions
			r = 'A' + (r-'A'+rune(offset))%26
		}
		output += string(r)
	}
	return output
}

func main() {
	fmt.Println(caesar(3, "abc"))                           // def
	fmt.Println(caesar(2, "alta"))                          // def
	fmt.Println(caesar(10, "alterraacademy"))               // kvdobbkkmknowi
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))    // bcdefghijklmnopqrstuvwxyza
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl
}
