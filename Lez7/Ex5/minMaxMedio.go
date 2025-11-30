package main

import (
	"fmt"
	"os"
	"strconv"
)
func main() {

	numeri := leggiNumeri()

	if len(numeri) == 0{
		fmt.Println("Nessun numero valido.")
		return
	}

	min := minimo(numeri)
	max := massimo(numeri)
	med := media(numeri)

	fmt.Printf("Minimo: %d\n", min)
	fmt.Printf("Massimo: %d\n", max)
	fmt.Printf("Media: %f\n", med)
	
}

func leggiNumeri() []int {
	var risultato []int
	argomenti := os.Args[1:]

	for _, arg := range argomenti {
		valore, err := strconv.Atoi(arg)
		if err == nil {
			// Se è un numero valido, lo aggiungiamo alla slice (array dinamico)
			risultato = append(risultato, valore)
		}
	}
	return risultato
}

func minimo(sl []int) int {
	min := sl[0]
	for _, v := range sl {
		if v < min {
			min = v
		}
	}
	return min
}

func massimo(sl []int) int {
	max := sl[0]
	for _, v := range sl {
		if v > max {
			max = v
		}
	}
	return max
}

func media(sl []int) float64 {
	somma := sl[0]
	for _, v := range sl{
		somma += v
	}
	return float64(somma) / float64(len(sl))
}


