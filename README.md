 # go-patterns



## Theory

#### Buffered and Unbuffered channels
A Channel with defined size is buffered channel, A Buffered and unbuffered channels in Go are used for communication between goroutines. Both serve the purpose of synchronizing goroutines, but they do so in slightly different ways, leading to different behaviors in terms of synchronous and asynchronous communication.

> Giving size 1 also creates a buffered channel with size one.

### Unbuffered Channels:

- **Synchronous Communication**: When you send a value to an unbuffered channel, the sender will block until the receiver is ready to receive the value. Similarly, when you receive a value from an unbuffered channel, the receiver will block until there is a sender ready to send a value.
- **Blocking Behavior**: Unbuffered channels block both the sender and the receiver until they are both ready to communicate. This means that both the sender and the receiver synchronize with each other, ensuring that the data exchange happens at the exact moment both are ready.
- **Guaranteed Synchronization**: Unbuffered channels ensure that the sender and receiver are synchronized, making them suitable for scenarios where precise synchronization is required.

### Buffered Channels:

- **Asynchronous Communication**: Buffered channels allow for asynchronous communication. When sending a value to a buffered channel, if there is space available in the buffer, the sender can continue immediately without blocking, even if there is no receiver ready to receive the value. Similarly, when receiving a value from a buffered channel, if there are values in the buffer, the receiver can receive them immediately without blocking, even if there is no sender ready to send.
- **Non-blocking Writes and Reads**: Buffered channels allow non-blocking writes and reads as long as there is space in the buffer (for writes) or there are values in the buffer (for reads). This makes them suitable for scenarios where immediate processing is more important than precise synchronization.
- **Potential for Deadlocks**: Since buffered channels allow asynchronous communication, there is a possibility of deadlocks if not used carefully. For example, if the buffer is filled and there are no receivers, or if the buffer is empty and there are no senders, the program may deadlock.

### Example:

```go
package main

import "fmt"

func main() {
    // Unbuffered channel
    unbuffered := make(chan int)

    // Buffered channel with capacity 2
    buffered := make(chan int, 2)

    // Sending on unbuffered channel (synchronous)
    go func() {
        unbuffered <- 1
        fmt.Println("Sent on unbuffered channel")
    }()

    // Sending on buffered channel (asynchronous)
    go func() {
        buffered <- 2
        fmt.Println("Sent on buffered channel")
    }()

    // Receiving from unbuffered channel (synchronous)
    fmt.Println("Received from unbuffered channel:", <-unbuffered)

    // Receiving from buffered channel (asynchronous)
    fmt.Println("Received from buffered channel:", <-buffered)
}
```

In this example:

- Sending and receiving on the unbuffered channel synchronizes the sender and receiver.
- Sending and receiving on the buffered channel does not synchronize the sender and receiver immediately.


