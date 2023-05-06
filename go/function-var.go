// While variables and their values are scoped to their functions, we do have built-in ways of passing information out of their native functions and into another namespace. Let’s describe the way that information can be sent from within a function to the call site, the place where the function is called. This is done by returning a value — when we return a value, we pass the value to another place in our code. A function can be given a return type, the type of a value that will be returned by the function. At the call site, the return value can be stored within a variable of the same type as the function’s return.

// func getLengthOfCentralPark() int32 {
//   var lengthInBlocks int32
//   lengthInBlocks = 51
//   return lengthInBlocks
// }
// Above, we wrote the function getLengthOfCentralPark(), we can also decide the return type by adding a type after the set of parentheses. In this case, our function has a return type of int32. Then, inside the function, we declare a variable lengthInBlocks with a value of 51. In our last line, we have our return statement. A return statement tells the function to pass back a value (or multiple values) and stops the function from executing any more code, i.e. if we put more code after our return statement, it would not run! Our function is all set up, we need to now call it in main():

// func main() {
//   var centralParkLength int32
//   centralParkLength = getLengthOfCentralPark()
//   fmt.Println(centralParkLength) // Prints: 51
// }
// Inside main() we were able to create a variable centralParkLength with type int32 and store the result (the returned value) from getLengthOfCentralPark() into centralParkLength. Then we can check the value of centralParkLength by printing it, which confirms what we said about return by printing the number 51. Even though we can’t access lengthInBlocks from getLengthOfCentralPark() directly, we can access the information we need through the return keyword!

package main

import (
	"fmt"
	"time"
)

// Add "string" as the return type of this function
func isItLateInNewYork() string {
	var lateMessage string
	t := time.Now()
	tz, _ := time.LoadLocation("America/New_York")
	nyHour := t.In(tz).Hour()
	if nyHour < 5 {
		lateMessage = "Goodness it is late"
	} else if nyHour < 16 {
		lateMessage = "It's not late at all!"
	} else if nyHour < 19 {
		lateMessage = "I guess it's getting kind of late"
	} else {
		lateMessage = "It's late"
	}

	// Return the string lateMessage
	return lateMessage
}

func main() {
	var nyLate string
	nyLate = isItLateInNewYork()
	fmt.Println(nyLate)
}
