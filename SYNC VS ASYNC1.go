package main // main goroutine

import (
	"fmt"
)

/* ''  SYNC CODE Function + OOPS { struct, method , interface}
       CODE EXECUTION DEPENDS ON THE DATA INSIDE THE MEMORY
''*/

func synccopy(fees int) {

	fees = 85000

	if fees <= 80000 {
		panic("not allowed for course")
	}

}
func sync() {

	var course string = "golang"
	var fees int

	if course != "golang" {
		fees = 75000
	}

	synccopy(fees)

	fmt.Println("sync function ")

} // release memory stack

/* Devloper allocates  memory for sync concepts */

/*'' A SYNC CODE Function => goroutine
       CODE EXECUTION DEPENDS ON THE DATA INSIDE THE CHANNEL MEMORY { HEAP MEMORY}
	   the channel memory is not residing in the same function outside
''*/

func async(cha chan string) {

	fmt.Println("async function started ")

	var course string = <-cha
	var fees int

	if course != "golang" {
		fees = 75000
	}

	synccopy(fees)

	fmt.Println("async function end ")

}

func main() { // main_main

	var cha chan string = make(chan string)

	go async(cha) // stack is allocated by os

	sync() // stack + heap + register

	fmt.Println("before writing in to channel ")

	cha <- "golang"

	fmt.Println(" end of main")
}
