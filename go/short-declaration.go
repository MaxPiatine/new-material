// We can also include a short variable declaration before we provide a condition in either if or switch statements. Here’s the syntax:

// x := 8
// y := 9
// if product := x * y; product > 60 {
//   fmt.Println(product, "  is greater than 60")
// }
// In our if statement, we first declare product. Notice that product is separated from the condition using a semi-colon ;. We can also have a short variable declaration inside a switch statement:

// switch season := "summer" ; season {
// case "summer":
//   fmt.Println("Go out and enjoy the sun!")
// }
// One thing to keep in mind when using the short variable declaration in if or switch statements is that the declared variable is scoped to the statement blocks. In programming, scope refers to where variables can be accessed. Having the variable scoped to the if… else if… else statement or switch statement means that variable is only accessible within the blocks of those statements and not anywhere else.

// x := 8
// y := 9
// if product := x * y; product > 60 {
//   fmt.Println(product, "  is greater than 60")
// }
// fmt.Println(product)
// The code above will throw the error:

// ./main.go:11:13: undefined: product
// Even though we defined product in our code snippet, we can only access product inside of the if block. Therefore, when we try to access it outside of the block, the compiler throws an error. We say that product is out-of-scope outside the if statement.

// Let’s use this handy shortcut in our code.

package main

import "fmt"

func main() {

	if success := true; success {
		fmt.Println("We're rich!")
	} else {
		fmt.Println("Where did we go wrong?")
	}

	amountStolen := 50000

	switch numOfThieves := 5; numOfThieves {
	case 1:
		fmt.Println("I'll take all $", amountStolen)
	case 2:
		fmt.Println("Everyone gets $", amountStolen/2)
	case 3:
		fmt.Println("Everyone gets $", amountStolen/3)
	case 4:
		fmt.Println("Everyone gets $", amountStolen/4)
	case 5:
		fmt.Println("Everyone gets $", amountStolen/5)
	default:
		fmt.Println("There's not enough to go around...")
	}
}
