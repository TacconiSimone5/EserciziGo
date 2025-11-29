package main

import "fmt"

func main() {
	var dimensione int

	fmt.Print("Inserire la quantità di numeri: ")
	fmt.Scan(&dimensione)

	//Creo lo Slice
	valori := make([]string, dimensione)
	var input string
	sliceFinale := make([]string, dimensione)

	fmt.Printf("Inserisci i %d valori:\n", dimensione)

	//Popolo lo slice
	for i := 0; i < dimensione; i++ {
		fmt.Printf("Valore %d: ", i+1)
		fmt.Scanln(&input) //Blocca l'iterazione e attende che l'utente digiti un valore e prema invio. Il valore digitato verrà memorizzato nella variabile input
		valori[i] = input
	}

	fmt.Println("\n--- Contrario ---")
	i := 0
	for j := dimensione - 1; j >= 0; j-- {
		//fmt.Printf("Elemento %d: %s\n", j, valori[j])

		inverto := valori[j]
		sliceFinale[i] = inverto
		i++

	}
	fmt.Println("Slice:", sliceFinale)
}
