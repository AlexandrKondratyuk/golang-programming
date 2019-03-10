package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
)

func main() {
	//res, err := http.Get("Environmental_Data_Deep_Moor_2015.txt")
	res, err := http.Get("https://lpo.dt.navy.mil/data/DM/Environmental_Data_Deep_Moor_2015.txt")
	if err != nil {
		log.Fatal(err)
	}
	rdr := csv.NewReader(res.Body)
	rdr.Comma = '\t'
	rdr.TrimLeadingSpace = true
	defer res.Body.Close()
	rows, err := rdr.ReadAll()
	if err != nil {
		panic(err)
	}

	// Don't count the header row in len(rows)
	fmt.Println("Total Records: ", len(rows)-1)
	fmt.Println("Air Temp:\t", mean(rows, 1), median(rows, 1))
	fmt.Println("Barometric:\t", mean(rows, 2), median(rows, 2))
	fmt.Println("Wind Speed:\t", mean(rows, 7), median(rows, 7))
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
