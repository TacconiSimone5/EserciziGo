package main;
import "fmt";

func main(){
	var scelta int;
	var v float64;

	fmt.Print("Scegli la conversione: \n",
	"1: secondi -> in ore\n",
	"2: secondi -> in minuti\n",
	"3: minuti -> in ore\n",
	"4: minuti -> in secondi\n",
	"5: ore -> in secondi\n",
	"6: ore -> in minuti\n",
	"7: minuti -> in giorni e ore\n",
	"8: minuti -> in anni e giorni\n",
	);
	fmt.Scan(&scelta);


	switch scelta {
	case 1:
		fmt.Printf("Inserisci il valore di secondi da convertire:");
		fmt.Scan(&v);
		res := v / 3600.0;
		fmt.Printf("%.0f secondi corrispondono a %.2f ore\n", v , res);
	case 2:
		fmt.Printf("Inserisci il valore di secondi da convertire:");
		fmt.Scan(&v);
		res := v / 60;
		fmt.Printf("%.0f secondi corrispondono a %.2f minuti\n", v , res);
	case 3:
		fmt.Printf("Inserisci il valore in minuti da convertire:");
		fmt.Scan(&v);
		res := v/60;
		fmt.Printf("%.0f minuti corrispondono a %.0f ore\n", v , res);
	case 4:
		fmt.Printf("Inserisci il valore in minuti da convertire:");
		fmt.Scan(&v);
		res := v * 60;
		fmt.Printf("%.0f minuti corrispondono a %.0f secondi\n", v , res);
	case 5:
		fmt.Printf("Inserisci il valore in ore da convertire:");
		fmt.Scan(&v);
		res := v * 3600;
		fmt.Printf("%.0f ore corrispondono a %.0f secondi\n", v , res);
	case 6:
		fmt.Printf("Inserisci il valore in ore da convertire:");
		fmt.Scan(&v);
		res := v * 60;
		fmt.Printf("%.0f ore corrispondono a %.2f minuti\n", v , res);
	case 7:
		fmt.Printf("Inserisci il valore in minuti da convertire:");
		fmt.Scan(&v);
		res := v / 1440;
		fmt.Printf("%.0f minuti corrispondono a %.2f giorni\n", v , res);
	case 8:
		fmt.Printf("Inserisci il valore in minuti da convertire:");
		fmt.Scan(&v);
		res := v / 525600;
		fmt.Printf("%.0f minuti corrispondono a %.2f anni\n", v , res);
	default:
		fmt.Println("Scelta non riconosciuta. Riprova.");
	
	}
}
