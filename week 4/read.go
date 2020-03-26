package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	var fileName string
	fmt.Print("Enter the file name: ")
	fmt.Scan(&fileName)
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Print("File not found")
		os.Exit(0)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var fileTextLines []Name

	for fileScanner.Scan() {
		names := strings.Fields(fileScanner.Text())
		var name Name
		if len(names[0]) > 20 {
			name.fname = names[0][:20]
		} else {
			name.fname = names[0]
		}
		if len(names[1]) > 20 {
			name.lname = names[1][:20]
		} else {
			name.lname = names[1]
		}
		fileTextLines = append(fileTextLines, name)
	}
	for _, name := range fileTextLines {
		fmt.Print("First Name: ", name.fname, "  Last Name: ", name.lname, "\n")
	}
}
