package main

import (
	"fmt"
	"math/rand"
	"time"
)

var flavorScale float32 = 5.8
var numOfFlavors int
var favoriteSnack string

// Define hoursInDay using var and = below:
var hoursInDay = 24

// Assign a value for numOfFlavors below:

// func Seed(seed int64)

func add(x int, y int) int {
	return x + y
}
func multiply(x int, y int) int {
	return x * y
}
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	numOfFlavors = 57
	favoriteSnack = "chips"
	daysOnVacation := 3
	fmt.Println("Hello World")
	fmt.Println("My name is Treveon")
	fmt.Println("The time is", time.Now())
	fmt.Println("My favorite number is", rand.Intn(3))
	fmt.Println(add(33, 36), ",nice")
	fmt.Println(multiply(11, 9))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println("You have spent", daysOnVacation*hoursInDay, "hours on vacation.")
	fmt.Println("My favorite snack is " + favoriteSnack)
	fmt.Println(numOfFlavors)

	// Declare flavorScale below:
	fmt.Println(flavorScale)

}
