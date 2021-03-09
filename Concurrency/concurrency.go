package main

import (
	"fmt"
	"time"
)

func main() {
	//go count("sheep") // run in background and cause a goroutine and continue in background.
	////count("fish") // will run both infinite loops simultaneously
	//go count("fish") // run in background and cause a goroutine and continue in background.
	//// when the first go routine it finishes, causing this to not output anything.

	c := make(chan string)

	go count("bukkit", c)

	//Manual checking if chan is closed.
	//for {
	//	msg, open := <- c
	//
	//	if !open {
	//		break
	//	}
	//
	//	fmt.Println(msg)
	//}

	for msg := range c {
		fmt.Println(msg)
	}
}

func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}

	close(c)
}