package main

import (
	"fmt"
	"time"
)

func main() {

	channel1, channel2 := make(chan string), make(chan string)

	go func() {
		time.Sleep(time.Millisecond * 500)

		for {
			channel1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			channel2 <- "Canal 2"
		}
	}()

	for {
		select {
		case messageChannel1 := <-channel1:
			fmt.Println(messageChannel1)
		case messageChannel2 := <-channel2:
			fmt.Println(messageChannel2)
		}

		//messageChannel1 := <-channel1
		//fmt.Println(messageChannel1)
		//
		//messageChannel2 := <-channel2
		//fmt.Println(messageChannel2)
	}
}
