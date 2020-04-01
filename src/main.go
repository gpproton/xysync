package main

import (
	"fmt"
	"time"
)

func main() {
	go echoTest()

	time.Sleep(100 *  time.Microsecond)
}

func echoTest() {
	fmt.Println("Hi Hi..")
}