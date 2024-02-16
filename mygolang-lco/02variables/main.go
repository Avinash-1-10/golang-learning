package main

import "fmt"

// constant
const Token string = "0xdlw334bndsf93l2l3l2" // public

func main() {
	// String
	var username string = "hitesh"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	// boolean
	var isLoogedIn bool = true
	fmt.Println(isLoogedIn)
	fmt.Printf("Variable is of type: %T \n", isLoogedIn)

	// uint
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	// fload
	var smallFloat float32 = 255.786392435
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit type
	var email = "john@gmail.com"
	fmt.Println(email)

	// no var style
	numberOfUser := 4500
	fmt.Println(numberOfUser)

	fmt.Println(Token)
}
