//Stampare il più grande valore di tipo int64

package main;

import "fmt";

func main (){
	var x , y int64;
	x = 1024;
	y = x * x * x; // --> 2^30 poiché 1024 è = a 2^10, quindi moltiplico x 3 volte
	fmt.Println("max int64:" , y * y * 2 * 2 * 2 - 1); // y * y = 2^60 * 2 * 2 = 2^64
	
}