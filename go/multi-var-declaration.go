// So far we’ve been declaring variables one by one, each on their own separate line. But Go actually allows us to declare multiple variables on a single line, in fact, there’s a few different syntaxes!

// Let’s start with declaring without assigning a value:

// var part1, part2 string
// part1 = "To be..."
// part2 = "Not to be..."
// Above, we declared both part1 and part2 on the same line both with the same type. If we’re using this syntax, both variables must be the same type.

// If we already know what values we want to assign our variables we can use := like so:

// quote, fact := "Bears, Beets, Battlestar Galactica", true
// fmt.Println(quote) // Prints: Bears, Beets, Battlestar Galactica
// fmt.Println(fact) // Prints: true
// In the example above, we declare both quote and fact in the same line with one operator (:=). These variables are then assigned their respective values based on the ordering of variables and value. Since quote is the first variable, and the string "Bears, Beets, Battlestar Galactica" is the first value, quote has a value of "Bears, Beets, Battlestar Galactica". Similarly, fact then is assigned the value true.

package main

import "fmt"

func main () {
  // Define magicNum and powerLevel below:
  var magicNum, powerLevel int32
  magicNum = 2048
  powerLevel = 9001
  fmt.Println("magicNum is:", magicNum, "powerLevel is:", powerLevel)
  
  
  // Define amount and unit below:
  amount, unit := 10, "doll hairs"
  
  fmt.Println(amount, unit, ", that's expensive...")
}