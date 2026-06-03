package main

import "fmt"

const NMax = 1000000

type arrInt [NMax]int

func selectionSortAsc(T *arrInt, n int) {
	var t, i, j, idx_min int
	i = 1
	for i <= n-1 {
		idx_min = i - 1
		j = i
		for j < n {
			if T[idx_min] > T[j] {
				idx_min = j
			}
			j = j + 1
		}
		t = T[idx_min]
		T[idx_min] = T[i-1]
		T[i-1] = t
		i = i + 1
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)
		var arr arrInt
		for j := 0; j < m; j++ {
			fmt.Scan(&arr[j])
		}

		selectionSortAsc(&arr, m)

		for j := 0; j < m; j++ {
			fmt.Print(arr[j], " ")
		}
		fmt.Println()
	}
}
