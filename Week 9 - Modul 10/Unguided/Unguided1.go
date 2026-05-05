package main

import (
	"fmt"
)

func main() {
	var n int
	var beratKelinci [1000]float64 

	fmt.Scan(&n)

	if n <= 0 {
		return
	}

	for i := 0; i < n; i++ {
		fmt.Scan(&beratKelinci[i])
	}

	min := beratKelinci[0]
	max := beratKelinci[0]

	for i := 1; i < n; i++ {
		if beratKelinci[i] < min {
			min = beratKelinci[i] 
		}
		if beratKelinci[i] > max {
			max = beratKelinci[i] 
		}
	}

	fmt.Printf("%.2f %.2f\n", min, max)
}