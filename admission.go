package main

import (
	"fmt"
	"math/rand"
	"time"
)

const maxCapacity = 10

func main() {
	buffer := make(chan int, maxCapacity)
	ch := make(chan int)

	go create_req(buffer)
	for i := 0; i < maxCapacity; i++ {
		go exec_req(buffer)
	}
	<-ch
}

func exec_req(buffer chan int) {
	for valor := range buffer {
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		fmt.Println(valor)
	}
}

func create_req(buffer chan int) {
	for {
		valor := rand.Intn(100)
		buffer <- valor
	}

}
