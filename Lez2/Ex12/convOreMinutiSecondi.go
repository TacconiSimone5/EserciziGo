/*
Scrivere un programma che prenda in ingresso un valore intero rappresentante
una quantità di tempo in secondi e restituisca il numeri di ore minuti e secondi

*/

package main;
import "fmt";

func main(){
	var n int;
	fmt.Print("Inserire secondi: ");
	fmt.Scan(&n);

	ore := n / 3600;
	secondiRimanenti := n % 3600;
	minuti := secondiRimanenti / 60;
	secondi := secondiRimanenti % 60;

	fmt.Printf("Ore - Minuti - Secondi: %d:%d:%d\n" , ore , minuti,secondi);


}