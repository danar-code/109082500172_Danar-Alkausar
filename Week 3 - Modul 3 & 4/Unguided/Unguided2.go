package main

import "fmt"

func f(x int) int { return x * x }
func g(x int) int { return x - 2 }
func h(x int) int { return x + 1 }

func main() {
	var x, y, z int

	fmt.Print("Masukkan nilai x : ")
	fmt.Scan(&x)
	fmt.Print("Masukkan nilai y : ")
	fmt.Scan(&y)
	fmt.Print("Masukkan nilai z : ")
	fmt.Scan(&z)

	fmt.Println("F(G(H(", x, "))) : ", f(g(h(x))))
	fmt.Println("G(H(F(", y, "))) : ", g(h(f(y))))
	fmt.Println("H(F(G(", z, "))) : ", h(f(g(z))))
}
