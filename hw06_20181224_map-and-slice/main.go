package main

import "fmt"

func main() {
	myMap := CreateMap() // create and initialize map by default (type = cacheMap)
	mySlice := CreateSlice(0, 10)

	resOfAddElemToMap := myMap.Add("first", "Kiev") //add an element to map if don't exists
	mySlice.Add("one", "two", "three")

	resultCheck := myMap.Check("first") // check if is element
	resultCheckSliceByIndex := mySlice.Check(0)

	resultVelueByKey, success := myMap.GetElemByKey("first")
	resultVelueByIndex, success := mySlice.GetElemByIndex(0)

	myMap.Add("second", "Kiev")
	isUpdated := myMap.Update("second", "Panama") // update map by key/value, return tru if is updated
	isUpdated2, err := mySlice.Update(0, "six")   // update slice by index, return tru if is updated

	myMap.Add("third", "Praha")
	isDeleted, success := myMap.Delete("third") // delete element by key, return true if is deleted
	isDeleted2, success2 := mySlice.Delete(1)   // delete element from slice by index, return true if is deleted

	fmt.Printf("%T = %v \n", myMap, myMap)
	fmt.Println("Result of adding =>>> ", resOfAddElemToMap)
	fmt.Printf("Func check = %v \n", resultCheck)
	fmt.Printf("Func getElemByKey = %v. Success => %v\n", resultVelueByKey, success)
	fmt.Printf("Func update = %v \n", isUpdated)
	fmt.Printf("%T = %v \n", myMap, myMap)
	fmt.Printf("%T = %v \n", myMap, myMap)
	fmt.Printf("Func delete = %v. Success => %v \n", isDeleted, success)
	fmt.Printf("%T = %v \n", myMap, myMap)

	fmt.Printf("%T = %v \n", mySlice, mySlice)
	fmt.Println("resultCheckSliceByIndex => ", resultCheckSliceByIndex)
	fmt.Println("resultVelueByIndex => ", resultVelueByIndex, ". Success => ", success)
	fmt.Println(isUpdated2, " --- ", err)
	fmt.Printf("******DELETE*****isDeleted2 =  %v, success2 =  %v\n", isDeleted2, success2)
}
