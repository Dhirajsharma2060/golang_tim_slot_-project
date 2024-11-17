package main

import (
	"fmt"
	"math/rand"
)

// The [] mentioned here is a fixed size array{more of a slice} means we cannot add more then the max
func GenerateSymbolArray(symbols map[string]uint) []string {
	symbolArr := []string{}
	for symbol, count := range symbols {
		for i := uint(0); i < count; i++ {
			symbolArr = append(symbolArr, symbol)
		}

	}
	return symbolArr

}

func GetRandomNumber(min int, max int) int {
	randomNumber := rand.Intn(max-min+1) + min
	return randomNumber
}
func GetSpin(reel []string, rows int, cols int) [][]string {
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
				randomIndex := GetRandomNumber(0, len(reel)-1)
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

func PrintSpin(spin [][]string) {
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
