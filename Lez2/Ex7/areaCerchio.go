/*
Da input si ha il raggio di un cerchio e ne calcoli la circonferenza e area

A = r * r * pi_greco
Circ = 2 * raggio * pi_greco
pi_greco è dentro la costante Pi bel package math
*/

package main;
import "fmt";
import "math";

func main(){
	var r float64;
	fmt.Print("Inserire il raggio: ");
	fmt.Scan(&r);

	circonferenza := 2 * r * math.Pi;
	area := r * r * math.Pi;

	fmt.Printf("La circonferenza del cerchio è =  %.3f\n L'area del cerchio = %.3f\n" , circonferenza, area);
}