package main /*
main - main package, for run ur programme
ur go file must be with main package and have main fun
*/

import (
	"fmt"
)

type MySuperMegaTyp int // type - like typealiase in kotlin

func main() {
	defer variables() // will be invoked in the end
	anonFun := func(x int, z int) int { return x + z }
	funWithAnonFun("Hello", anonFun)
	withCircle()

	var variable MySuperMegaTyp = 1
	fmt.Println(variable)
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

func funWithAnonFun(x string, fun func(int, int) int) {
	res := fun(1, 4)
	fmt.Println(fmt.Sprintf("String %s Int %d", x, res))
}

func withCircle() {
	var counter int

	for i := 0; i > 10; i++ {
		counter++
	}

	names := [...]string{"Sanya", "Egor", "Roma"}

	for index, value := range names {
		fmt.Println(fmt.Sprintf("Index %d; Value %s", index, value))
	}
}

type user struct {
	name string
	age  int
}

func objectOperation() {
	person := user{}

	person.age = 10
	person.name = "SerGay"

}
