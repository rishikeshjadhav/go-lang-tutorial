package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	fmt.Println("Go Routines")
	say("Hey")
	go say("There")

	time.Sleep(time.Second)
}
