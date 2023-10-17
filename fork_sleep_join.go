package main

import (
	"fmt"
	"math/rand"
	"time"
)

const maxCapacity = 5

func main() {
	ch := make(chan int)

	for i := 0; i < maxCapacity; i++ {
		go exect(ch)
	}

	for i := 0; i < maxCapacity; i++ {
		<-ch
	}
	fmt.Println("fim")

}

func exect(ch chan int) {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	ch <- 0
}
