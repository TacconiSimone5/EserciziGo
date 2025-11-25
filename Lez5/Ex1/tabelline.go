package main
import "fmt"

func tabellina (numero int) bool{
if numero >= 1 && numero <= 9  {
	//fmt.Printf("Tabellina del: &d " , numero)
	for i := 0; i <= 10; i++{
		ris := numero * i
		fmt.Print(ris, " ");
	}
	return true
}
fmt.Print("Porgramma temrinato.\n")
return false 
}

func main(){
	var n int
	fmt.Print("Inserire un numero: ")
	fmt.Scan(&n)
	fmt.Printf("Tabellina del: %d = " , n)
	tabellina(n)
}
