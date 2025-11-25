package main

import "fmt"

func main() {
	var n, n1, sum, max, min, mag, inf, nulli int
	fmt.Print("")
	fmt.Scan(&n)
	fmt.Printf("Inserirci %d numeri: ", n)
	for i := 0; i < n; i++ {
		fmt.Scan(&n1)
		//fmt.Printf("i numeri: %d\n", n1)
		sum += n1

		if n1 < min {
			min = n1
		}

		if n1 > max {
			max = n1
		}

		if n1 > 0 {
			mag++
		}

		if n1 < 0 {
			inf++
		}

		if n1 == 0 {
			nulli++
		}
	}
	fmt.Printf("Somma = %d\nValore minimo = %d\nValore massimo = %d\nInteri > 0 = %d\nInteri < 0 = %d\nInteri = 0 = %d\n", sum, min, max, mag, inf, nulli)

}
