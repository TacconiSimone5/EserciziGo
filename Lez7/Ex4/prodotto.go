package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	// os.Args[0] è il nome del programma, quindi partiamo da [1:]
	argomenti := os.Args[1:]
	prodotto := 1
	numeriTrovati := false

	// 3. Iteriamo su ogni argomento (stringa) passato
	for _, argomento := range argomenti {

		// Proviamo a convertire la stringa in un intero
		numero, err := strconv.Atoi(argomento)

		// Se err è nil, significa che la conversione ha avuto successo (è un numero)
		if err == nil {
			prodotto = prodotto * numero
			numeriTrovati = true
		}
		// Se err != nil (es. "ciao" o "a"), ignoriamo semplicemente l'argomento
	}

	// 4. Stampiamo il risultato solo se abbiamo trovato dei numeri validi
	if numeriTrovati {
		fmt.Printf("Il risultato della moltiplicazione tra i numeri interi è %d\n", prodotto)
	}
}
