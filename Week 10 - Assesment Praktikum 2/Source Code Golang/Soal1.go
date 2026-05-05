package main

import "fmt"

type set [2022]int

func exist(T set, n int, val int) bool {
	for i := 0; i < n; i++ {
		if T[i] == val {
			return true
		}
	}
	return false
}

func inputSet(T *set, n *int) {
	*n = 0
	var val int

	for *n < 2022 {
		fmt.Scan(&val)

		if exist(*T, *n, val) {
			return
		}

		T[*n] = val
		*n++
	}
}

func findIntersection(s1, s2 set, n1, n2 int, s3 *set, n3 *int) {
	*n3 = 0

	for i := 0; i < n1; i++ {
		if exist(s2, n2, s1[i]) && !exist(*s3, *n3, s1[i]) {
			s3[*n3] = s1[i]
			*n3++
		}
	}
}

func printSet(T set, n int) {
	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(T[i])
	}
	fmt.Println()
}

func main() {
	var s1, s2, s3 set
	var n1, n2, n3 int

	inputSet(&s1, &n1)
	inputSet(&s2, &n2)
	findIntersection(s1, s2, n1, n2, &s3, &n3)
	printSet(s3, n3)
}
