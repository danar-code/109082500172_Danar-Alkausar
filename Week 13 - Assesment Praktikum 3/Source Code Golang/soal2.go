package main

import "fmt"

const NMAX = 1001

type pemain struct {
	namaDepan    string
	namaBelakang string
	gol          int
	assist       int
}

type arrPemain [NMAX]pemain

func selectionSort(A *arrPemain, n int) {
	var i, j, idxMax int
	var temp pemain

	for i = 0; i < n-1; i++ {
		idxMax = i

		for j = i + 1; j < n; j++ {
			if A[j].gol > A[idxMax].gol ||
				(A[j].gol == A[idxMax].gol && A[j].assist > A[idxMax].assist) {
				idxMax = j
			}
		}

		temp = A[i]
		A[i] = A[idxMax]
		A[idxMax] = temp
	}
}

func main() {
	var data arrPemain
	var n int

	fmt.Print("Masukkan Data Input : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&data[i].namaDepan, &data[i].namaBelakang, &data[i].gol, &data[i].assist)
	}

	selectionSort(&data, n)

	fmt.Println()
	fmt.Println("Hasil Sorting :")
	for i := 0; i < n; i++ {
		fmt.Println(data[i].namaDepan, data[i].namaBelakang, data[i].gol, data[i].assist)
	}
}
