package main

import (
	"fmt"
)

func main() {
	var x, y int
	var beratIkan [1000]float64

	fmt.Scan(&x, &y)

	for i := 0; i < x; i++ {
		fmt.Scan(&beratIkan[i])
	}

	var totalWadah []float64
	var beratSementara float64
	var hitungIkan int

	for i := 0; i < x; i++ {
		beratSementara += beratIkan[i]
		hitungIkan++

		if hitungIkan == y || i == x-1 {
			totalWadah = append(totalWadah, beratSementara)
			beratSementara = 0 
			hitungIkan = 0
		}
	}

	var totalSemuaWadah float64

	for i := 0; i < len(totalWadah); i++ {
		fmt.Printf("%.2f", totalWadah[i])
		if i < len(totalWadah)-1 {
			fmt.Print(" ")
		}
		totalSemuaWadah += totalWadah[i]
	}
	fmt.Println()

	rataRata := totalSemuaWadah / float64(len(totalWadah))
	fmt.Printf("%.2f\n", rataRata)
}