package main
import (
	"fmt"
)

func main(){
	x := []string{"a","b"};
	fmt.Println(len(x));
	fmt.Println(cap(x));
	// make() is used to create an slice variable
	sli := make([]int,1,2);
	// here only 2 arguements, 10 is the length and capacity
	sli[0] = 10;
	fmt.Println(len(sli))
	sli = append(sli,20)
	fmt.Println(len(sli))
	fmt.Println(sli[len(sli)-1])
	sli = append(sli,30)
	fmt.Println(sli[len(sli)-1])
	sli = append(sli,40)
	fmt.Println(sli[len(sli)-1])

	fmt.Println(sli)
	// sli1 := make([]int,5,20);
	// if 2 arguements are passed, 5 is the length, capacity is the 20
	// fmt.Println(len(x))
	// for check := 0 ; check <= (len(x)-1) ; check++ {
	// 	// fmt.Println(check)
	// 	fmt.Print(x[check])
	// }

}