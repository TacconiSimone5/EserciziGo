package main
import "fmt"

func main (){
	var n int
	fmt.Println("Inserire un numero intero: ")
	fmt.Scan(&n);
	var i,j int
	for i = 0; i < n; i++{ //Righe
		for j = 0; j < n; j++{ //Colonne
			if i%3 == 0{
				fmt.Print("* ");
			}else if i%3 == 1{
				fmt.Print("+ ");
			}else{
				fmt.Print("° ");
			}
		}
		fmt.Println();
	}
	
}