package main

import "fmt"

func main() {
	var suara [21]int
	var masuk, sah, num int

	for {
		fmt.Scan(&num)

		if num == 0 {
			break
		}

		masuk++

		if num >= 1 && num <= 20 {
			sah++
			suara[num]++
		}
	}

	fmt.Printf("Suara masuk: %d\n", masuk)
	fmt.Printf("Suara sah: %d\n", sah)

	for i := 1; i <= 20; i++ {
		if suara[i] > 0 {
			fmt.Printf("%d: %d\n", i, suara[i])
		}
	}
}
