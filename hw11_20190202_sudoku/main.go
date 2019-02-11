package main

import (
	"fmt"
)

func main() {

	// part_1 ---> TEST finding solution from program
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

	fmt.Println("sudoku before")
	View(grid)

	FindSolution(&grid, 0)
	fmt.Println(" \n sudoku after")
	View(grid)

	// part_2 ---> TEST finding solution from file
	grid3 := make([]string, 9)
	grid3 = []string{
		"3..6.....",
		".758.23..",
		"86.....94",
		"..731.98.",
		".5..4..2.",
		".39.284..",
		"72.....68",
		"..32.514.",
		".....7..9",
	}

	createFile() // create file
	writeFile(&grid3)  // write file from '[]string'
	readFile()   // read data from file and save them to slice GridFromFile

	fmt.Println("sudoku before")
	View(GridFromFile)

	FindSolution(&GridFromFile, 0)
	fmt.Println(" \n sudoku after")
	View(GridFromFile)

}