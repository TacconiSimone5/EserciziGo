
package main;
import "fmt";

func main(){
	var n int;
	fmt.Print("Inserire un numero: ");
	fmt.Scan(&n);

	if n%10 == 0{
		fmt.Printf("%d è multiplo di 10\n" , n);

	}else{
		fmt.Printf("%d non è multiplo di 10\n" , n);
	}
}
