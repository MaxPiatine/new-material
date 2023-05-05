// else if statements are great for checking multiple conditions. However, we can find ourselves needing to check so many conditions that writing all the necessary else if statements can feel tedious. For example:

// clothingChoice := "sweater"

// if clothingChoice == "shirt" {
//   fmt.Println("We have shirts in S and M only.")
// } else if clothingChoice == "polos" {
//   fmt.Println("We have polos in M, L, and XL.")
// } else if clothingChoice == "sweater" {
//   fmt.Println("We have sweaters in S, M, L, and XL.")
// } else {
//   fmt.Println("Sorry, we don't carry that.")
// }
// In the code above, we have a series of conditions checking for a value that matches a clothingChoice variable. Our code works fine, but imagine if we needed to check more values! Having to write additional else if statements sounds like a pain!

// Instead, we can use a switch statement that uses alternative syntax that is easier to read and write. Below is an example of a switch statement:

// clothingChoice := "sweater"

// switch clothingChoice {
// case "shirt":
//   fmt.Println("We have shirts in S and M only.")
// case "polos":
//   fmt.Println("We have polos in M, L, and XL.")
// case "sweater":
//   fmt.Println("We have sweaters in S, M, L, and XL.")
// case "jackets":
//   fmt.Println("We have jackets in all sizes.")
// default:
//   fmt.Println("Sorry, we don't carry that")
// }
// // Prints: We have sweaters in S, M, L, and XL.
// Let’s break down what happens in a switch statement:

// The switch keyword initiates the statement and is followed by a value. In the example, the value after switch is compared to the value after each case, until there is a match.
// Inside the switch block, { ... }, there are multiple cases. The case keyword checks if the expression matches the specified value that comes after it. The value following the first case is "shirt". Since the value of clothingChoice is not the same as "shirt", that case‘s code does not run.
// The value of clothingChoice is "sweater", so the third case‘s code runs and "We have sweaters in S, M, L, and XL." is printed. No more case statements are checked.
// At the end of our switch statement is a default statement. If none of the cases are true, then the code in the default statement will run.

package main

import "fmt"

func main() {
	name := "H. J. Simp"

	// Add your switch statement below:
	switch name {
	case "Butch":
		fmt.Println("Head to Robbers Roost.")
	case "Bonnie":
		fmt.Println("Stay put in Joplin.")
	default:
		fmt.Println("Just hide!")
	}
}
