package main

import "fmt"

func checkWinner(board [6][7]string, player string) bool {
	// Check horizontal
	for row := 0; row < 6; row++ {
		for col := 0; col < 4; col++ {
			if board[row][col] == player && board[row][col+1] == player && board[row][col+2] == player && board[row][col+3] == player {
				return true
			}
		}
	}

	// Check vertical
	for col := 0; col < 7; col++ {
		for row := 0; row < 3; row++ {
			if board[row][col] == player && board[row+1][col] == player && board[row+2][col] == player && board[row+3][col] == player {
				return true
			}
		}
	}

	// Check diagonal (top right to bottom left)
	for row := 0; row < 3; row++ {
		for col := 3; col < 7; col++ {
			if board[row][col] == player && board[row+1][col-1] == player && board[row+2][col-2] == player && board[row+3][col-3] == player {
				return true
			}
		}
	}

	// Check diagonal (top left to bottom right)
	for row := 0; row < 3; row++ {
		for col := 0; col < 4; col++ {
			if board[row][col] == player && board[row+1][col+1] == player && board[row+2][col+2] == player && board[row+3][col+3] == player {
				return true
			}
		}
	}
	return false
}

func main() {

	//fill the board
	var board [6][7]string
	for row := 0; row < 6; row++ {
		for col := 0; col < 7; col++ {
			board[row][col] = "."

		}

	}

	//main logic in for loop (to repeat many times)
	currentPlayer := "R"

	for {
		//print the board
		fmt.Println("0 1 2 3 4 5 6")
		for row := 0; row < 6; row++ {
			for col := 0; col < 7; col++ {
				fmt.Print(board[row][col], " ")
			}
			fmt.Println()
		}

		fmt.Println("Player", currentPlayer, "'s turn")

		//ask for input
		var col int
		fmt.Println("Enter the column number (0-6) to drop your piece:")
		fmt.Scan(&col)

		if col < 0 || col > 6 {
			fmt.Println("Invalid column!")
		} else {
			placed := false
			for row := 5; row >= 0; row-- {
				if board[row][col] == "." {
					board[row][col] = currentPlayer
					placed = true
					break
				}
			}
			if !placed {
				fmt.Println("Column is full!")
			}
			if placed {
				// Check for a winner
				if checkWinner(board, currentPlayer) {
					fmt.Println("Player", currentPlayer, "wins!")
					break
				}
				// Check for a draw
				boardFull := true
				for row := 0; row < 6; row++ {
					for col := 0; col < 7; col++ {
						if board[row][col] == "." {
							boardFull = false
							break // Exit the inner loop if boardfull is false
						}
					}
					if !boardFull {
						break // Exit the outer loop if boardfull is false
					}
				}

				if boardFull {
					fmt.Println("It's a draw!")
					break
				}

				// Switch players
				if currentPlayer == "R" {
					currentPlayer = "Y"
				} else {
					currentPlayer = "R"
				}

			}
		}

	}
	// Print the final board
	fmt.Println("\nFinal Board:")
	fmt.Println("0 1 2 3 4 5 6")
	for row := 0; row < 6; row++ {
		for col := 0; col < 7; col++ {
			fmt.Print(board[row][col], " ")
		}
		fmt.Println()
	}
}
