package main;
import "fmt";

func main(){
	var n int;
	fmt.Print("Inserire numero intero: ");
	fmt.Scan(&n);
	if n%3 == 0 && n%5 == 0{
		fmt.Printf("Fizz Buzz \n");
	}else{
		if n%3 == 0{
		fmt.Printf("Fizz \n");
	}else if n%5 == 0{
		fmt.Printf("Buzz \n");
	}
	}
	
}
