package main;
import "fmt";
func main(){
	var m, q float64;
	var px, py float64;
	fmt.Print("Inserire m e q: ");
	fmt.Scan(&m,&q);
	fmt.Print("Inserire px e py: ");
	fmt.Scan(&px,&py);

	yRetta := m*px+q;
	if py == yRetta{
		fmt.Printf("Il punto P(%.2f, %.2f) APPARTIENE alla retta y = %.2fx + %.2f\n", px, py, m, q)
	} else if py > yRetta {
		
		fmt.Printf("Il punto P(%.2f, %.2f) STA SOPRA la retta y = %.2fx + %.2f\n", px, py, m, q)
	} else {
		fmt.Printf("Il punto P(%.2f, %.2f) STA SOTTO la retta y = %.2fx + %.2f\n", px, py, m, q)
	}
}
