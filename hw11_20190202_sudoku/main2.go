package main

import (
	"fmt"
)

func main3() {
	//grid := make([][]string, 9)
	//grid[0] = []string{".", ".", ".", "1", "4", ".", ".", "2", "."}
	//grid[1] = []string{".", ".", "6", ".", ".", ".", ".", ".", "."}
	//grid[2] = []string{".", ".", ".", ".", ".", ".", ".", ".", "."}
	//grid[3] = []string{".", ".", "1", ".", ".", ".", ".", ".", "."}
	//grid[4] = []string{".", "6", "7", ".", ".", ".", ".", ".", "9"}
	//grid[5] = []string{".", ".", ".", ".", ".", ".", "8", "1", "."}
	//grid[6] = []string{".", "3", ".", ".", ".", ".", ".", ".", "6"}
	//grid[7] = []string{".", ".", ".", ".", ".", "7", ".", ".", "."}
	//grid[8] = []string{".", ".", ".", "5", ".", ".", ".", "7", "."}


	grid := make([][]string, 9)
	grid[0] = []string{"3", ".", ".", "6", ".", ".", ".", ".", "."}
	grid[1] = []string{".", "7", "5", "8", ".", "2", "3", ".", "."}
	grid[2] = []string{"8", "6", ".", ".", ".", ".", ".", "9", "4"}
	grid[3] = []string{".", ".", "7", "3", "1", ".", "9", "8", "."}
	grid[4] = []string{".", "5", ".", ".", "4", ".", ".", "2", "."}
	grid[5] = []string{".", "3", "9", ".", "2", "8", "4", ".", "."}
	grid[6] = []string{"7", "2", ".", ".", ".", ".", ".", "6", "8"}
	grid[7] = []string{".", ".", "3", "2", ".", "5", "1", "4", "."}
	grid[8] = []string{".", ".", ".", ".", ".", "7", ".", ".", "9"}

	res := CheckTask(grid)
	fmt.Println(res)

	for key, value := range grid {
		fmt.Println(key, value)
	}

	criticalCounter := 0
	for countDots(grid) > 0 && criticalCounter < 10{
		SolveProblem(grid)

		for key, value := range grid {
			fmt.Println(key, value)
		}

		criticalCounter++
	}
}


func CheckTask3(grid [][]string) bool {

	// CHECK - is correct values or no in vertical and horizontal rows
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] != "." {
				for k := 0; k < len(grid); k++ { // additional param 'k' to check with current value[i][j]
					if grid[i][k] == grid[i][j] && j != k { // check horizontally
						return false
					}
					if grid[k][j] == grid[i][j] && i != k { // check vertically
						return false
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
									return false
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
	return true
}

func SolveProblem3(grid [][]string) [][]string {
	patternMap := map[string]string{"1": "", "2": "", "3": "", "4": "", "5": "", "6": "", "7": "", "8": "", "9": ""}
	currentMap := make(map[string]string)
	resultSl := append(grid[:])
	success := 0

	fmt.Println("Count before: ", countDots(resultSl))

	for     i := range resultSl {
		for j := range resultSl[i] {

			for key, val := range patternMap {
				currentMap[key] = val
			}

			for k := 0; k < 9; k++ {
				if _, ok := currentMap[resultSl[i][k]]; ok {
					delete(currentMap, resultSl[i][k])
					//fmt.Println("hor_delete", i, ":", k, resultSl[i][k],currentMap)
				}
				if _, ok := currentMap[resultSl[k][j]]; ok {
					delete(currentMap, resultSl[k][j])
					//fmt.Println("ver_delete", i, ":", k, resultSl[k][j],currentMap)
				}
			}

			for     a := i / 3 * 3; a < i/3*3+3; a++ {
				for b := j / 3 * 3; b < j/3*3+3; b++ {
					//for kk := range currentMap {
					//	if resultSl[a][b] == kk {
							if _, ok := currentMap[resultSl[a][b]]; ok {
								delete(currentMap, resultSl[a][b])
								//fmt.Println("squ_delete", a, ":", b, resultSl[a][b], currentMap)
							}
						//}
					//}
				}
			}

			if len(currentMap) == 1 {
				for key := range currentMap {
					resultSl[i][j] = key
					success++
					fmt.Println("Bingoooo", i, ":", j, "success=", success)
				}
			}
		}
	}
	fmt.Println("Count after: ", countDots(resultSl))
	return resultSl
}

func countDots3(grid [][]string) int {
	counter := 0

	for _, v1 := range grid {
		for _, v2 := range v1 {
			if v2 == "." {
				counter++
			}
		}
	}
	return counter
}






func SolveProblem(grid [][]string) [][]string {
	patternMap := map[string]string{"1": "", "2": "", "3": "", "4": "", "5": "", "6": "", "7": "", "8": "", "9": ""}
	currentMap := make(map[string]string)
	resultSl := append(grid[:])
	success := 0

	fmt.Println("Count before: ", countDots(resultSl))

	for     i := range resultSl {
		for j := range resultSl[i] {

			for key, val := range patternMap {
				currentMap[key] = val
			}

			for k := 0; k < 9; k++ {
				if _, ok := currentMap[resultSl[i][k]]; ok {
					delete(currentMap, resultSl[i][k])
					//fmt.Println("hor_delete", i, ":", k, resultSl[i][k],currentMap)
				}
				if _, ok := currentMap[resultSl[k][j]]; ok {
					delete(currentMap, resultSl[k][j])
					//fmt.Println("ver_delete", i, ":", k, resultSl[k][j],currentMap)
				}
			}

			for     a := i / 3 * 3; a < i/3*3+3; a++ {
				for b := j / 3 * 3; b < j/3*3+3; b++ {
					//for kk := range currentMap {
					//	if resultSl[a][b] == kk {
					if _, ok := currentMap[resultSl[a][b]]; ok {
						delete(currentMap, resultSl[a][b])
						//fmt.Println("squ_delete", a, ":", b, resultSl[a][b], currentMap)
					}
					//}
					//}
				}
			}

			if len(currentMap) == 1 {
				for key := range currentMap {
					resultSl[i][j] = key
					success++
					fmt.Println("Bingoooo", i, ":", j, "success=", success)
				}
			}
		}
	}
	fmt.Println("Count after: ", countDots(resultSl))
	return resultSl
}

func countDots(grid [][]string) int {
	counter := 0

	for _, v1 := range grid {
		for _, v2 := range v1 {
			if v2 == "." {
				counter++
			}
		}
	}
	return counter
}
