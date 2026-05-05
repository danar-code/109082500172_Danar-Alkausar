package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const nProv = 10

type NamaProv [nProv]string
type PopProv [nProv]int
type TumbuhProv [nProv]float64

func InputData(prov *NamaProv, pop *PopProv, tumbuh *TumbuhProv) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("--- Masukkan Nama Provinsi, Populasi Provinsi, Angka Pertumbuhan Provinsi ---")
	for i := 0; i < nProv; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		scanner.Scan()
		line := scanner.Text()
		parts := strings.Split(line, " ")

		lastIdx := len(parts) - 1
		prov[i] = strings.Join(parts[:lastIdx-1], " ")
		fmt.Sscanf(parts[lastIdx-1], "%d", &pop[i])
		fmt.Sscanf(parts[lastIdx], "%f", &tumbuh[i])
	}
}

func ProvinsiTercepat(tumbuh TumbuhProv) int {
	idx := 0
	for i := 1; i < nProv; i++ {
		if tumbuh[i] > tumbuh[idx] {
			idx = i
		}
	}
	return idx
}

func IndeksProvinsi(prov NamaProv, nama string) int {
	for i := 0; i < nProv; i++ {
		if prov[i] == nama {
			return i
		}
	}
	return -1
}

func Prediksi(prov NamaProv, pop PopProv, tumbuh TumbuhProv) {
	fmt.Println("\n--- Prediksi Jumlah Penduduk Tahun Depan Pada Provinsi Dengan Pertumbuhan Diatas 2% ---")
	for i := 0; i < nProv; i++ {
		if tumbuh[i] > 0.02 {
			prediksi := float64(pop[i]) * (tumbuh[i] + 1)
			fmt.Printf("%s %.0f\n", prov[i], prediksi)
		}
	}
}

func main() {
	var prov NamaProv
	var pop PopProv
	var tumbuh TumbuhProv
	var namaCari string

	InputData(&prov, &pop, &tumbuh)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	namaCari = scanner.Text()

	tercepat := ProvinsiTercepat(tumbuh)
	fmt.Printf("\nProvinsi dengan angka pertumbuhan tercepat : %s\n", prov[tercepat])

	idxCari := IndeksProvinsi(prov, namaCari)
	fmt.Printf("\nData provinsi yang dicari : %s\n", prov[idxCari])

	Prediksi(prov, pop, tumbuh)
}
