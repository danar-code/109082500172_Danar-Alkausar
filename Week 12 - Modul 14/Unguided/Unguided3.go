package main

import "fmt"

const NMax = 1000000

type arrInt [NMax]int

func insertionSortAsc(T *arrInt, n int) {
	var temp, i, j int
	i = 1
	for i <= n-1 {
		j = i
		temp = T[j]
		for j > 0 && temp < T[j-1] {
			T[j] = T[j-1]
			j = j - 1
		}
		T[j] = temp
		i = i + 1
	}
}

func main() {
	var arr arrInt
	var n, val int

	for {
		fmt.Scan(&val)
		if val < 0 {
			break
		}
		arr[n] = val
		n++
	}

	insertionSortAsc(&arr, n)

	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	if n > 1 {
		jarak := arr[1] - arr[0]
		tetap := true
		for i := 1; i < n-1; i++ {
			if arr[i+1]-arr[i] != jarak {
				tetap = false
				break
			}
		}
		if tetap {
			fmt.Printf("Data berjarak %d\n", jarak)
		} else {
			fmt.Println("Data berjarak tidak tetap")
		}
	}
}
