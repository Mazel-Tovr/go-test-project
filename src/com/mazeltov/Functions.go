package main

import "fmt"

func funWithOut(num int) (out int, z int) {
	out = 1
	z = 2
	return
}

// same
func funWithOut2(num int) (out int, z int) {

	return 1, 2
}

func funWithError() (int, error) {
	return 0, fmt.Errorf("smt happend")
}

func funWithVararg(nums ...int) {
	fmt.Println(nums)
}

func test() {
	slice := []int{1, 2, 4}
	funWithVararg(slice...) // need to unpack slice
}
