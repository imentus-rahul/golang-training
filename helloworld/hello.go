package main

import "fmt"
// import "rsc.io/quote"


// func main() {
//     // fmt.Println("Hello, World!")
// 	fmt.Println(quote.Go())
// }

import "example.com/greetings" // changed to local package in go.mod
func main() {
    // Get a greeting message and print it.
    message := greetings.Greet("Gladys") // calling method from external package
    fmt.Println(message)
}