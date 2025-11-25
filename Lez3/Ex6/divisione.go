package main;
import "fmt";

func main(){
	var n1 , n2 float64;
	var divisione float64;
	fmt.Print("Inserire primo numero: ");
	fmt.Scan(&n1);
	fmt.Print("Inserire secondo numero: ");
	fmt.Scan(&n2);

	divisione = n1/n2;

	if n2 == 0{
		fmt.Printf("Impossibile\n");
	}else {
		fmt.Printf("Quoziente: %f\n", divisione);
	}
}
