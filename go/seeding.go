// Previously, we saw how our random numbers weren’t entirely random. The reason for this behavior is due to how Go seeds or chooses a number as a starting point for generating random numbers. By default, the seed value is 1. We can provide a new seed value using the method rand.Seed() and passing in an integer.

// However, we encounter the same problem if we pass in a constant, i.e. pass in 5. Each time we run our program, we’ll always print the same set of numbers. Therefore, each time we run our program, we also need a unique number as a seed. One way to get this unique number each time we run our program is to use time. The reason why we can use time to be our unique number is because it’s always a different time when our program is run! Take for example:

// package main

// import (
//   "fmt"
//   "math/rand"
//   "time"
// )

// func main() {
//   rand.Seed(time.Now().UnixNano())
//   fmt.Println(rand.Intn(100))
// }
// In the example above, we import the time library. We then use the time library and call .Now() with the .UnixNano() method chained to it. This results in the difference in time (in nanoseconds) since Janurary 1st, 1970 UTC. (Check out the UnixNano documentation for more details). The number that we get from time.Now().UnixNano() is passed as an argument to rand.Seed() resulting in a different seed value each time we run our program. Since each run has a unique seed value, each run will also print out a random number from 0 to 99.

// Let’s see this change in our own code.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Add your code below:
	rand.Seed(time.Now().UnixNano())

	amountLeft := rand.Intn(10000)

	fmt.Println("amountLeft is: ", amountLeft)

	if amountLeft > 5000 {
		fmt.Println("What should I spend this on?")
	} else {
		fmt.Println("Where did all my money go?")
	}
}
