package main

import (
	"fmt"
	"time"
)

func doWork(done <-chan bool) {
	for {
		select {
		case <-done:
			fmt.Println("go work exit")
			return
		default:
			fmt.Println(" Doing Work")
		}
	}

}

func main() {
	fmt.Println("for-select and Done channel pattern")
	done := make(chan bool)

	go doWork(done)

	time.Sleep(time.Second * 3)

	done <- true
	fmt.Println("3 sec")
	time.Sleep(time.Second * 3)
	fmt.Println("3 more end")
	close(done)

}
