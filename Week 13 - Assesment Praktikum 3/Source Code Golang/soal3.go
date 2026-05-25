package main

import "fmt"

const NMAX = 1000000

type partai struct {
	nama  int
	suara int
}

type tabPartai [NMAX]partai

func posisi(t tabPartai, n int, nama int) int {
	for i := 0; i < n; i++ {
		if t[i].nama == nama {
			return i
		}
	}
	return -1
}

func insertionSort(p *tabPartai, n int) {
	var i, j int
	var temp partai

	for i = 1; i < n; i++ {
		temp = p[i]
		j = i - 1

		for j >= 0 && p[j].suara < temp.suara {
			p[j+1] = p[j]
			j--
		}

		p[j+1] = temp
	}
}

func main() {
	var p tabPartai
	var n, suaraMasuk, idx int

	fmt.Println("Masukkan proses input suara :")
	fmt.Scan(&suaraMasuk)

	for suaraMasuk != -1 {
		idx = posisi(p, n, suaraMasuk)

		if idx == -1 {
			p[n].nama = suaraMasuk
			p[n].suara = 1
			n++
		} else {
			p[idx].suara++
		}

		fmt.Scan(&suaraMasuk)
	}

	insertionSort(&p, n)

	fmt.Println()
	fmt.Println("Hasil Perhitungan suara :")
	for i := 0; i < n; i++ {
		fmt.Printf("%d(%d) ", p[i].nama, p[i].suara)
	}
}
