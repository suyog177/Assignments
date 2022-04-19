package main

import (
	"fmt"
	"strconv"
)

//function to return the first integer
func printInt(str string) int {

	for _, obj := range str {
		if val, err := strconv.Atoi(string(obj)); err == nil {
			return val
		}
	}

	return 0
}

func main() {
	fmt.Println("Enter the string :")
	var str string
	fmt.Scanln(&str)
	if printInt(str) == 0 {
		fmt.Println("No integer found in given string")
	} else {
		fmt.Println(printInt(str))
	}
}
