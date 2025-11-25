package main;
import "fmt";

func main(){
	var n1, n2 int;
	fmt.Print("Inserire due numeri interi: ");
	fmt.Scan(&n1,&n2);
	div := n1 / n1;
    res := n1 % n2;
    fmt.Printf("Quoziente = %d\n Resto = %d\n" , div, res);
}