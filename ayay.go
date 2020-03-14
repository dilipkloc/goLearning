package main
import (
	"fmt"
)

func main(){
	var x [10]float32;
	var y = [5]int{1,2,3,4,5};
	z := [...]int{4,56,32,3,42,12};
	fmt.Println(z[5]);
	fmt.Println(y[4]);

	// Looping through array
	for index,value := range z{
		fmt.Print(index);
		fmt.Print(":");
		fmt.Println(value);
	}

	// Below array initialization will not work
	// var z [5]int = [5]{1,2,3,4,5};
	x[0] = 1;
	fmt.Print(x[9]);
	

}