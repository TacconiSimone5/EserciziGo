//Programma che inverte le cifre di un numero int a tre cifre. Esempio n = 123 --> 321

package main;
import "fmt";

func main(){
	var n int;
	fmt.Print("Inserisci un numero intero a tre cifre: ")
	fmt.Scan(&n);
	//Estrai cifre delle unità. il modulo restituisce sempre l'ultima cifra. unita 123 % 10 --> 3
	unita := n % 10;
	fmt.Printf("unità: %d\n" , unita);
	//La divisione intera per 10 rimuove l'ultima cifra. n = 123 --> 123 / 10 --> 12
	n = n / 10;
	decine := n % 10; //Applico modulo 10 al nuovo n (che nell'esempio di prima è 12). decine = 12 % 10 --> 2
	fmt.Printf("decine: %d\n" , decine);
	//La divisione intera per 10 rimuove l'ultima cifra. n = 123 --> 12 / 10 --> 1
	n = n / 10;
	centinaia := n;
	fmt.Printf("centinaia: %d\n" , centinaia);
	//Per invertitlo ovviamente andrò a moltiplicare per 100 l'unità che diventerà così centinaio, per 10 le decine e infine per 1 le centinaia.
	numeroInvertito := (unita * 100) + (decine *10) + centinaia;

	fmt.Printf("Il numero invertito è: %d\n", numeroInvertito);

}