package main

import (
	"fmt"
	"strconv" 
)

func printBoard(board [9][9]int) {
	fmt.Println("+-------+-------+-------+")
	for row := 0; row < 9; row++ {
		fmt.Print("| ")
		for col := 0; col < 9; col++ {
			if col == 3 || col == 6 {
				fmt.Print("| ")
			}
			fmt.Printf("%d ", board[row][col])
			if col == 8 {
				fmt.Print("|")
			}
		}
		if row == 2 || row == 5 || row == 8 {
			fmt.Println("\n+-------+-------+-------+")
		} else {
			fmt.Println()
		}
	}
}

func isValidCandidate(board *[9][9]int, candidate int, row int, col int) bool {
	for j := 0; j < 9; j++ {
		if board[row][j] == candidate { return false }
	}
	for i := 0; i < 9; i++ {
		if board[i][col] == candidate { return false }
	}
	quadRow := int(row/3) * 3
	quadCol := int(col/3) * 3
	for i := quadRow; i < quadRow+3; i++ {
		for j := quadCol; j < quadCol+3; j++ {
			if board[i][j] == candidate { return false }
		}
	}
	return true
}

func backtrack(board *[9][9]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 { 
				for candidate := 1; candidate <= 9; candidate++ {
					if isValidCandidate(board, candidate, i, j) {
						board[i][j] = candidate
						if backtrack(board) {
							return true
						}
						board[i][j] = 0 // reset point for each false return on backtrack
					}
				}	
				return false // there was no valid candidate
			}
		}
	}
	return true // there are no more emtpy cells
}

func parseLayout(input string) [9][9]int {
	board := [9][9]int{}
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			board[row][col], _ = strconv.Atoi(string(input[9*row+col]))
		}
	}
	return board
}
func main() {
	var layout string
	fmt.Scanln(&layout)
	board := parseLayout(layout)
	printBoard(board)
	backtrack(&board)
	printBoard(board)
}