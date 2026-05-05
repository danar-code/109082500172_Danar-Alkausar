package main

import "fmt"

func main() {
	var n, x, idxHapus int
	fmt.Print("Masukkan jumlah elemen: ")
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Print("Indeks Ganjil: ")
	for i := 1; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}

	fmt.Print("\nMasukkan x: ")
	fmt.Scan(&x)
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}

	fmt.Print("\nHapus indeks: ")
	fmt.Scan(&idxHapus)
	arr = append(arr[:idxHapus], arr[idxHapus+1:]...)
	fmt.Println("Isi array setelah hapus:", arr)
}
