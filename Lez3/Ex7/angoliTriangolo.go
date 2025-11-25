package main;
import "fmt";

func main(){
	var a1 , a2 float64;
	
	fmt.Print("Inserire le ampiezze dei due angoli: ");
	fmt.Scan(&a1,&a2);
	
	angoliInterni := a1 + a2;
	a3 := 180 - angoliInterni;
	
	if angoliInterni > 180{
		fmt.Printf("I due angoli non appartengono ad un triangolo.");
	}else{
		fmt.Printf("Ampiezza terzo angolo: %f\n" , a3);
	}
}
