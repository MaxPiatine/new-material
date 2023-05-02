// Inferring Type
// There is a way to declare a variable without explicitly stating its type using the short declaration := operator. We might use the := operator if we know what value we want our variable to store when creating it. For instance:

// nuclearMeltdownOccurring := true
// radiumInGroundWater := 4.521
// daysSinceLastWorkplaceCatastrophe := 0
// externalMessage := "Everything is normal. Keep calm and carry on."
// Above, we were able to define a bool, a float, an int, and a string without specifying the type. We used the := to create a variable and infer its type based on the value provided. Floats created in this way are of type float64. Integers created in this way are either int32 or int64 (we’ll discuss how this is determined in the next exercise).

// Go also offers a separate syntax to declare a variable and infer its type. We could’ve written the same code from above as:

// var nuclearMeltdownOccurring = true
// var radiumInGroundWater = 4.521
// var daysSinceLastWorkplaceCatastrophe = 0
// var externalMessage = "Everything is normal. Keep calm and carry on."
// Notice, in the second example, that we used the var keyword and the = operator. In both examples, we’ve declared and initialized variables while leaving the Go compiler to infer the type of the value assigned.

// Instructions
// 1.
// Create an integer called daysOnVacation using the := operator and assign it a number between 0 to 9.

package main

import "fmt"

func main() {
  // Define daysOnVacation using := below:
  daysOnVacation := 9
  
  // Define hoursInDay using var and = below:
  var hoursInDay = 24
  
  fmt.Println("You have spent", daysOnVacation * hoursInDay, "hours on vacation.")
}