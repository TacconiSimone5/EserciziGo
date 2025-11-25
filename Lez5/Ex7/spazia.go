package main
import "fmt"

func main(){
	var s string
	fmt.Print("Inserisci una stringa di testo: ")
	fmt.Scan(&s)

	for _, aRune := range s{
		fmt.Printf("%c ",aRune )
	}
	fmt.Println()
}
