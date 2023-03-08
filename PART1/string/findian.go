
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
	"strings"
)

func main(){
	var data string
	fmt.Print("Put the string:\n")
	_,err := fmt.Scan(&data)

	if(err != nil){fmt.Print("Some errors here:  ",  err)}

	data = strings.ToLower(data)

	if strings.Index(data,"i") == 0 && strings.Index(data,"a") != 0 && strings.HasSuffix(data,"n") {
		fmt.Print("Found!")		
	}else{
		fmt.Print("Not Found!")

	}

}