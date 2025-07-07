package main
/**
Concurrency

Concurrency is the ability to perform multiple tasks at the same time. Typically, our code is executed one line at a time, one after the other. This is called sequential execution or synchronous execution.

If the computer we're running our code on has multiple cores, we can even execute multiple tasks at exactly the same time. If we're running on a single core, a single core executes code at almost the same time by switching between tasks very quickly. Either way, the code we write looks the same in Go and takes advantage of whatever resources are available.
How Does Concurrency Work in Go?

Go was designed to be concurrent, which is a trait fairly unique to Go. It excels at performing many tasks simultaneously safely using a simple syntax.

There isn't a popular programming language in existence where spawning concurrent execution is quite as elegant, at least in my opinion.

Concurrency is as simple as using the go keyword when calling a function:

go doSomething()

In the example above, doSomething() will be executed concurrently with the rest of the code in the function. The go keyword is used to spawn a new goroutine.

package main

import (
	"fmt"
	"time"
)

func sendEmail(message string) {
	go func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received: '%s'\n", message)
	}()
	fmt.Printf("Email sent: '%s'\n", message)
}

// Don't touch below this line

func test(message string) {
	sendEmail(message)
	time.Sleep(time.Millisecond * 500)
	fmt.Println("========================")
}

func main() {
	test("Hello there Kaladin!")
	test("Hi there Shallan!")
	test("Hey there Dalinar!")
}

Email sent: 'Hello there Kaladin!'
Email received: 'Hello there Kaladin!'
========================
Email sent: 'Hi there Shallan!'
Email received: 'Hi there Shallan!'
========================
Email sent: 'Hey there Dalinar!'
Email received: 'Hey there Dalinar!'
========================

Channels

Channels are a typed, thread-safe queue. Channels allow different goroutines to communicate with each other.
Create a Channel

Like maps and slices, channels must be created before use. They also use the same make keyword:

ch := make(chan int)

Send Data to a Channel

ch <- 69

The <- operator is called the channel operator. Data flows in the direction of the arrow. This operation will block until another goroutine is ready to receive the value.
Receive Data from a Channel

v := <-ch

This reads and removes a value from the channel and saves it into the variable v. This operation will block until there is a value in the channel to be read.
Blocking and Deadlocks

A deadlock is when a group of goroutines are all blocking so none of them can continue. This is a common bug that you need to watch out for in concurrent programming.


package main

import (
	"time"
)

type email struct {
	body string
	date time.Time
}

func checkEmailAge(emails [3]email) [3]bool {
	isOldChan := make(chan bool)

	go sendIsOld(isOldChan, emails)

	isOld := [3]bool{}
	isOld[0] = <-isOldChan
	isOld[1] = <-isOldChan
	isOld[2] = <-isOldChan
	return isOld
}

// don't touch below this line

func sendIsOld(isOldChan chan<- bool, emails [3]email) {
	for _, e := range emails {
		if e.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
			isOldChan <- true
			continue
		}
		isOldChan <- false
	}
}

if there wasn't a go routine:


    sendIsOld is called (synchronously): checkEmailAge calls sendIsOld. The program execution completely moves into sendIsOld.
    First isOldChan <- value: sendIsOld attempts to send a value to isOldChan.
    No Receiver: At this point, the checkEmailAge function (the caller) is paused, waiting for sendIsOld to finish. It hasn't yet reached the lines where it tries to <-isOldChan (receive from the channel).
    Blocking: Since there's no other goroutine ready to receive the value that sendIsOld is trying to send, the send operation blocks indefinitely.
    Deadlock: Because sendIsOld is blocked, it never returns control to checkEmailAge. And because checkEmailAge never gets control back, it can never reach its receive operations. Both sides are waiting for the other, resulting in a deadlock, and the program never finishes.

Think of it like two people trying to pass a ball, but only one person is present. If the person with the ball tries to pass it, they'll just stand there with their arm extended, waiting for someone to catch it, and nothing else will happen.

Channels

In the previous lesson, we saw how we can receive values from channels like this:

v := <-ch

However, sometimes we don't care what is passed through a channel. We only care when and if something is passed. In that situation, we can block and wait until something is sent on a channel using the following syntax.

<-ch

This will block until it pops a single item off the channel, then continue, discarding the item.

In cases like this, Empty structs are often used as a unary value so that the sender communicates that this is only a "signal" and not some data that is meant to be captured and used by the receiver.

Here is an example:

func downloadData() chan struct{} {
	downloadDoneCh := make(chan struct{})

	go func() {
		fmt.Println("Downloading data file...")
		// after the download is done, send a "signal" to the channel
		downloadDoneCh <- struct{}{}
	}()

	return downloadDoneCh
}

func processData(downloadDoneCh chan struct{}) {
	// any code here can run normally
	fmt.Println("Preparing to process data...")

	// block until `downloadData` sends the signal that it's done
	<-downloadDoneCh

	// any code here can assume that data download is complete
	fmt.Println("Data download ensured, starting data processing...")
}

**/