// There is one more common way to define an int in Go. Computers actually have a default length for the data in their Read-Only Memory (ROM). Some newer comps may have more processing power and can store/handle bigger chunks of data. These computers might be using a 64-bit architecture, but other computers still run on 32-bit architecture and work just fine. By providing the type int or uint, Go will check to see if the computer architecture is running on is 32-bit or 64-bit. Then it will either provide a 32-bit int (or uint) or a 64-bit one depending on the computer itself.

// It’s recommended to use int unless there’s a reason to specify the size of the int (like knowing that value will be larger than the default type, or needing to optimize the amount of space used).

// var timesWeWereFooled int
// var foolishGamesPlayed uint
// Above, we declared two variables, timesWeWereFooled an integer of either 32 or 64 bits. foolishGamesPlayed, an unsigned integer of either 32 or 64 bits.

// consolationPrizes := 2
// When a variable is declared and assigned a value using the := operator, it will be the same type as if it were declared as an int. In the example above, consolationPrize has the type int.

package main

import "fmt"

func main() {
  // Define cupsOfCoffeeConsumed here
  var cupsOfCoffeeConsumed int
  // Give a value to cupsOfCoffeeConsumed
  cupsOfCoffeeConsumed = 1
  // Print out cupsOfCoffeeConsumed
  fmt.Println(cupsOfCoffeeConsumed)
}