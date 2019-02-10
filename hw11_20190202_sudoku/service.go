package main

import "fmt"

func CheckTask2(grid [][]string) bool {

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

func SolveProblem2(grid [][]string) [][]string {
	patternMap := map[string]string{"1": "", "2": "", "3": "", "4": "", "5": "", "6": "", "7": "", "8": "", "9": ""}
	currentMap := make(map[string]string)
	resultSl := append(grid[:])
	success := 0

	fmt.Println("Count before: ", countDots(resultSl))

	for i := range resultSl {
		for j := range resultSl[i] {

			for key, val := range patternMap {
				currentMap[key] = val
			}

			//fmt.Println("currentMap=", currentMap)
			for k := 0; k < 9; k++ {
				//for kk := range currentMap {
					//if grid[i][k] == kk {
					if _, ok := currentMap[resultSl[i][k]]; ok {
						//if _, ok := currentMap[kk]; ok {
							delete(currentMap, resultSl[i][k])
							fmt.Println("Hor_delete", i, ":", k, currentMap)
						//}
					}
					//if grid[k][j] == kk {
					if _, ok := currentMap[resultSl[k][j]]; ok {
						//if _, ok := currentMap[kk]; ok {
							delete(currentMap, resultSl[k][j])
							fmt.Println("Ver_delete", i, ":", k, currentMap)
						//}
					}
				//}
			}
			//fmt.Println(1, len(currentMap))

			for a := i / 3 * 3; a < i/3*3+3; a++ {
				for b := j / 3 * 3; b < j/3*3+3; b++ {
					for kk := range currentMap {
						if resultSl[a][b] == kk {
							if _, ok := currentMap[kk]; ok {
								delete(currentMap, kk)
								fmt.Println("squ_delete", i, ":", kk, currentMap)
							}
						}
					}
				}
			}

			if len(currentMap) == 1 {
				for key := range currentMap {
					resultSl[i][j] = key
					success++
					fmt.Println("Bingoooo", i, ":", j, "success=", success)
				}
			}
			//fmt.Printf("[%v][%v]=%v\n", i, j, currentMap)
			//fmt.Println(2, len(currentMap))
		}
	}
	fmt.Println("Count after: ", countDots(resultSl))
	return resultSl
}

func countDots2(grid [][]string) int {
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
