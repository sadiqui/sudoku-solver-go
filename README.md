### Sudoku Solver

This program is designed to solve Sudoku puzzles. It ensures that the provided puzzle is valid and has a unique solution.

#### Instructions

1. **Run the Program**: Execute the program by providing the Sudoku puzzle as command-line arguments. Ensure that each row of the puzzle is provided as a separate argument.

2. **Output**: If a valid solution exists, the program will print the solved Sudoku grid. Otherwise, it will display an error message indicating that the puzzle is invalid or unsolvable.

#### Usage

**Line by line**: dots represent empty cells.

```bash
$ go run . ".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7" | cat -e
3 9 6 2 4 5 7 8 1$
1 7 8 3 6 9 5 2 4$
5 2 4 8 1 7 3 9 6$
2 8 7 9 5 1 6 4 3$
9 3 1 4 8 6 2 7 5$
4 6 5 7 2 3 9 1 8$
7 1 2 6 3 8 4 5 9$
6 5 9 1 7 4 8 3 2$
8 4 3 5 9 2 1 6 7$
```

#### The code functionality

- `checkArgs`: Verifies if the provided Sudoku puzzle arguments are valid.
- `Print`: Prints the Sudoku grid with appropriate formatting.
- `isValidPlace`: Checks if placing a number at a specific position in the grid is valid.
- `backTracking`: Recursively solves the Sudoku puzzle using backtracking algorithm.
- `fillCell` : Fills an empty cell with a valid number and recursively solves the puzzle.
- `solveSudoku`: Creates a Grid from the given puzzle arguments.
