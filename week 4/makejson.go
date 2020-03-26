package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var name, address string

	fmt.Print("Enter the name: ")
	fmt.Scan(&name)
	// json1["name"] = s
	fmt.Print("Enter the address: ")
	fmt.Scan(&address)
	// json1["address"] = s
	json1 := map[string]string{
		"name":    name,
		"address": address,
	}
	// fmt.Print(json1)
	jsonObject, err := json.Marshal(json1)
	if err == nil {
		fmt.Print("\n\n\nPrinting the Marshalled data: \n")
		os.Stdout.Write(jsonObject)
		json2 := make(map[string]string)
		err1 := json.Unmarshal(jsonObject, &json2)
		if err1 == nil {
			fmt.Print("\n\n\nPrinting the Unmarshalled data:\n")
			fmt.Print(json2)
		}

	} else {
		fmt.Print(err)
	}

}
