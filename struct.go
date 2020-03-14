package main
import "fmt"



type Person struct {
	name string
	addr string
	phone int
	idMap map[string]int
}

func main(){
	var p1 Person;
	p1.name = "dilip"
	p1.addr = "bangalore"
	p1.phone = 8722255043

	fmt.Println(p1.phone)


	// make should be used inorder use the map from struct
	p1.idMap = make(map[string]int)

	p1.idMap[p1.name] = 1

	fmt.Println(p1.idMap[p1.name])

	// initializing struct
	// this method initialisez all the values to zero int, empty for string(single space)
	p2 := new(Person)
	fmt.Println(p2.name,"hi")

	// can initialize values while creating
	p3 := Person{name:"joe",
	addr:"white field",
	phone:8755522654,
	idMap : make(map[string]int)}
	
	p3.idMap["dilip"] = 213

	fmt.Println(p3.name)

	fmt.Println(p3.idMap["dilip"])



}