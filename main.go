package main

import (
	"fmt"
	
	"time"
)

func main() {

	for {
		fmt.Println("===========================================================")
		fmt.Println("Welcome To Adewah Global Tech. ")
		fmt.Println("current:", time.Now().Format("15: 00 pm"))
		fmt.Println("===========================================================")
		fmt.Println("Enter Your FullName: ")
		var result string
		processor(reader(inputs(result)))

		
	}

}
