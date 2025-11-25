package main
import "fmt"

func StringheAlternate(s1, s2 string) (risultato string) {
	len1 := len(s1)
	len2 := len(s2)

	// Trovo la lunghezza massima per il ciclo.
	maxLen := len1
	if len2 > maxLen {
		maxLen = len2
	}

	risultato = ""
	const fillChar = '-'

	
	for i := 0; i < maxLen; i++ {
		
		// 1. Carattere da s1
		if i < len1 {
			// Accede al byte s1[i] e lo converte in stringa per la concatenazione.
			risultato += string(s1[i])
		} else {
			// Aggiunge il riempimento.
			risultato += string(fillChar)
		}

		// 2. Carattere da s2
		if i < len2 {
			// Accede al byte s2[i] e lo converte in stringa per la concatenazione.
			risultato += string(s2[i])
		} else {
			// Aggiunge il riempimento.
			risultato += string(fillChar)
		}
	}

	return risultato
}

func main() {
    var s1, s2 string
        
    fmt.Print("Inserisci la prima stringa di testo: ")
    fmt.Scan(&s1)
    
    fmt.Print("Inserisci la seconda stringa di testo: ")
    fmt.Scan(&s2)
    
    // Chiama la funzione e stampa il risultato
    risultato := StringheAlternate(s1, s2)
    fmt.Println(risultato)
}