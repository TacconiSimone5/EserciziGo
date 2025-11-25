
package main;
import "fmt";

func main(){
	var n int;
	fmt.Print("Inserire numero: ");
	fmt.Scan(&n);

	if n > 0 {
		fmt.Printf("+%d\n" , n);
	}else {
		fmt.Printf("%d\n",n );
	}
}