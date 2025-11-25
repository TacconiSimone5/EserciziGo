package main
import "fmt"
func StampaRigaInizioAsterisco (lunghezza int) {
	var i int
	for i = 0; i < lunghezza; i++{ //Righe
			if i%2 == 0{
				fmt.Print("* ");
			}else{
				fmt.Print("+ ");
			}
	}
}
func StampaRigaInizioPiù (lunghezza int) {
	var i int
	for i = 0; i < lunghezza; i++{ //Righe
			if i%2 == 0{
				fmt.Print("+ ");
			}else{
				fmt.Print("* ");
			}
	}
}
func StampaScacchiera(dimensione int) {
	if dimensione <= 0{
		return
	}else{
		for i := 0; i < dimensione; i++{
			if i%2 == 0{
			StampaRigaInizioAsterisco(dimensione)
			}else{
			StampaRigaInizioPiù(dimensione)
			}
			fmt.Println()
		}
	return
	}
}
func main(){
	var d int
	fmt.Print("Inserire la dimensione: ")
	fmt.Scan(&d)
	StampaScacchiera(d)
}


