package main

import (
	"fmt"
	"math/rand"
)

// in go we have return the type of the return before the body of the function
func getName() string {
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

func getBet(balance uint) uint {
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

// The [] mentioned here is a fixed size array{more of a slice} means we cannot add more then the max
func generateSymbolArray(symbols map[string]uint) []string {
	symbolArr := []string{}
	for symbol, count := range symbols {
		for i := uint(0); i < count; i++ {
			symbolArr = append(symbolArr, symbol)
		}

	}
	return symbolArr

}

func getRandomNumber(min int, max int) int {
	randomNumber := rand.Intn(max-min+1) + min
	return randomNumber
}
func getSpin(reel []string, rows int, cols int) [][]string {
	// we want to return a 2d slice or array
	//{{"A","B","A"},{"C","D""D"},{}} some like this
	result := [][]string{}
	for i := 0; i < rows; i++ {
		result = append(result, []string{})
	}
	for col := 0; col < cols; col++ {
		selected := map[int]bool{}
		for row := 0; row < rows; row++ {
			for true {
				randomIndex := getRandomNumber(0, len(reel)-1)
				_, exists := selected[randomIndex]
				if !exists {
					selected[randomIndex] = true
					result[row] = append(result[row], reel[randomIndex])
					break
				}

			}

		}

	}
	return result
}

func printSpin(spin [][]string) {
	for _, row := range spin {
		for j, symbol := range row {
			fmt.Printf(symbol)
			if j != len(row)-1 {
				fmt.Printf("|")
				//A|B|C
			}

		}
		fmt.Println()
	}
}

// TODO: checkwin function video time stamp:58:00
func main() {
	//map is more like the hashmap of go where we need to define what is the input  like key and value of the map will be value
	symbols := map[string]uint{
		"A": 4,
		"B": 7,
		"C": 12,
		"D": 20,
	}
	/*multipliers := map[string]uint{
		"A": 20,
		"B": 10,
		"C": 5,
		"D": 2,
	}*/
	symbolArr := generateSymbolArray(symbols)
	balance := uint(500)

	getName()
	for balance > 0 {
		bet := getBet(balance)
		if bet == 0 {
			break
		}
		balance = balance - bet
		spin := getSpin(symbolArr, 3, 3)
		printSpin(spin)
		//check if we win and update the balance
	}
	fmt.Printf("Now the left with balance , Rs%d\n", balance)

}
