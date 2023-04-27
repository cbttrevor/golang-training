package main

import (
	"fmt"
	"sync"
)

type Airplane struct {
	manufacturer  string
	distanceFlown uint32
}

func (a *Airplane) Fly(distance uint32) {
	a.distanceFlown += distance
}

type SupportedChannels interface {
	uint8 | string | uint16 | int16
}

func ProcessMessage[T SupportedChannels](inputChan chan T, wg *sync.WaitGroup) {
	fmt.Println(<-inputChan)
	fmt.Println(<-inputChan)
	fmt.Println(<-inputChan)
	wg.Done()
}

func DoSomeFlying(inputPlanes chan *Airplane, wg *sync.WaitGroup) {
	plane := <-inputPlanes

	plane.Fly(500)
	fmt.Println("Flew 500 miles")

	wg.Done()
}

func main() {
	mychan := make(chan uint8, 5)
	fmt.Println(mychan)

	// Adding uint8s to channel
	mychan <- 5
	mychan <- 50
	mychan <- 55

	myStringChan := make(chan string, 5)
	fmt.Println(myStringChan)
	myStringChan <- "Trevor"
	myStringChan <- "John"
	myStringChan <- "Daniel"

	var myWaitGroup sync.WaitGroup
	myWaitGroup.Add(3)

	go ProcessMessage[uint8](mychan, &myWaitGroup)
	go ProcessMessage[string](myStringChan, &myWaitGroup)

	plane01 := Airplane{manufacturer: "Boeing"}
	planeChan := make(chan *Airplane, 2)
	planeChan <- &plane01
	go DoSomeFlying(planeChan, &myWaitGroup)

	myWaitGroup.Wait()

	plane01.Fly(900)
	fmt.Printf("%+v\n", plane01)

	// time.Sleep(500 * time.Millisecond)
	//	for {
	//		select {
	//		case myChanValue := <-mychan:
	//			fmt.Println(myChanValue)
	//		case myName := <-myStringChan:
	//			fmt.Println(myName)
	//		default:
	//			fmt.Println("No messages available on any channels")
	//			// time.Sleep(1 * time.Second)
	//			goto DoneProcessing
	//		}
	//	}
	//
	// DoneProcessing:
}
