package main

import (
	"fmt"
)

func main() {
	var sequenza string

	fmt.Print("Inserire sequenza: ")
	fmt.Scan(&sequenza)

	var caratteriSlice []string

	for _, char := range sequenza {
		lettera := string(char) //char è un rune dobbiamo convertirlo in stringa

		caratteriSlice = append(caratteriSlice, lettera)
	}
	fmt.Println("Slice creato:", caratteriSlice)
	
	for _, elem := range caratteriSlice {
         // Fai qui i controlli su 'elem' (che ora è una stringa, es "1")
         // Nota: confrontare stringhe è diverso da confrontare rune
         if elem < "0" || elem > "9" {
             fmt.Println("Non è un numero:", elem)
         }
    }
}
