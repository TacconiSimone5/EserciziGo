package main
import "fmt"

func main(){
	var n1, n2 int
	var res int
	fmt.Print("Inserisci due numeri: ");
	fmt.Scan(&n1,&n2);
	for i := n1+1; i < n2; i++{
		if i%2 == 0{
			res  +=  i
		}
	}
	fmt.Printf("Somma = %d\n", res);

}
