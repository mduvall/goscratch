package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		}
	}
}
