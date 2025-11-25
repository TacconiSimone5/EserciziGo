package main

import (
	"fmt"
)

func main() {
	var scelta, n int
	var totale int
	const consegna = 2
	var check = true 

	for check {
		fmt.Println("\nCosa vuoi ordinare:")
		fmt.Println("1. Patatine - 2€")
		fmt.Println("2. Hamburger - 5€")
		fmt.Println("3. Coca-Cola - 2€")
		fmt.Println("0. Termina")
		fmt.Print("Inserisci il numero della tua scelta: ")
		fmt.Scan(&scelta) 

		switch scelta {
		case 1:
			fmt.Print("Patatine? Ottimo, quante: ")
			fmt.Scan(&n)
			totale += 2 * n
			fmt.Printf("Aggiunte %d porzioni di patatine.\n", n)
		case 2:
			fmt.Print("Hamburger? Ottimo, quanti: ")
			fmt.Scan(&n)
			totale += 5 * n
			fmt.Printf("Aggiunti %d hamburger.\n", n)
		case 3:
			fmt.Print("Coca-Cola? Ottimo, quante: ")
			fmt.Scan(&n)
			totale += 2 * n
			fmt.Printf("Aggiunte %d Coca-Cola.\n", n)
		case 0:
			check = false 
		default:
			fmt.Println("Scelta non valida, riprova.")
		}
	}

	if totale == 0 {
		fmt.Println("Nessun ordine effettuato. Arrivederci!")
		return
	}

	totaleConConsegna := totale + consegna
	fmt.Printf("Sono %d€ + 2€ di consegna. Totale: %d€\n", totale, totaleConConsegna)
}
