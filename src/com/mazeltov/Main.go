package main /*
main - main package, for run ur programme
ur go file must be with main package and have main fun
*/

import (
	"fmt"
)

func main() {
	variables()
}

func variables() {
	/*
		unused variables is NOT allowed
		:= - init variable with some value
		 = - set new value in all read inited variable
	*/

	// Examples
	helloWorldMessage := "Hello world"
	fmt.Println(helloWorldMessage)

	helloWorldMessage = "New value setted"
	fmt.Println(helloWorldMessage)

	// creation late init variable
	var lateInitVar string

	lateInitVar = "late init"

	fmt.Println(lateInitVar)

	/*
		var lateInitVar string = "init"
		same as shit
		lateInitVar := "init"
	*/

	/*
		types https://metanit.com/go/tutorial/2.3.php
		rune - char
	*/
	var iAmChar rune = 'h'
	fmt.Println(iAmChar)

	/*
	  Multiple initializing
	*/
	a, b, c := 1, 2, 3

	// a and b switching values
	a, b = b, a

	// c get new value
	_, _, c = 1, c, 4
	// must be  2  1  4
	fmt.Println(a, b, c)
}

// Constant can be unused
const constantName string = "I am constant"

// Global var can be also unused
var globVar string
