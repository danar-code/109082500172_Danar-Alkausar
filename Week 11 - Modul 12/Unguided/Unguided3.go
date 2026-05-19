package main

import "fmt"

const NMAX = 1000000

var data [NMAX]int

func isiArray(n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
}

func posisi(n, k int) int {
	kr := 0
	kn := n - 1
	found := -1

	for kr <= kn && found == -1 {
		med := (kr + kn) / 2

		if data[med] < k {
			kr = med + 1
		} else if data[med] > k {
			kn = med - 1
		} else {
			found = med
		}
	}

	return found
}

func main() {
	var n, k int

	fmt.Scan(&n, &k)

	isiArray(n)

	idx := posisi(n, k)

	if idx == -1 {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Println(idx)
	}
}
