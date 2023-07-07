package main

import (
	"log"
	"time"
)

const N = 100

func printCompletionPercentage(c chan int) {
	for i := range c {
		log.Printf("%d%% loop completed ...", i)
	}
}

func main() {
	c := make(chan int)
	// b := make(chan bool)

	defer close(c)
	go printCompletionPercentage(c)

	for i := 0; i < N; i++ {
		if (i+1)%int(N/4) == 0 {
			c <- i + 1
		}
		time.Sleep(100 * time.Millisecond)
	}

}
