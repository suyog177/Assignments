package main

import "fmt"

func main() {
	//declared the map for storing the student and scores respectively
	Score := map[string]int{}

	//for getting the number of students user wants to enter dynamically
	var i int 
	fmt.Println("How many student's score you want to enter :")
	fmt.Scanln(&i)
	
	//iteration for the student name and marks 
	for j:=1;j<=i;j++{
	fmt.Println("Enter student name :")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Enter marks of ",name)
	var marks int
	fmt.Scanln(&marks)
		
	//stored data in a map
	Score[name] = marks
}
fmt.Println("Team      Grades")

	//iteration on the score map for assigning the grade to students and displaying it as required
	for name, marks := range Score {
		switch {
		case marks >= 0 && marks <= 49:
			fmt.Println(name,"      ","F")
		case marks >= 50 && marks <= 59:
			fmt.Println(name,"      ","E")
		case marks >= 60 && marks <= 69:
			fmt.Println(name,"      ","D")
		case marks >= 70 && marks <= 79:
			fmt.Println(name,"      ","C")
		case marks >= 80 && marks <= 89:
			fmt.Println(name,"      ","B")
		case marks >= 90 && marks <= 100:
			fmt.Println(name,"      ","A")
		default:
			break
		}
	}

}
