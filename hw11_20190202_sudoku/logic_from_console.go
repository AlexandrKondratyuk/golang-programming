package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getSolutionFromConsole() [][]string{
	gridFromConsole := make([][]string, 9)
	str := bufio.NewReader(os.Stdin)
	fmt.Println("Type 9 rows of data")

Start:
	for index := range gridFromConsole {
		fmt.Printf("%d-st row --> ", index+1)
		//text, _ = str.ReadLine()
		text, _ := str.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		gridFromConsole[index] = strings.Split(text, "")
		if len(gridFromConsole[index]) < 9 {
			fmt.Println("Error: Sudoku is incorrect and hasn't enough numbers.")
			fmt.Println("Try again type 9 rows from 9 numbers or dots")
			goto Start
		} else if len(gridFromConsole[index]) > 9 {
			gridFromConsole[index] = gridFromConsole[index][0:9]
		}

	}
	return gridFromConsole
}
