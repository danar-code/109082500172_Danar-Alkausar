package main

import "fmt"

func ganjil(n int) {
	if n < 1 {
		return
	}

	ganjil(n - 1)

	if n%2 != 0 {
		fmt.Print(n, " ")
	}
}

func main() {
	var n int

	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)

	fmt.Println("Bilangan ganjil:")
	ganjil(n)
}
