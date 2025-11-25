package main;
import "fmt";
import "math"

func main(){
	var n, l int;
	fmt.Print("Inserire il numero di lati del poligono: ");
	fmt.Scan(&n);
	fmt.Print("Inserire la lunghezza dei lati del poligono: ");
	fmt.Scan(&l);

	area := float64(float64(n) * math.Pow(float64(l),2)) / (4 * math.Tan(math.Pi / float64(n)));
	fmt.Printf("L'area del poligono regolare è = %.4f\n" , area);
}