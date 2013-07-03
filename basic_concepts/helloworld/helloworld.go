// This is a comment!  
/* Comments in Go can be added using // if it is a single line.
   Multiline comments can be added by using the syntax in this comment! */

// Below is a package declaration.
// Packages in Go are used as libraries to import libraries.
package main

// import is used to include public functions and constants
// from the "fmt" package.  "fmt", short for format, is used
// to format input and output
import "fmt"

// Function declaration is declared by using the func keyword
// followed by the name of the function (main in this case).
// main is a special function. It is called when our program
// is executed.
func main(){
  fmt.Println("Hello World!")
}

// The statement inside of the main function is made up of three parts:
//
// fmt - Declare that we are using the fmt package
// Println - Println is a function inside the fmt package used to print a line of text
// ("Hello World!") - The String that will be returned from Println
