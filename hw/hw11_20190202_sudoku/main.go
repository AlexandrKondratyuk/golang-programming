package main

import (
	"fmt"
)

func main() {

	// PART_1 ---> TEST finding solution from program
	fmt.Println("START --> finding solution from slice")
	gridFromSlice := make([][]string, 9)
	gridFromSlice[0] = []string{".", ".", ".", "1", "4", ".", ".", "2", "."}
	gridFromSlice[1] = []string{".", ".", "6", ".", ".", ".", ".", ".", "."}
	gridFromSlice[2] = []string{".", ".", ".", ".", ".", ".", ".", ".", "."}
	gridFromSlice[3] = []string{".", ".", "1", ".", ".", ".", ".", ".", "."}
	gridFromSlice[4] = []string{".", "6", "7", ".", ".", ".", ".", ".", "9"}
	gridFromSlice[5] = []string{".", ".", ".", ".", ".", ".", "8", "1", "."}
	gridFromSlice[6] = []string{".", "3", ".", ".", ".", ".", ".", ".", "6"}
	gridFromSlice[7] = []string{".", ".", ".", ".", ".", "7", ".", ".", "."}
	gridFromSlice[8] = []string{".", ".", ".", "5", ".", ".", ".", "7", "."}

	fmt.Println("sudoku before")
	View(gridFromSlice)

	FindSolution(&gridFromSlice, 0)
	fmt.Println(" \n sudoku after")
	View(gridFromSlice)



	// PART_2 ---> TEST finding solution from file
	fmt.Println("START --> finding solution from FILE")
	gridFromFile := make([]string, 9)
	gridFromFile = []string{
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
	writeFile(&gridFromFile)  // write file from '[]string'
	readFile()   // read data from file and save them to slice GridFromFile

	fmt.Println("sudoku before")
	View(GridFromFile)

	FindSolution(&GridFromFile, 0)
	fmt.Println(" \n sudoku after")
	View(GridFromFile)


	// PART_3 ---> TEST finding solution from console
	fmt.Println("START --> finding solution from CONSOLE")
	gridFromConsole := getSolutionFromConsole()

	fmt.Println(gridFromConsole)
	fmt.Println("sudoku before")
	View(gridFromConsole)

	FindSolution(&gridFromConsole, 0)
	fmt.Println(" \n sudoku after")
	View(gridFromConsole)


}