package main

import (
	"log"
	"time"
)

const (
	Max = 5
)

func process(c chan int) {
	for i := range c {
		log.Println("start", i)
		time.Sleep(time.Duration(i) * time.Second)
		log.Println("finish", i)
	}
}

func numbers(q chan bool) chan int {
	c := make(chan int)
	i := 0

	go func() {
		for {
			i += 1
			c <- i

			if i >= (Max * 5) + Max {
				q <- true
			}
		}
	}()

	return c
}

func main() {
	quit := make(chan bool)

	numberChan := numbers(quit)

	for i := 0; i < Max; i++ {
		go process(numberChan)
	}

	<-quit
}
