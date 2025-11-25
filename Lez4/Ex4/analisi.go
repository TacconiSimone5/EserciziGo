package main
import "fmt"

func main(){
	var s string
	var maiuscola, minuscola, vocale int
	fmt.Print("Inserire parola: ")
	fmt.Scan(&s)

	for _, lettera := range s{
		switch {
		case lettera >= 'A' && lettera <= 'Z':
		maiuscola ++
		case lettera >= 'a' && lettera <= 'z':
		minuscola ++
		}
		switch lettera {
		case 'a' ,'A' , 'e' , 'E' , 'I' , 'i' , 'O' , 'o' , 'U' , 'u':
		vocale ++
		}
	}
}