#Program to Read Names from a Text File
This program reads a text file containing a series of names and creates a slice of name structs, where each struct represents a single name from the file. The program then prints the first name and last name attributes of each struct to the console.

##Getting Started
To use this program, you will need to have Go installed on your computer. If you don't have Go installed, you can download it from the official website.

##Usage
Open a terminal or command prompt.
Navigate to the directory where the program file is located.
Type go run main.go to run the program.
When prompted, enter the name of the text file containing the names.
##Example Input File
Here's an example of what the input file should look like:

Copy code
John Smith
Jane Doe
Bob Johnson
Alice Williams
Each line of the text file should contain a first name and a last name, separated by a space.

###Output
The program will create a slice of name structs, where each struct represents a single name from the input file. The program will then print the first name and last name attributes of each struct to the console, like this:

sql
Copy code
First name: John, Last name: Smith
First name: Jane, Last name: Doe
First name: Bob, Last name: Johnson
First name: Alice, Last name: Williams
