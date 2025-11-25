/*
Scrivere un programma che legga da standard input una distanza in
Km ed effettui la conversione di tale distanza in miglia (1 Km = 0.62137 mi).
*/

package main;
import "fmt";
 func main(){
	var d float64;
	fmt.Print("Inserire una distanza in km: ");
	fmt.Scan(&d);
	miglia := d * 0.62137;
	fmt.Printf("La distanza in miglia è: %f\n" , miglia)
 }