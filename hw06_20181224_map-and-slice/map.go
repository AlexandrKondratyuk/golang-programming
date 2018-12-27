package main

import "errors"

// CacheMap - create custom type cacheMap
type CacheMap map[string]string

//CreateMap - creator of cacheMap entity
func CreateMap() CacheMap {
	return make(CacheMap)
}

// Add - add element to cacheMap
func (CMap CacheMap) Add(key string, value string) (err error) { // add element to cacheMap
	if CMap.Check(key) {
		err = errors.New("Error - this key exist. Try another to add")
		return
	}
	CMap[key] = value
	return
}

// Check -  check element if exists
func (CMap CacheMap) Check(key string) (ok bool) {
	_, ok = CMap[key]
	return
}

// GetElemByKey - check element if exists
func (CMap CacheMap) GetElemByKey(key string) (val string, err error) {
	if CMap.Check(key) {
		val = CMap[key]
		return
	}
	err = errors.New("Error - the element by this key doesn't exist. Try another key")
	return
}

// Update - check element if exists
func (CMap CacheMap) Update(key string, value string) bool {
	if CMap.Check(key) {
		CMap[key] = value
		return true
	}
	return false
}

// Delete - check element if exists
func (CMap CacheMap) Delete(key string) (ok bool, err error) {
	if CMap.Check(key) {
		delete(CMap, key)
		ok = true
		return
	}
	err = errors.New("Error - element for delitting doesn't exist! Try another")
	return
}
