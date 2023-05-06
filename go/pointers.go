// In the previous exercise we went over addresses, now let’s learn how to store them. In Go, pointers do that job for us. Pointers are variables that specifically store addresses.

// We even set the data type of the addresses’ value for the pointer. For instance:

// var pointerForInt *int
// In the example above pointerForInt will store the address of a variable that has an int data type. To break it down further, the * operator signifies that this variable will store an address and the int portion means that the address contains an integer value.

// With pointerForInt initialized, we can assign it value like so:

// var pointerForInt *int

// minutes := 525600

// pointerForInt = &minutes

// fmt.Println(pointerForInt) // Prints 0xc000018038
// Notice in our example that minutes has a value of 525600, an integer type. Since we’ve initialized pointerForInt as a pointer that will hold the address of an integer value, we can then assign the address of minutes (&minutes) to pointerForInt. Printing out pointerForInt, we get another hexadecimal number: 0xc000018038.

// We can also declare a pointer implicitly like other variables:

// minutes := 55

// pointerForInt := &minutes
// Let’s see how we would create a pointer for a string instead!

package main

import "fmt"

func main() {
	star := "Polaris"

	// Add your code below:
	starAddress := &star
	fmt.Println("The address of star is", starAddress)

}
