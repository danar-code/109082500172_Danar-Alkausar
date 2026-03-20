package main

import (
	"fmt"
	"math"
)

func hitungPersegi(sisi int) {
	fmt.Println("Luas persegi : ", sisi*sisi)
	fmt.Println("Keliling persegi : ", 4*sisi)
}

func hitungPersegiPanjang(panjang, lebar int) {
	fmt.Println("Luas persegi panjang : ", panjang*lebar)
	fmt.Println("Keliling persegi panjang : ", 2*(panjang+lebar))
}

func hitungLingkaran(jarijari float64) {
	fmt.Printf("Luas lingkaran : %f\n", math.Pi*jarijari*jarijari)
	fmt.Printf("Keliling lingkaran : %.4f\n", 2*math.Pi*jarijari)
}

func main() {
	var pilihan int

	fmt.Println("--- PROGRAM BANGUN DATAR ---")
	fmt.Println("1. Hitung luas & keliling persegi")
	fmt.Println("2. Hitung luas & keliling persegi panjang")
	fmt.Println("3. Hitung luas & keliling lingkaran")
	fmt.Print("Pilihan : ")
	fmt.Scan(&pilihan)
	fmt.Println()

	switch pilihan {
	case 1:
		var sisi int
		fmt.Print("Masukkan sisi : ")
		fmt.Scan(&sisi)
		hitungPersegi(sisi)
	case 2:
		var panjang, lebar int
		fmt.Print("Masukkan panjang : ")
		fmt.Scan(&panjang)
		fmt.Print("Masukkan lebar : ")
		fmt.Scan(&lebar)
		hitungPersegiPanjang(panjang, lebar)
	case 3:
		var jarijari float64
		fmt.Print("Masukkan jari-jari : ")
		fmt.Scan(&jarijari)
		hitungLingkaran(jarijari)
	}
}
