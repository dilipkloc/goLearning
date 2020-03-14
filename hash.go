package main
import "fmt"
func main(){
	// map is hash in golang

	// to create map, 
	// var idMap map[string]int;
	idMap1 := make(map[string]int);

	// initialize map
	idMap2 := map[string]int {
		"dilip":1}
	fmt.Println(idMap2["dilip"])
	// adding a new key and value

	idMap1["kumar"] = 34

	fmt.Println(idMap1["kumar"])

	// delete a key

	delete(idMap1,"kumar")


	fmt.Println(idMap1["kumar123123"])

	// below return id=> value and p boolean if the key exists or not
	id,p := idMap1["kumar"]

	fmt.Println(p)
	fmt.Println(id)

	fmt.Println(len(idMap2))


	idMap3 := map[string]int{
		"dilip":123,
		"kumar":234,
		"m":354}

	for key, val := range idMap3 {
		fmt.Println(key,":",val);
	}


}