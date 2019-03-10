package main

import (
	"errors"
	"fmt"
	"strconv"
)

// CacheSlice - creator for new slice with len() and cap()
type CacheSlice []string

// CreateSlice - creator for type CacheSlice
func CreateSlice(len int, cap int) (cSlice CacheSlice) {
	if len > cap {
		fmt.Println(fmt.Errorf("Error -> LAN(%d) is bigger then CAP(%d). Please chose correct arguments", len, cap))
		return
	}
	return make(CacheSlice, len, cap)
}

// Add - Func for adding element into slice
func (CSlice *CacheSlice) Add(value ...string) {
	*CSlice = append(*CSlice, value...)
}

// Check -  check element if exists
func (CSlice *CacheSlice) Check(index int) bool {
	if index >= 0 && index < len(*CSlice) {
		return true
	}
	fmt.Println(fmt.Errorf("Error -> Index(%d) is out of range. Please chose correct idex, less then %d", index, len(*CSlice)))
	return false
}

// GetElemByIndex - return element by certain index (if exists)
func (CSlice *CacheSlice) GetElemByIndex(index int) (val string, err error) {
	if index >= len(*CSlice) {
		err = errors.New("Error here -> Index(" + strconv.Itoa(index) + " is out of range. Please chose correct index, less then " + strconv.Itoa(len(*CSlice)))
		return
	}
	val = (*CSlice)[index]
	return
}

// Update - Func update element by index if exists
func (CSlice CacheSlice) Update(index int, value string) (ok bool, err error) {
	if !CSlice.Check(index) {
		err = errors.New("Error here -> Index(" + strconv.Itoa(index) + " doesn\t exist. Please chose another index for update")
		return
	}
	CSlice[index] = value
	return
}

// Delete - Func for deleting element by index if exists
func (CSlice *CacheSlice) Delete(index int) (ok bool, err error) {
	if !CSlice.Check(index) {
		err = errors.New("Error here -> Element with index(" + strconv.Itoa(index) + " doesn\t exist. Please chose another index for delete element")
		return
	}
	*CSlice = append((*CSlice)[:index], (*CSlice)[index+1:]...)
	return
}
