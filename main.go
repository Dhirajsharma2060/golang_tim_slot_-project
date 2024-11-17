package main

import (
	"fmt"
)

// TODO: checkwin function video time stamp:58:00

func checkWin(spin [][]string, multipilers map[string]uint) []uint {
	lines := []uint{}
	for _, row := range spin {
		win := true
		checkSymbol := row[0]
		for _, symbol := range row[1:] {
			if checkSymbol != symbol {
				win = false
				break
			}

		}
		if win {
			lines = append(lines, multipilers[checkSymbol])

		} else {
			lines = append(lines, 0)
		}

	}
	return lines
}

func main() {
	//map is more like the hashmap of go where we need to define what is the input  like key and value of the map will be value
	symbols := map[string]uint{
		"A": 4,
		"B": 7,
		"C": 12,
		"D": 20,
	}
	multipliers := map[string]uint{
		"A": 20,
		"B": 10,
		"C": 5,
		"D": 2,
	}
	symbolArr := GenerateSymbolArray(symbols)
	balance := uint(500)

	GetName()
	for balance > 0 {
		bet := GetBet(balance)
		if bet == 0 {
			break
		}
		balance = balance - bet
		spin := GetSpin(symbolArr, 3, 3)
		PrintSpin(spin)
		//check if we win and update the balance
		winningLines := checkWin(spin, multipliers)
		//fmt.Println(winningLines)
		for i, multi := range winningLines {
			win := multi * bet
			balance += win
			if multi > 0 {
				fmt.Printf("Won $%d,(%dx) on Line #%d\n", win, multi, i+1)
			}
		}
	}
	fmt.Printf("Now the left with balance , Rs%d\n", balance)

}
