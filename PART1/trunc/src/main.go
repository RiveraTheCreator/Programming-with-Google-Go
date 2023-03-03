package main

import(
	"fmt"
)
	
func main(){
	numf := 0.00
	fmt.Print("Put a float number:  ")
	fmt.Scan(&numf)
	fmt.Print("\nThe number trunked is: ", int64(numf))
	
}