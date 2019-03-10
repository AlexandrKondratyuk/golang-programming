package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strconv"
)

func main() {
	//res, err := http.Get("Environmental_Data_Deep_Moor_2015.txt")
	res, err := http.Get("http://www.lynda.com/")
	if err != nil {
		log.Fatal(err)
	}
	page, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", page)
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
