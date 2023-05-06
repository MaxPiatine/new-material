// We often include conditionals for when we need to account for different situations and outcomes. When it comes to coming up with a plan and executing it, there’s nothing more uncertain than a bank heist! (Watch any robbery themed movie if you need convincing).

// In this project, we’re going to use conditionals to simulate this situation along with hi-jinks and hiccups that may pop up. Who knows? Maybe, just maybe, if all goes well, we’ll have a few more dollars after.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	isHeistOn := true
	eludedGuards := rand.Intn(100)

	if eludedGuards >= 50 {
		fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
	} else {
		isHeistOn = false
		fmt.Println("Plan a better disguise next time?")
	}

	openedVault := rand.Intn(100)

	if isHeistOn && openedVault >= 70 {
		fmt.Println("Grab and GO!")
	} else if isHeistOn {
		isHeistOn = false
		fmt.Println("Get the combination next time?")
	}

	leftSafely := rand.Intn(2)

	if isHeistOn {
		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println("Plan a better escape next time?")
		case 1:
			isHeistOn = false
			fmt.Print("Turns out vault doors don't open from the inside...")
		default:
			fmt.Println("Start the getaway car!")
		}
	}

	if isHeistOn {
		amtStolen := 10000 + rand.Intn(1000000)
		fmt.Println("Succesfull Heist. Stolen: ", amtStolen)
	}

	fmt.Println("Heist is currently: ", isHeistOn)
}
