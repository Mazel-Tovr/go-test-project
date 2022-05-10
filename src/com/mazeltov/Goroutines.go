package main

import "fmt"

func testGoroutines(i int) {
	fmt.Println("Iteration num ", i)
}

func tests() {
	for i := 0; i < 10; i++ {
		go testGoroutines(i) // will be executed suspend
	}
	fmt.Scanln()
}

func main() {
	channelSimple()
}

func channelSimple() {
	ch1 := make(chan int) // creating channel
	go func(in chan int) {
		val := <-in // reading from channel
		fmt.Println("GO: get from chan", val)
		fmt.Println("GO: after read from chan")
	}(ch1)
	ch1 <- 42 // writing to  from channel
	//ch1 <- 21 // dead lock we should make buffered channel make(chan int, 10)
	fmt.Println("MAIN: after put to chan")
	fmt.Scanln()
}
