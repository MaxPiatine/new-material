
// Let’s focus on the first line package main. 
// This line is called a package declaration and it tells the Go compiler to which package this ‘.go’ file belongs.
//  If this package declaration is ‘package main’, then this program will be compiled into an executable.


package main

// Then we have an import statement, import "fmt".
// The import keyword allows us to use code from other packages, 
// in this case the Println function from the fmt package.
// Note how the package name is enclosed with double quotes ".

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
}

// to run this code inside terminal, you must:
// >> go run main.go
