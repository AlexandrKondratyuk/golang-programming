package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.Open("Environmental_Data_Deep_Moor_2015.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	rdr := csv.NewReader(f)
	rdr.Comma = '\t'
	rdr.TrimLeadingSpace = true
	rows, err := rdr.ReadAll()
	if err != nil {
		panic(err)
	}
}

func mean(rows [][]string, idx int) float64 {
	var total float64
	for i, row := range rows {
		if i != 0 {
			val, _ := strconv.ParseFloat(row[idx], 64)
			total += val
		}
	}
	return total / float64(len(rows)-1)
}

func median(rows [][]string, idx int) float64 {
	var sorted []float64
	for i, row := range rows {
		if i != 0 {
			val, _ := strconv.ParseFloat(row[idx], 64)
			sorted = append(sorted, val)
		}
	}

	sort.Float64s(sorted)

	if len(rows)%2 == 0 {
		middle := len(sorted) / 2
		lower := sorted[middle-1]
		higher := sorted[middle]
		return (lower + higher) / 2
	}

	middle := len(sorted) / 2
	return sorted[middle]
}
