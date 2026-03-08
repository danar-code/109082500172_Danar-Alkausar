package main

import "fmt"

func main() {
	var year int

	fmt.Print("Masukkan tahun : ")
	fmt.Scan(&year)

	fmt.Println("Kabisat :", isLeapYear(year))
}

// isLeapYear memeriksa apakah suatu tahun adalah tahun kabisat
func isLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	}

	if year%4 == 0 && year%100 != 0 {
		return true
	}

	return false
}
