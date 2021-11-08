package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("  __      _")
	fmt.Println("o'')}____//")
	fmt.Println("`_/      )")
	fmt.Println("(_(_/-(_/")
	fmt.Println("the time is now", time.Now())
	coolSneakers := 65.99
	niceNecklace := 45.50

	// Define taxCalculation here

	// Add coolSneakers to taxCalculation
	var taxCalculation float64
	taxCalculation += coolSneakers
	taxCalculation += niceNecklace
	taxCalculation *= .08875

	// Add niceNecklace to taxCalculation

	// Compute the NYC sales tax
	// 8.875% of the purchase here:

	// Uncomment this line for a receipt!
	fmt.Println("Purchase of", coolSneakers+niceNecklace, "with 8.875% sales tax", taxCalculation, "equal to", coolSneakers+niceNecklace+taxCalculation)
}

func main2() {
	// Define magicNum and powerLevel below:
	var magicNum, powerLevel int32
	magicNum = 2048
	powerLevel = 9001

	fmt.Println("magicNum is:", magicNum, "powerLevel is:", powerLevel)

	// Define amount and unit below:
	amount, unit := 10, "doll hairs"

	fmt.Println(amount, unit, ", that's expensive...")
}
