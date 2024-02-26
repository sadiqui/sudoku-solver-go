package main

import (
	"fmt"
	"os"
)

type Grid [][]byte

// Checks if the provided arguments are valid
func checkArgs(args []string) bool {
	if len(args) != 9 {
		return false
	}
	for _, row := range args {
		if len(row) != 9 {
			return false
		}
		for _, cell := range row {
			if cell != '.' && (cell < '1' || cell > '9') {
				return false
			}
		}
	}
	return true
}

// Prints the Sudoku grid
func (grid Grid) Print() {
	for _, row := range grid {
		for i, cell := range row {
			fmt.Printf("%c", cell)
			if i < len(row)-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

// Checks if placing a number at a specific position is valid
func (grid Grid) isValidPlace(x, y int, num byte) bool {
	for i := 0; i < 9; i++ {
		if grid[y][i] == num || grid[i][x] == num || grid[(y/3)*3+i/3][(x/3)*3+i%3] == num {
			return false
		}
	}
	return true
}

// Recursively solves the Sudoku using backtracking algorithm
func (grid Grid) backTracking() bool {
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if grid[y][x] == '.' {
				if grid.fillCell(x, y) {
					return true
				}
				return false
			}
		}
	}
	return true
}

// Fills an empty cell with a valid number in recursive solving
func (grid Grid) fillCell(x, y int) bool {
	for num := byte('1'); num <= '9'; num++ {
		if grid.isValidPlace(x, y, num) {
			grid[y][x] = num
			if grid.backTracking() {
				return true
			}
			grid[y][x] = '.' // Backtrack if no solution is found
		}
	}
	return false
}

// Creates a Grid from the given puzzle arguments
func solveSudoku(args []string) Grid {
	var grid Grid
	for _, row := range args {
		grid = append(grid, []byte(row))
	}
	return grid
}

func main() {
	args := os.Args
	if !checkArgs(args[1:]) {
		fmt.Println("Error")
		return
	}
	grid := solveSudoku(args[1:])
	if grid.backTracking() {
		grid.Print()
	} else {
		fmt.Println("No solution exists.")
	}
}
