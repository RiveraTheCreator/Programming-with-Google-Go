
package main
/* Description
Write a program which prompts the user to enter a string. 
The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’. 
The program should print “Found!” if the entered string starts with the character ‘i’, ends with the character ‘n’, and contains the character ‘a’. 
The program should print “Not Found!” otherwise. 
The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.

*/

import(
	"fmt"
	"strconv"
	
)

func main() {
	i, _ := strconv.Atoi("10")
	y := i * 2
	fmt.Println(y)
  }