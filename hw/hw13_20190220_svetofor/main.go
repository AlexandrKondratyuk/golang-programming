package main

import (
	"golang.org/x/exp/errors/fmt"
	"math/rand"
	"time"
)

type chanStruct struct {
	ch      chan int
	name    string
	counter int
}

var (
	chanMap map[string]*chanStruct
)

//var chanMap = make(map[string]chan int, 12)
//var counter = make(map[string]int, 12)

func init() {
	chanMap = make(map[string]*chanStruct, 12)

	// Creating 12 chans←, →, ↑, ↓
	chanMap["↓↓car↓↓"] = &chanStruct{make(chan int, 100), "↓↓car↓↓", 0}
	chanMap["↓↓car↓↓"] = &chanStruct{make(chan int, 100), "↓↓car↓↓", 0}
	chanMap["↑↑car↑↑"] = &chanStruct{make(chan int, 100), "↑↑car↑↑", 0}
	chanMap["→→car→→"] = &chanStruct{make(chan int, 100), "→→car→→", 0}
	chanMap["←←car←←"] = &chanStruct{make(chan int, 100), "←←car←←", 0}
	chanMap["→→pedTop→→"] = &chanStruct{make(chan int, 100), "→→pedTop→→", 0}
	chanMap["←←pedTop←←"] = &chanStruct{make(chan int, 100), "←←pedTop←←", 0}
	chanMap["→→pedBot→→"] = &chanStruct{make(chan int, 100), "→→pedBot→→", 0}
	chanMap["←←pedBot←←"] = &chanStruct{make(chan int, 100), "←←pedBot←←", 0}
	chanMap["↓↓pedLeft↓↓"] = &chanStruct{make(chan int, 100), "↓↓pedLeft↓↓", 0}
	chanMap["↑↑pedLeft↑↑"] = &chanStruct{make(chan int, 100), "↑↑pedLeft↑↑", 0}
	chanMap["↓↓pedRight↓↓"] = &chanStruct{make(chan int, 100), "↓↓pedRight↓↓", 0}
	chanMap["↑↑pedRight↑↑"] = &chanStruct{make(chan int, 100), "↑↑pedRight↑↑", 0}

	go generateRoutineSender(&(chanMap["↓↓car↓↓"].ch))
	go generateRoutineSender(&(chanMap["↑↑car↑↑"].ch))
	go generateRoutineSender(&(chanMap["→→car→→"].ch))
	go generateRoutineSender(&(chanMap["←←car←←"].ch))
	go generateRoutineSender(&(chanMap["→→pedTop→→"].ch))
	go generateRoutineSender(&(chanMap["←←pedTop←←"].ch))
	go generateRoutineSender(&(chanMap["→→pedBot→→"].ch))
	go generateRoutineSender(&(chanMap["←←pedBot←←"].ch))
	go generateRoutineSender(&(chanMap["↓↓pedLeft↓↓"].ch))
	go generateRoutineSender(&(chanMap["↑↑pedLeft↑↑"].ch))
	go generateRoutineSender(&(chanMap["↓↓pedRight↓↓"].ch))
	go generateRoutineSender(&(chanMap["↑↑pedRight↑↑"].ch))

	go generateRoutineReciever(&(chanMap["↓↓car↓↓"].ch))
	go generateRoutineReciever(&(chanMap["↑↑car↑↑"].ch))
	go generateRoutineReciever(&(chanMap["→→car→→"].ch))
	go generateRoutineReciever(&(chanMap["←←car←←"].ch))
	go generateRoutineReciever(&(chanMap["→→pedTop→→"].ch))
	go generateRoutineReciever(&(chanMap["←←pedTop←←"].ch))
	go generateRoutineReciever(&(chanMap["→→pedBot→→"].ch))
	go generateRoutineReciever(&(chanMap["←←pedBot←←"].ch))
	go generateRoutineReciever(&(chanMap["↓↓pedLeft↓↓"].ch))
	go generateRoutineReciever(&(chanMap["↑↑pedLeft↑↑"].ch))
	go generateRoutineReciever(&(chanMap["↓↓pedRight↓↓"].ch))
	go generateRoutineReciever(&(chanMap["↑↑pedRight↑↑"].ch))
}

func main() {
	// infinite loop for simulation work of traffic lights
	i := 0
	flag := false
	tempStruct := chanStruct{}
	tempCounter := 0

	for range time.Tick(time.Duration(time.Second * 3)) {
		switch {
		case i == 0:
			setResult(*chanMap["→→car→→"])
			tempStruct = *chanMap["→→car→→"]
			tempCounter = -(tempStruct.counter - <-chanMap["→→car→→"].ch)
			tempStruct.counter = tempCounter
			tempStruct = *chanMap["←←car←←"]
			tempCounter = -(tempStruct.counter - <-chanMap["←←car←←"].ch)
			tempStruct.counter = tempCounter
			tempStruct = *chanMap["→→pedTop→→"]
			tempCounter = -(tempStruct.counter - <-chanMap["→→pedTop→→"].ch)
			tempStruct.counter = tempCounter
			tempStruct = *chanMap["←←pedTop←←"]
			tempCounter = -(tempStruct.counter - <-chanMap["←←pedTop←←"].ch)
			tempStruct.counter = tempCounter
			tempStruct = *chanMap["→→pedBot→→"]
			tempCounter = -(tempStruct.counter - <-chanMap["→→pedBot→→"].ch)
			tempStruct.counter = tempCounter
			tempStruct = *chanMap["←←pedBot←←"]
			tempCounter = -(tempStruct.counter - <-chanMap["←←pedBot←←"].ch)
			tempStruct.counter = tempCounter
			/*fmt.Println("MOVE")
			printResult(chanMap["→→car→→"])
			printResult(chanMap["←←car←←"])
			printResult(chanMap["→→pedTop→→"])
			printResult(chanMap["←←pedTop←←"])
			printResult(chanMap["→→pedBot→→"])
			printResult(chanMap["←←pedBot←←"])*/

			flag = false
			i++
		case i == 1 && (flag == false || flag == true):
			if flag {
				i--
			} else {
				i++
			}
			/*fmt.Println("STOP")
			printResult(chanMap["↓↓car↓↓"])
			printResult(chanMap["↑↑car↑↑"])
			printResult(chanMap["→→car→→"])
			printResult(chanMap["←←car←←"])
			printResult(chanMap["→→pedTop→→"])
			printResult(chanMap["←←pedTop←←"])
			printResult(chanMap["→→pedBot→→"])
			printResult(chanMap["←←pedBot←←"])
			printResult(chanMap["↓↓pedLeft↓↓"])
			printResult(chanMap["↑↑pedLeft↑↑"])
			printResult(chanMap["↓↓pedRight↓↓"])
			printResult(chanMap["↑↑pedRight↑↑"])*/
		case i == 2:
			//<-chanMap["↓↓car↓↓"]
			//<-chanMap["↑↑car↑↑"]
			//<-chanMap["↓↓pedLeft↓↓"]
			//<-chanMap["↑↑pedLeft↑↑"]
			//<-chanMap["↓↓pedRight↓↓"]
			//<-chanMap["↑↑pedRight↑↑"]

			/*fmt.Println("MOVE")
			printResult(chanMap["↓↓car↓↓"])
			printResult(chanMap["↑↑car↑↑"])
			printResult(chanMap["↓↓pedLeft↓↓"])
			printResult(chanMap["↑↑pedLeft↑↑"])
			printResult(chanMap["↓↓pedRight↓↓"])
			printResult(chanMap["↑↑pedRight↑↑"])*/

			flag = true
			i--
		}
	}

}

func generateRoutineSender(ch *chan int) {
	i := 0
	for range time.Tick(time.Duration(rand.Intn(60)) * time.Millisecond) {
		i++
		*ch <- i
	}
}

func generateRoutineReciever(ch *chan int) {
	for range time.Tick(time.Duration(20) * time.Millisecond) {
		<-*ch
	}
}

func setResult(tempStruct *chanStruct) {
	tempStruct = chanMap["→→car→→"]
	tempCounter := -(tempStruct.counter - <-chanMap["→→car→→"].ch)
	tempStruct.counter = tempCounter
}


func printResult(obj chanStruct) {
	//fmt.Println("Move:")
	fmt.Printf("%s = %d \n", obj.name, obj.counter)
}
