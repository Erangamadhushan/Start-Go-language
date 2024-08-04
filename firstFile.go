package main
import(
	"fmt"
)

//how run and build go progrmme into vs code terminal(instruction below )
func main() {
	//for get output we can use 'fmt.Println(output)'
	//below given example about me 
	fmt.Println("My name is  Eranga Madhushan");
	fmt.Println("I'm also BCS Student at University of Ruhuna in Sri Lanka");

	//Declare variable with 'var' keyword

	/*
		syntax
			var variableName variableType = value
	*/

	//Mainly Go have 4 variable types
		//1.int
		//2.float32 or float64
		//3.string
		//4.bool
	var myInt int = 956
	var myFloat float32 = 1.234
	var myString string = "EM956"
	var myBool bool = true

	//print variable outputs
	fmt.Println(myInt)
	fmt.Println(myFloat)
	fmt.Println(myString)
	fmt.Println(myBool)

	/*

		=====  OUTPUT  =====

		My name is  Eranga Madhushan
		I'm also BCS Student at University of Ruhuna in Sri Lanka
		956
		1.234
		EM956
		true
	*/
}
//===============================================================================================
//    Instruction

// first you check go version in your device
//	open cmd check go version 'go version'
//if you got the answer you can run and build go programme via vs code terminal

// 'go run .\filename.go
// 'go build .\filename.go