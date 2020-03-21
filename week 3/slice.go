package main
import (
	"fmt"
	"strconv"
	"sort"
)
func main(){
	x := make([]int,3);
	var y string;
	i := 0;
	fmt.Println("Enter an integer");
	fmt.Scan(&y);
	for y != "x" && y != "X" {
		k, err := strconv.Atoi(y);
		if i < 3 && err == nil{
			if i == 2{
				x[0] =k;
			}else{
				x[i] = k;
			}
			i++;
		}else if err == nil{
			x = append(x,k);
		}
		if err != nil{
			fmt.Println("It is not an integer", y);
		}
		sort.Ints(x);
		fmt.Println(x);
		fmt.Println("Enter an integer");
		fmt.Scan(&y);
	}
}