package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {

	if len(a) <= len(b) {
		if strings.Contains(b, a) {
			return a
		}
	} else {
		if strings.Contains(a, b) {
			return b
		}
	}

	return "invalid"
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))
	fmt.Println(Compare("KANGOORO", "KANG"))
	fmt.Println(Compare("KI", "KIJANG"))
	fmt.Println(Compare("KUPU-KUPU", "KUPU"))
	fmt.Println(Compare("ILALANG", "ILA"))
}
