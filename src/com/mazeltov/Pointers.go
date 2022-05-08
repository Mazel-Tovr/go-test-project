package main

func pointers() {

	a := 2

	var b *int = &a // b - is a pinter to a

	//*b - we set value for b
	*b = 3 // a value will be 3

	c := &a // new pointer to a

	//var d *int
	d := new(int) // we create new pointer ( default value for int = 0 )

	*d = 12 // change default value of d (0) to 12

	// operating with value only
	*c = *d // c value = 12 , so a value = 12 ; cos c - is pointer to a

	*d = 15 // only d value will be changed

	c = d // we are changing pointers; a no longer have pointers

	*c = 14 // c value = 14 , d value = 14, a doesn't change = 12

}
