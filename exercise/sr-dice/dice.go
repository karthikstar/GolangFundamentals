//--Summary:
//  Create a program that can perform dice rolls using various configurations
//  of number of dice, number of rolls, and number of sides on the dice.
//  The program should report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these cirsumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must use variables to configure:
//  - number of times to roll the dice
//  - number of dice used in the rolls
//  - number of sides of all the dice (6-sided, 10-sided, etc determined
//    with a variable). All dice must use the same variable for number
//    of sides, so they all have the same number of sides.
//
//--Notes:
//* Use packages from the standard library to complete the project
//* Try using different values for your variables to check the results

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func roll(sides int) int {
	return rand.Intn(sides) + 1 // starts from 0, so we need +1 so that it can start from 1, and range ends excl the no passed in so +1 makes sense
}

func main() {
	noOfDiceUsed := 2
	fmt.Println("No Of Dices Used", noOfDiceUsed)
	noOfSides := 6
	fmt.Println("No Of Sides for each dice", noOfSides)
	noOfTimesToRollDice := 2
	fmt.Println("no of times to roll dice", noOfTimesToRollDice)
	rand.Seed(time.Now().UnixNano()) // seed value will be the current itme, which we can use the unix nano function in time package

	for i := 1; i <= noOfTimesToRollDice; i++ {
		totalRoll := 0
		for j := 1; j <= noOfDiceUsed; j++ {
			roll := roll(noOfSides)
			totalRoll += roll
			fmt.Println("Roll #", i, ", Dice #", j, ":", roll)
		}

		fmt.Println("total roll is:", totalRoll)
		switch totalRoll := totalRoll; {
		case totalRoll == 2 && noOfDiceUsed == 2:
			fmt.Println("Snake eyes")
			fallthrough
		case totalRoll%2 == 0:
			fmt.Println("Even")
		case totalRoll == 7:
			fmt.Println("Lucky 7")
			fallthrough
		default: // odd no
			fmt.Println("Odd")
		}
	}

}
