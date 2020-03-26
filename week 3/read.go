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
		str := strings.Fields(fileScanner.Text())
		fileTextLines = append(fileTextLines, Name{
			fname: str[0],
			lname: str[1],
		})
	}
	for _, str := range fileTextLines {
		fmt.Print("First Name: ", str.fname, "  Last Name: ", str.lname, "\n")
	}
}
