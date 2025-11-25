package main

import "fmt"

func primo(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func numeriPrimi(limite int) {
	for i := 2; i < limite; i++ {
		if primo(i) {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}
func main() {
	var n int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)
	if n <= 0 {
		fmt.Println("La soglia inserita non è positiva.")
	} else {
		fmt.Printf("I numeri primi inferiori a %d\n", n)
		numeriPrimi(n)
	}

}
