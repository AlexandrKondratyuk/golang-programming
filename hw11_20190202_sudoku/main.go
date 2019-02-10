package main

import (
	"fmt"
	. "github.com/AlexandrKondratyuk/library/cast"

)

func main() {
	grid := make([][]string, 9)
	grid[0] = []string{".", ".", ".", "1", "4", ".", ".", "2", "."}
	grid[1] = []string{".", ".", "6", ".", ".", ".", ".", ".", "."}
	grid[2] = []string{".", ".", ".", ".", ".", ".", ".", ".", "."}
	grid[3] = []string{".", ".", "1", ".", ".", ".", ".", ".", "."}
	grid[4] = []string{".", "6", "7", ".", ".", ".", ".", ".", "9"}
	grid[5] = []string{".", ".", ".", ".", ".", ".", "8", "1", "."}
	grid[6] = []string{".", "3", ".", ".", ".", ".", ".", ".", "6"}
	grid[7] = []string{".", ".", ".", ".", ".", "7", ".", ".", "."}
	grid[8] = []string{".", ".", ".", "5", ".", ".", ".", "7", "."}

	//grid := make([][]string, 9)
	//grid[0] = []string{"3", ".", ".", "6", ".", ".", ".", ".", "."}
	//grid[1] = []string{".", "7", "5", "8", ".", "2", "3", ".", "."}
	//grid[2] = []string{"8", "6", ".", ".", ".", ".", ".", "9", "4"}
	//grid[3] = []string{".", ".", "7", "3", "1", ".", "9", "8", "."}
	//grid[4] = []string{".", "5", ".", ".", "4", ".", ".", "2", "."}
	//grid[5] = []string{".", "3", "9", ".", "2", "8", "4", ".", "."}
	//grid[6] = []string{"7", "2", ".", ".", ".", ".", ".", "6", "8"}
	//grid[7] = []string{".", ".", "3", "2", ".", "5", "1", "4", "."}
	//grid[8] = []string{".", ".", ".", ".", ".", "7", ".", ".", "9"}

	fmt.Println("sudoku before")
	View(grid)

	FindSolution(&grid, 0)
	fmt.Println(" \n sudoku after")
	View(grid)
}

// Logic : is correct task for finding solution. If incorrect - return 'false'
func CheckTask(grid [][]string) error {

	// CHECK - is correct values or no in vertical and horizontal rows
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] != "." {
				for k := 0; k < len(grid); k++ { // additional param 'k' to check with current value[i][j]
					if grid[i][k] == grid[i][j] && j != k { // check horizontally
						return fmt.Errorf("Error: Sudoku is incorrect and hasn't correct solution. Incorret when check horizontal line")
					}
					if grid[k][j] == grid[i][j] && i != k { // check vertically
						return fmt.Errorf("Error: Sudoku is incorrect and hasn't correct solution. Incorret when check vertical line")
					}
				}
			}
		}
	}

	// CHECK - is correct values in every square 3*3
	var val string // temp variable 'val' for comparing it with current value[i][j]

	for r := 0; r < len(grid); r += 3 { // iterating through all square (9*9)
		for c := 0; c < len(grid); c += 3 {

			for i := r; i < r+3; i++ { // iterating through current square (3*3)
				for j := c; j < c+3; j++ {
					if grid[i][j] != "." {
						val = grid[i][j] // var for checking with - has initial value [i][j]
						//fmt.Println(val)
						for a := r; a < r+3; a++ {
							for b := c; b < c+3; b++ { // don't check if curent value == [i][j]
								if a == i && b == j {
									continue
								}
								if val == grid[a][b] { // false if current value == [i][j]
									return fmt.Errorf("Error: Sudoku is incorrect and hasn't correct solution. Incorret when check square 3*3")
								}
								if a == r+2 && b == c+2 { // change 'val' to next value for checking
									val = grid[a][b]
								}
								if val != grid[a][b] { // ??? test is needed
									continue
								}
							}
						}
					}
				}
			}
		}
	}
	return nil
}

// View : result to console
func View(grid [][]string) {
	for row := 0; row < 9; row++ {
		fmt.Println(" ")
		if row == 3 || row == 6 {
			fmt.Println(" ")
		}
		for col := 0; col < 9; col++ {
			if col == 3 || col == 6 {
				fmt.Print(" ")
			}
			if col == 8 {
				fmt.Print(grid[row][col])
			} else if col != 3 || col != 6 {
				fmt.Print(grid[row][col], " ")
			}
			if row == 8 && col == 8 {
				fmt.Println()
			}
		}
	}
}

// Logic : for recursive finding correct results
func FindSolution(grid *[][]string, currentPosition int) bool {
	if CheckTask(*grid) != nil {    //  if task correct return solution, else error
		fmt.Println(" \n Task incorrect")
		return false
	}

	var (
		row, col, checkedInt int
		checkedString        string
	)

	if currentPosition == 81 {      // check if all values are in solution - return True
		return true
	}

	row = currentPosition / 9
	col = currentPosition % 9

	if (*grid)[row][col] != "." {   // recursive check not empty cells
		return FindSolution(grid, currentPosition+1)
	}

	for checkedInt = 1; checkedInt <= 9; checkedInt++ {     // check if result doesn't exist, write result into current cell
		checkedString, _ = ToString(checkedInt)

		if CheckColumn(checkedString, *grid, row) && CheckRow(checkedString, *grid, col) && CheckSquare(checkedString, *grid, row, col) {

			(*grid)[row][col] = checkedString

			if FindSolution(grid, currentPosition+1) {      //  after write we are going to find next result for next position
				return true
			}
		}
	}

	(*grid)[row][col] = "."         // if result is incorrect, write into current cell '.' and return 'false'
	return false
}

// Logic : check current value in row and return 'true' if doesn't exist
func CheckColumn(checkedString string, grid [][]string, row int) bool {
	for col := 0; col < 9; col++ {
		if grid[row][col] == checkedString {
			return false
		}
	}
	return true
}

// Logic : check current value in column and return 'true' if doesn't exist
func CheckRow(checkedString string, grid [][]string, col int) bool {
	for row := 0; row < 9; row++ {
		if grid[row][col] == checkedString {
			return false
		}
	}
	return true
}

// Logic : check current value in square 3*3 and return 'true' if doesn't exist
func CheckSquare(checkedString string, grid [][]string, row int, col int) bool {
	startFromRow := row - (row % 3)
	startFromCol := col - (col % 3)

	for row = startFromRow; row < startFromRow+3; row++ {
		for col = startFromCol; col < startFromCol+3; col++ {
			if grid[row][col] == checkedString {
				return false
			}
		}
	}
	return true
}
