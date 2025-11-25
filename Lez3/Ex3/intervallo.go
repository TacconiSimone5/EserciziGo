package main;
import "fmt";

func main (){
	var voto int;
	fmt.Print("Inserire voto in centesimi: ");
	fmt.Scan(&voto);

	if voto < 60{
		fmt.Printf("Il voto %d è insufficiente. \n" , voto);
	}else if voto < 70{
		fmt.Printf("Il voto %d è sufficiente. \n" , voto);
	}else if voto < 80 {
		fmt.Printf("Il voto %d è buono. \n" ,voto);
	}else if voto < 90{
		fmt.Printf("Il voto %d è distinto. \n" , voto);
	}else if voto <= 100{
		fmt.Printf("Il voto %d è ottimo. \n" , voto);
	}else{
		fmt.Print("Errore");
	}
}
