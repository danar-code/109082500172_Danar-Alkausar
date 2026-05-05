package main

import "fmt"

const nMax = 51

type mahasiswa struct {
	NIM   string
	nama  string
	nilai int
}

type arrayMahasiswa [nMax]mahasiswa

func inputData(T *arrayMahasiswa, n *int) {
	fmt.Print("Masukkan jumlah data : ")
	fmt.Scan(n)

	for i := 0; i < *n; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&T[i].NIM, &T[i].nama, &T[i].nilai)
	}
}

func cariNilaiPertama(T arrayMahasiswa, n int, nim string) int {
	for i := 0; i < n; i++ {
		if T[i].NIM == nim {
			return T[i].nilai
		}
	}
	return -1
}

func cariNilaiTerbesar(T arrayMahasiswa, n int, nim string) int {
	max := -1
	found := false

	for i := 0; i < n; i++ {
		if T[i].NIM == nim {
			if !found {
				max = T[i].nilai
				found = true
			} else if T[i].nilai > max {
				max = T[i].nilai
			}
		}
	}
	return max
}

func main() {
	var T arrayMahasiswa
	var n int
	var nim string

	inputData(&T, &n)

	fmt.Print("Masukkan NIM mahasiswa yang ingin dicari nilai pertama dan nilai terbesarnya : ")
	fmt.Scan(&nim)

	pertama := cariNilaiPertama(T, n, nim)
	terbesar := cariNilaiTerbesar(T, n, nim)

	fmt.Printf("Nilai pertama dari NIM %s adalah %d\n", nim, pertama)
	fmt.Printf("Nilai terbesar dari NIM %s adalah %d\n", nim, terbesar)
}
