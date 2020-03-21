package main
import (
	"fmt"
	
)
// func main() {
// // 	i, _ := strconv.Atoi("10")
// // 	y := i * 2
// // 	fmt.Println(y)
// //   }
// s := strings.Replace("ianianian", "ni", "in", 2)
//   fmt.Println(s)

// }
func main() {
	s := make([]int, 0, 3)
	s = append(s, 100)
	fmt.Println(len(s), cap(s))
  }

// func main() {
// 	// var xtemp int
// 	// x1 := 0
// 	// x2 := 1
// 	// for x:=0; x<5; x++ {
// 	//   xtemp = x2
// 	//   x2 = x2 + x1
// 	//   x1 = xtemp
// 	// }
// 	// fmt.Println(x2)
// 	func main() {
// 		x := [...]int {4, 8, 5}
// 		y := x[0:2]
// 		z := x[1:3]
// 		y[0] = 1
// 		z[1] = 3
// 		fmt.Print(x)
// 	  }
//   }