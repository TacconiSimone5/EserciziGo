/*
Scrivere un programma che legga da standard input le misure dell’altezza e della base 
di un rettangolo e ne calcoli il perimetro e l’area.

Esempio output:
$ go run rettangolo.go
Inserisci la base: 20
Inserisci l'altezza: 10
Perimetro = 60
Area = 200
*/

package main;

import "fmt";

func main(){
	var altezza, base float64;
	fmt.Print("Inserisci la base:");
	fmt.Scan(&base);
	fmt.Print("Inserire l'altezza: ");
	fmt.Scan(&altezza);
	//fmt.Printf("base = %f Altezza = %f" , base, altezza);

	perimetro := 2 * (base + altezza);
	area := base * altezza;
	
	fmt.Printf("Perimetro = %.0f\n" , perimetro);
	fmt.Printf("Area = %.0f\n" , area);
}