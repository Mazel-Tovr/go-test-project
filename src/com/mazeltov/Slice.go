package main

import "fmt"

func slice() {

	// it's an array
	var array = [1]int{1}

	//slice it's like array, but with capacity
	var slice = []int{}             //lent=0, cap=0
	var slice1 = []int{1}           //lent=1, cap=1
	var slice2 = make([]int, 5)     //lent=5, cap=5
	var slice3 = make([]int, 5, 10) //lent=5, cap=10

	fmt.Println(array, slice, slice1, slice2, slice3)

	// adding elements to slice

	var buf []int //lent=0, cap=0

	buf = append(buf, 1, 2) //lent=2, cap=2
	buf = append(buf, 5)    //lent=3, cap=4

	// slice merging

	otherSlice := make([]int, 3) // [0,0,0]

	buf = append(buf, otherSlice...) //lent=6, cap=8

	// check len and cap

	var bufLen, bufCap = len(buf), cap(buf)

	fmt.Println(bufLen, bufCap)

}

func sliceCut() {
	slice := []int{1, 2, 3, 4, 5}

	sl1 := slice[1:4] // [2,3,4] (slice[1]  slice[2]  slice[3])
	sl2 := slice[:2]  // [1,2] (slice[0]  slice[1])
	sl3 := slice[2:]  // [3, 4, 5] (slice[2]  slice[3]  slice[4])

	fmt.Println(sl2, sl3, sl1)

	// when we cut slices, value are still pointing on origin slice

	newSlice := slice[:]

	newSlice[0] = 6 // slice will be [6, 2, 3, 4, 5]

	//to avoid this you must copy slices

	emptySlice := make([]int, len(slice), cap(slice))

	copy(emptySlice, slice)
}
