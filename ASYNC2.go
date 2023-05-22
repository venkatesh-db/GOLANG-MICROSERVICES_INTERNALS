package main // main goroutine

/* ''  SYNC CODE Function + OOPS { struct, method , interface}
       CODE EXECUTION DEPENDS ON THE DATA INSIDE THE MEMORY
''*/

func synccopy(fees int) {


	// SYNC FUNCTIONS ==> Isolated address space in ram
	// RAM -> stack ,heap ,data
	// Always avialble for allocations

	// What if we do this operations in  sync function somes time it will pause
	//  RESOURCE -> CHAN , IO operations { iobuffer} ,FILE

}

/* Devloper allocates  memory for sync concepts */

/*'' A SYNC CODE Function => goroutine
       CODE EXECUTION DEPENDS ON THE DATA INSIDE THE CHANNEL MEMORY { HEAP MEMORY}
	   the channel memory is not residing in the same function outside
''*/

func async(cha chan string) {

	
	// ASYNC FUNCTIONS in different application  ==> Isolated address space in ram
	// but to communicate between two applications we need resource 
	// 1. shared memory 2. message queue 3. pipe 

	// 1.ASYNC FUNCTIONS in one  application ==> two go routines 
	//  1. RESOURCE           ->   CHAN , IO operations { iobuffer} ,FILE
	//   1.1 HARDISK          ->   FILE

	// 2. MICROSERVICES      ->    OS MEMORY
	//2.1  OS MEMORY FOR COMMUNICATION -> computer1 --> computer2  CHAN , RABBITMQ , KAFKA, etc ..
           || 2.1 use Layered approach
	// LAYERED -> Functions  :=> write in to channel or read from channel
	// WIth out function CHANNEL ->               direct reading  data from and to a channel

	//  3. HARDISK -> DATABASE
	//  3.1 DATABASE -> MYSQL , POSTGRESS , REDIS , MONGODB
	//  3.2  DATABASE -> HUGHE FILE DATA ASYNC FUNCTIONS
	                  // TRanformation TExt file content is tranformed to map 
	//  3.4 DATATBASE -> FILE SIZE 10GB => split file data in to chunks of 
	 //         1. 500MB   2. 500MB  3. 500MB  ==> all processing using async 

	   // 4.      Client         - Server   computer1 --> computer2
	  //      application1       application2 
      //      async function      async function 
	  //    4.1. Socket -> two way communications   
	  //   4.2. Server we use async function

	// 5. UR application  ASYNC FUNCTIONS in one  application ==> two go routines 
    //    We have wait on resources -> Mulitple resources
	//   EX: vlc player volume & system volume are two different resources 
	
	//6. CodeRegistrations : 1. go routine 2. function 3. api 4. struct 
	                    //    5. method 6. interface 7. objects  

	//7. Events + CODE   : MONEYDEPOSIT+ code { curentbalance = curentbalance +MONEYDEPOSIT }				
             ||
		Application1   : can be run like vlc as three instances 

     //8.  Events + CODE   : => Application  ony instnce of application be runned Server1
	  
	                               


	res := func() {

		cha <- "string"

	}

	res()

}

func main() { // main_main

}
