// In this exercise, we’ll focus on the first two, && and ||. When we use the And (&&) operator, we are checking that both expressions are true:

// if storeLights == "on" && doorsOpen {
//   fmt.Println("You can enter the store!")
// }
// When using the && operator, both conditions must evaluate to true for the entire condition to evaluate to true and execute. Otherwise, if either condition evaluate as false, the && condition will evaluate to false and the code inside the if block will not execute.

// If we only care about either condition being true, we can use the Or (||) operator:

// if day == "Saturday" || day == "Sunday" {
//   fmt.Println("Enjoy the weekend!")
// } else {
//   fmt.Println("Do some work.")
// }
// When using the || operator, only one of the conditions must evaluate to true for the overall statement to evaluate to true. In the code example above, if either day == "Saturday" or day == "Sunday" is true the statement’s code block will execute. If the first operand in the || expression evaluates to true, the second operand won’t even be checked. Only if day == "Saturday" evaluates to false will day == "Sunday" be evaluated. The code in the else statement above will execute only if both comparisons evaluate to false.

// Let’s implement these operators to add an additional layer of logic!

package main

import "fmt"

func main() {
	rightTime := true
	rightPlace := true

	// Edit this condition for the FIRST checkpoint
	if rightTime && rightPlace {
		fmt.Println("We're outta here!")
	} else {
		fmt.Println("Be patient...")
	}

	enoughRobbers := false
	enoughBags := true

	// Edit this condition for the SECOND checkpoint
	if enoughRobbers || enoughBags {
		fmt.Println("Grab everything!")
	} else {
		fmt.Println("Grab whatever you can!")
	}
}
