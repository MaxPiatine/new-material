// In Go, values can be many things. Just to name a few, values can be numbers (like 109), or text wrapped in quotes (like "Hello world"). These values can be written into code as is, and are called literals. They are literally what they say they are.

// We can perform arithmetic in Go with literals (or named values, covered in the next exercise) using the following operators:

// + to add
// - to subtract
// * to multiply
// / to divide
// % to take the remainder (the modulus operator) between two numbers.
// fmt.Println(20 * 3) // Prints: 60
// fmt.Println(55.21 / 5) // Prints: 11.042
// fmt.Println(9 % 2) // Prints: 1
// Imagine the code above as appearing inside the main function body of a Go program. In this snippet, we used the Go programming language as a calculator. We printed out the product of multiplying 20 and 3 (60). We next printed out the quotient of dividing 55.21 by 5 (11.042). We lastly printed out the remainder left over after dividing 9 by 2 (9 divided by 2 has a remainder of 1).

package main

import "fmt"

func main() {
  fmt.Println(2235 * 1231)
}