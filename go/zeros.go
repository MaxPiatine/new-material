// Even before we assign anything to our variables they hold a value. Go’s designers attempted to make these “sensible defaults” that we can anticipate based on the variable’s types. All numeric variables have a value of 0 before assignment. String variables have a default value of "", which is also known as the empty string because it contains no characters. Boolean variables have a default value of false. For example:

// var classTime uint32
// var averageGrade float32
// var teacherName string
// var isPassFail bool
 
// fmt.Println(classTime) // Prints 0
// fmt.Println(averageGrade) // Prints 0
// fmt.Println(teacherName) // Doesn't print anything
// fmt.Println(isPassFail) // Prints false
// Above we declared four variables, an unsigned 32-bit int, a 32-bit floating point number, a string, and a boolean. Without assigning any of the variables to a value we print them out to see their default value. The two numbers print the same result, 0, a valid value for both types. The empty string, when printed, displays nothing. The boolean value prints false.

package main

import "fmt"

func main() {
  // Create three variables
  // emptyInt an int8
  var emptyInt int8
  // emptyFloat a float32
  var emptyFloat float32
  // and emptyString a string
  var emptyString string
  // Finally, print them all out
  fmt.Println(emptyInt, emptyFloat, emptyString)
}