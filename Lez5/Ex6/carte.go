package main

import "fmt"

func main(){
	// Asso di Cuori (Unicode 127153)
	var assoCuori rune = '\U0001F0B1' 

	fmt.Println("Sequenza Carte di Cuori (Asso al 10):")

	// Stampiamo l'Asso una volta (l'indice i=0 lo stamperà per primo)
	// Il ciclo for i < 10 stamperà 10 carte: i=0 (Asso) fino a i=9 (Dieci)
	for i := 0 ;i <= 9; i++{ 
		// Calcola il codice numerico della carta successiva.
		cardCode := int(assoCuori) + i
		
		// Converte il codice numerico (int) in un carattere Unicode (rune).
		currentCard := rune(cardCode) 
		fmt.Printf("Simbolo: %c - Codice numerico in base 10: %d\n", currentCard, cardCode)
	}
}
