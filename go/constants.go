// In addition to literal (aka unnamed) values, there are also named values. Naming a value in Go means creating a word that will represent that value. One example of named values are constants, which cannot be updated while the program is running. Another example of named values are variables which we can update their value. We’ll focus on constants in this exercise and discuss variables at length in later exercises.

// We use the const keyword to create a constant. We immediately assign a value to the constant using a literal. Throughout the rest of the program, we can use the constant’s name instead of the literal.

// const funFact = "Hummingbirds' wings can beat up to 200 times a second."
 
// fmt.Println("Did you know?")
// fmt.Println(funFact)
// Which prints:

// Did you know? 
// Hummingbirds' wings can beat up to 200 times a second.
// Above, we created a constant named funFact, which contains the text "Hummingbirds' wings can beat up to 200 times a second.". We were then able to print out this value using the names we applied. This is useful if a value doesn’t change throughout a program and it also helps to convey the developer’s intention of keeping a consistent value.

// Let’s also take a second to examine the peculiar name, funFact. Constants have names without spaces: spaces aren’t allowed! This means that for us to have descriptive names (and it is important to have descriptive names so that we can read the code we’ve written) we need to use camelCase or PascalCase, capitalizing each subsequent word instead of adding spaces.

package main

import "fmt"

func main() {
  // Create the constant earthsGravity
  // here and set it to 9.80665
  const earthsGravity = 9.80665
  // Here's where we print out the gravity:
  fmt.Println(earthsGravity)
}