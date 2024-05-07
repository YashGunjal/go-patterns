Here we will discuss two patterns, For-select and Done Channel.

When we are running a long running go routine, or possibilly routines with infinite loop.
For safetely or to have more control over the routine we can pass a channel variable and continue and calculation in the routine if the channel is still active else we exist the routine.

Here we  are passing a bool chan called done in doWork function. As we close the channel after 3 seconds in the Main function, the doWork function also exits.

Inside DOWork:
In doWork function we are selecting from done channel or default.
If something is published to Done channel or channle is closed the doWork function will exist, else it will continue to perform the default work.
Here we are using a combination of for and select to achieve this and this is know as for-select pattern.


```go
package main

import (
	"fmt"
	"time"
)

func doWork(done <-chan bool) {
	for {
		select {
		case <-done:
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

	close(done)
}



```