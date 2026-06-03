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

func selectionSortDesc(T *arrInt, n int) {
	var t, i, j, idx_max int
	i = 1
	for i <= n-1 {
		idx_max = i - 1
		j = i
		for j < n {
			if T[idx_max] < T[j] {
				idx_max = j
			}
			j = j + 1
		}
		t = T[idx_max]
		T[idx_max] = T[i-1]
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
		var arrGanjil, arrGenap arrInt
		var nGanjil, nGenap int

		for j := 0; j < m; j++ {
			var val int
			fmt.Scan(&val)
			if val%2 != 0 {
				arrGanjil[nGanjil] = val
				nGanjil++
			} else {
				arrGenap[nGenap] = val
				nGenap++
			}
		}

		selectionSortAsc(&arrGanjil, nGanjil)
		selectionSortDesc(&arrGenap, nGenap)

		for j := 0; j < nGanjil; j++ {
			fmt.Print(arrGanjil[j], " ")
		}
		for j := 0; j < nGenap; j++ {
			fmt.Print(arrGenap[j], " ")
		}
		fmt.Println()
	}
}
