package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	intSlice := make([]int, 0, 3) // create a slice

	for {
		var input string
		fmt.Println("Pur a integer (or 'X' for out): ")
		fmt.Scanln(&input)

		if input == "X" {
			break // break the loop
		}

		num, err := strconv.Atoi(input) 
		if err != nil {
			fmt.Println("put a valid input")
			continue 
		}

		intSlice = append(intSlice, num) 
		sort.Ints(intSlice)              

		fmt.Println(intSlice) 
	}
}
