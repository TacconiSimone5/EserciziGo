package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Println("Inserire una stringa: ")
	fmt.Scan(&s)
	occorrenze(s)
	fmt.Println("Occorrenze:", occorrenze)

}

func occorrenze(s string) [26]int {
	//Data una stringa restituisce un array con le occorrenze delle 26 lettere dell'alfabeto contenute nella stringa
	//strings.toLower(s)
	sMinuscola := strings.ToLower(s)

	var occorrenze [26]int

	for i := 0; i < len(s); i++{
		occorrenze[i] = s[i]
	}
	return occorrenze
}
