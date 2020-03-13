package main
import (
	"fmt"
	"bufio"
	"os"
)
func main(){
	consoleReader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter the string: ");
	str, _ := consoleReader.ReadString('\n')
	if test(string(str)) {
		fmt.Print("Found");
	}else{
		fmt.Print("Not Found");
	}
	
	

}

func test(str string) bool{
	// fmt.Print(len(str)-3);
	if str[0] == 105  && str[len(str)-3] == 110{
		var i int;
		for i = 1 ; i < len(str) - 3 ; i++{
			if str[i] == 97 {
				return true	
			}
		}
	}
	return false
	
}