package main

import "fmt"

func main() {
	var berat, kg, gram int
	var biayaKg, biayaGram, total int

	fmt.Print("Masukkan total berat (gram): ")
	fmt.Scan(&berat)

	kg = berat / 1000
	gram = berat % 1000

	biayaKg = kg * 10000

	if kg > 10 {
		biayaGram = 0
	} else if gram >= 500 {
		biayaGram = gram * 5
	} else {
		biayaGram = gram * 15
	}

	total = biayaKg + biayaGram

	fmt.Println()
	fmt.Println("===== Detail Perhitungan =====")
	fmt.Println("Detail berat :", kg, "kg +", gram, "gram")
	fmt.Println("Detail biaya : Rp.", biayaKg, "+ Rp.", biayaGram)
	fmt.Println("Total biaya: Rp", total)
}
