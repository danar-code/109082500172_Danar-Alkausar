package main

import "fmt"

type suhu float64

func CelciusToReamur(c suhu) suhu {
	return 0.8 * c
}

func CelciusToFahrenheit(c suhu) suhu {
	return (9.0/5.0)*c + 32
}

func CelciusToKelvin(c suhu) suhu {
	return c + 273.15
}

func main() {
	var c suhu

	fmt.Println("=== KONVERTER CELCIUS ===")
	fmt.Print("Masukkan suhu (celcius) : ")
	fmt.Scan(&c)

	fmt.Println()
	fmt.Printf("%.0f celcius = %.1f reamur\n", c, CelciusToReamur(c))
	fmt.Printf("%.0f celcius = %.1f fahrenheit\n", c, CelciusToFahrenheit(c))
	fmt.Printf("%.0f celcius = %.2f kelvin\n", c, CelciusToKelvin(c))
}
