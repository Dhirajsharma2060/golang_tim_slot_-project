# Golang Slot Machine

![Slot Machine](istockphoto-817307644-612x612.jpg)

A simple slot machine game implemented in Golang. This project demonstrates basic Go programming concepts such as functions, loops, maps, and slices.

## Features

- User input for name and bet amount
- Random symbol generation for slot machine reels
- Win checking and balance updating
- Simple command-line interface

## Getting Started

### Prerequisites

- Go 1.23.2 or later

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/golang-slot-machine.git
    cd golang-slot-machine
    ```

2. Initialize the Go module:
    ```sh
    go mod init casino
    ```

### Running the Game

To run the game, use the following command:
```sh
go run main.go util.go spin.go
```
## How to Play
1. Enter your name when prompted.
2. Enter your bet amount. The bet must be less than or equal to your current balance.
3. The slot machine will spin and display the result.
4. If you win, your balance will be updated based on the winning lines.
5. Repeat steps 2-4 until you decide to quit by entering a bet of 0 or your balance reaches 0.

## File Structure
- main.go: Contains the main game loop and the checkWin function.
- util.go: Contains utility functions for user input.
- spin.go: Contains functions for generating and printing the slot machine   spin.
- go.mod: Go module file.
- README.md: Project documentation.

# License
This project is licensed under the MIT License. See the LICENSE file for details.

# Acknowledgments
- Inspired by various slot machine implementations and Go programming tutorials.