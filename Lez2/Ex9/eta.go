/*
Scrivere un programma che legga da standard input le età di due persone
 (espresse in anni) e calcoli:

la somma delle età inserite;
la media delle età inserite;
la media delle età inserite arrotondata per difetto all'intero inferiore;
la media delle età inserite arrotondata per eccesso all'intero superiore;
la somma e la media delle età che le due persone avranno fra 10 anni.
Suggerimento: la media arrotondata per difetto può essere calcolata usando 
la funzione math.Floor del package math nel seguente modo:
var mediaArrotondataDifetto float64 = math.Floor(media)
*/

package main;
import "fmt";
import "math";
 func main(){
	var eta1 , eta2 int;

	fmt.Print("Inserire la prima età: ");
	fmt.Scan(&eta1);
	fmt.Print("Inserire la seconda età: ");
	fmt.Scan(&eta2);

	//Somma
	sum := eta1 + eta2;
	//Media
	avg := float64(eta1 + eta2) / 2;

	fmt.Printf("Somma età = %d\n La media è = %f\n La media arrotondata per difetto è = %f\n La media arrotondata per eccesso è = %f\n La somma delle età che le due persone avranno fra 10 anni = %d\n La media delle età che le due persone avranno fra 10 anni = %f\n" , sum, avg, math.Floor(avg), math.Ceil(avg), sum + 20, float64(sum+20)/2);


 }