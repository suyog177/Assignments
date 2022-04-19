package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter the string :")
	var str string
	fmt.Scanln(&str)
if validIP(str) {
	fmt.Println("Entered IP Address is valid")
}else {
	fmt.Println("Entered IP Address is invalid")
}
}

func validIP(str string) bool {
	tempArr := strings.Split(str, ".")

	if len(tempArr) != 4 {
		return false
	}
	for _, obj := range tempArr {
		if val, err := strconv.Atoi(string(obj)); err != nil {
			return false
		} else if val < 0 || val > 255 {
			return false
		}
	}

	return true
}


