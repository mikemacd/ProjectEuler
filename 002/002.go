package main

import "fmt"

// By considering the terms in the Fibonacci sequence whose values do not exceed four million, 
// find the sum of the even-valued terms.

func main(){

	sum:=0 
	f1 := 1 // second most recently seen Fibonacci number
	f2 := 1 // most recently seen Fibonacci number
	for f2 < 4000000 {
		// fmt.Printf("f1: %d f2: %d",f1,f2)
		if f2%2==0 {
			sum+=f2
		}
		f2,f1=f2+f1,f2
		 
		// fmt.Printf(" sum: %d \n",sum)
	}
	fmt.Printf("Sum: %d \n",sum)
}
