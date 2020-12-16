package main

import "fmt"

var leftHand = 0
var rightHand = 0
var left = "L"
var right = "R"
var quit = "XXX"
var choice string

func main() {
	fmt.Println("Enter L for left, R for right, and XXX to quit")
	for choice != quit {
		fmt.Scanln(&choice)
		if choice == right {
			fmt.Println("Right hand added")
			rightHand++
		} else if choice == left {
			fmt.Println("Left handed added")
			leftHand++
		} else if choice == quit {
			var total = leftHand + rightHand
			fmt.Println("Left Handed:", leftHand)
			fmt.Println("Right Handed:", rightHand)
			fmt.Println("Total Counted:", total)
			break
		} else {
			fmt.Println("")
			fmt.Println("Invalid Input")
			fmt.Println("Enter L for left, R for right, and XXX to quit")
		}
	}
}
