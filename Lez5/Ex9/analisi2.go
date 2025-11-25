package main

import (
	"fmt"
	"unicode"
)
func vocaleMin (l rune) bool{
	l_lower := unicode.ToLower(l)
	switch l_lower {
	case 'a','e','i','o','u':
		return true
	default:
		return false
	}
}
func main(){
	var s string
	var vocaliMaiuscola, VocaliMinuscola, consonantiMaiuscola, consonantiMinuscola int
	fmt.Print("Inserire parola: ")
	fmt.Scan(&s)
	
	for _, lettera := range s{
		if unicode.IsLetter(lettera){
			if vocaleMin(lettera){
				if unicode.IsUpper(lettera){
					vocaliMaiuscola++
				}else{
					VocaliMinuscola++
				}
			}else{
				if unicode.IsUpper(lettera){
					consonantiMaiuscola++
				}else{
					if unicode.IsLower(lettera){
						consonantiMinuscola++
					}
				}
			}
		}
	}
	fmt.Println("\nRisultati dell'analisi:")
    fmt.Printf("Numero di vocali maiuscole: %d\n", vocaliMaiuscola)
    fmt.Printf("Numero di vocali minuscole: %d\n", VocaliMinuscola)
    fmt.Printf("Numero di consonanti maiuscole: %d\n", consonantiMaiuscola)
    fmt.Printf("Numero di consonanti minuscole: %d\n", consonantiMinuscola)
}