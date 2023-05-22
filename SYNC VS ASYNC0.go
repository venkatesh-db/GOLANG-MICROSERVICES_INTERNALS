package main

import (
	"fmt"
	"time"
)

/* '' SYNC CODE Function + OOPS { struct, method , interface} ''*/

func sync() {

	// statemennt1
	// statemennt7 with condition
	// statemennt10

	fmt.Println("sync function ")

} // release memory stack

/* Devloper allocates  memory for sync concepts */

/*'' A SYNC CODE Function => goroutine ''*/

func async() {

	// statemennt1
	//  statemennt7 with condition
	//  statemennt10

	fmt.Println("async function ")

}

func main() { // main_main

	go async() // stack is allocated by os

	sync() // stack + heap + register

	time.Sleep(1)

}
