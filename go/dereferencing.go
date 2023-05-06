// We know that addresses are where values are stored and pointers allow us to keep track of addresses. But what if we want the address to store a different value? Well, we can actually use our pointer to access the address and change its value! This action is called dereferencing or indirecting.

// We’ll need to use the * operator again on a pointer and then assign a new value like so:

// lyrics := "Moments so dear"
// pointerForStr := &lyrics

// *pointerForStr = "Journeys to plan"

// fmt.Println(lyrics) // Prints: Journeys to plan
// In our example, we have our variables: lyrics that has the value of "Moments so dear" and pointerForStr which is a pointer for lyrics. Then we use the * operator on pointerForStr to dereference it and assign a new value of "Journeys to plan". When we check the value of lyrics it’s now "Journeys to plan"!

package main

import "fmt"

func main() {
	star := "Polaris"

	starAddress := &star

	// Add your code below:
	*starAddress = "Sirius"

	fmt.Println("The actual value of star is", star)
}
