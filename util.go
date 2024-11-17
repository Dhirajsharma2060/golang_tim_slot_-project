package main

import "fmt"

// in go we have return the type of the return before the body of the function
func GetName() string {
	//:= is used to define the value of the left according to the value of the right
	name := ""
	fmt.Println("Enter the name : ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		return ""
	}
	fmt.Printf("Welcome %s,lets's start learning !!\n", name)
	return name
}

func GetBet(balance uint) uint {
	var bet uint
	//Note: in Golang there is no while so , instead of this we use "for true" is equivalent of the  while true
	for true {
		fmt.Printf("Enter Your Bet, or 0 to quit (balance = Rs%d): ", balance)
		fmt.Scan(&bet)

		if bet > balance {
			fmt.Println("bet cannot be larger then valid ")

		} else {
			break
		}

	}
	return bet

}
