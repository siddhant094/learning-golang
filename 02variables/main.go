package main

import "fmt"

func main() {
	var isLoggedin int = 256
	fmt.Println(isLoggedin)
	fmt.Println("Variables")
	fmt.Printf("Variable is of type: %T \n", isLoggedin)

	var num2 float64 = 255.234234234234
	fmt.Println(num2)
	fmt.Println("Variables")
	fmt.Printf("Variable is of type: %T \n", isLoggedin)
	
	//default values and some aliases
	var newVar string
	fmt.Println(newVar)
	fmt.Printf("Variable is of type: %T \n", newVar)

	//implicit type
	var website = "learncode"
	fmt.Println(website)

	noOfUser := 300000 
	fmt.Println(noOfUser)
}

