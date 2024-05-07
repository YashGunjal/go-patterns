// Generator: function that returns a channel
// channels are first class values, just like strings and integers

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Making Generator")
	c := boring("boring!")
	for i := 0; i < 5; i++ {
		fmt.Printf(" You say %q\n", <-c)
	}
	fmt.Println("Main exit")
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c // return s the channel
}

// we can even have multiple instance of boring function in main, and both will produce a output to their own channel
