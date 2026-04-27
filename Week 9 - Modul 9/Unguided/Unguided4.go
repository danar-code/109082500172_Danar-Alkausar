package main

import "fmt"

func main() {
	var input string
	fmt.Scan(&input)

	karakter := []rune(input)
	n := len(karakter)

	isPalindrom := true
	for i := 0; i < n/2; i++ {
		if karakter[i] != karakter[n-1-i] {
			isPalindrom = false
			break
		}
	}

	fmt.Print("Reverse: ")
	for i := n - 1; i >= 0; i-- {
		fmt.Print(string(karakter[i]))
	}
	fmt.Printf("\nPalindrom: %t\n", isPalindrom)
}
