package main

import "fmt"

var LoginToken = 30000 // Public variable

func main() {
	var username string = "Vijay"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isBool bool = true
	fmt.Println(isBool)
	fmt.Printf("Variable is of type: %T \n", isBool)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.65465451384364545
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	//default values of declaritive variables
	var anotherVar int
	fmt.Println(anotherVar)
	fmt.Printf("Type is %T \n", anotherVar)

	var anotherStr string
	fmt.Println(anotherStr)
	fmt.Printf("Type is %T \n", anotherStr)

	fmt.Println(LoginToken)
	fmt.Printf("type is %T", LoginToken)

}
