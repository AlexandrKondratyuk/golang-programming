package main

import (
	"fmt"
	"io"
	"os"
)

var path = "test.txt"
var grid = make([][]string, 9)
var GridFromFile = make([][]string, 9)

func createFile() {
	// check if file exists
	var _, err = os.Stat(path) // Stat returns a FileInfo describing the named file.

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}

	fmt.Println("File Created Successfully", path)
}

func writeFile(data *[]string) {
	// Open file using READ & WRITE permission.
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// Write some text line-by-line to file.
	for i := 0; i < 9; i++ {
		_, err = file.WriteString((*data)[i])
		if isError(err) {
			return
		}
	}

	// Save file changes.
	err = file.Sync()
	if isError(err) {
		return
	}

	fmt.Println("File Updated Successfully.")
}

func readFile() {
	// Open file for reading.
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// Read file, line by line
	var text = make([]byte, 1024)
	for {
		_, err = file.Read(text)

		// Break if finally arrived at end of file
		if err == io.EOF {
			break
		}

		// Break if error occured
		if err != nil && err != io.EOF {
			isError(err)
			break
		}

	}

	fmt.Println("Reading from file.")
	//fmt.Println(string(text))

	//write data from file into new slice 'gridFromFIle'
	for i:= 0; i < 9; i++ {
		GridFromFile[i] = make([]string, 9)
	}
	for i:= 0; i < 9*9; i++ {
		GridFromFile[i/9][i%9] = string(text[i])
	}

	fmt.Println("Reading from slice what read from file.")
	//fmt.Println(GridFromFile)

}

func deleteFile() {
	// delete file
	var err = os.Remove(path)
	if isError(err) {
		return
	}

	fmt.Println("File Deleted")
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
