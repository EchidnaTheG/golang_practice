package main

/**
x := 42
// "x" is the name of a location in memory. That location is storing the integer value of 42

A pointer is a variable that stores the memory address of another variable. This means that a pointer "points to" the location of where the data is stored, not the actual data itself.

The * syntax defines a pointer:

var p *int

The & operator generates a pointer to its operand.

myString := "hello"
myStringPtr := &myString

It's possible to define an empty pointer. For example, an empty pointer to an integer:

var p *int

fmt.Printf("value of p: %v\n", p)
// value of p: <nil>

Its zero value is nil, which means it doesn't point to any memory address. Empty pointers are also called "nil pointers".

Instead of starting with a nil pointer, it's common to use the & operator to get a pointer to its operand:

myString := "hello"      // myString is just a string
myStringPtr := &myString // myStringPtr is a pointer to myString's address

fmt.Printf("value of myStringPtr: %v\n", myStringPtr)
// value of myStringPtr: 0x140c050

Dereference

The * operator dereferences a pointer to get the original value.

*myStringPtr = "world"                              // set myString through the pointer
fmt.Printf("value of myString: %s\n", *myStringPtr) // read myString through the pointer
// value of myString: world

Unlike C, Go has no pointer arithmetic.


package main

import (
	"strings"
)

func removeProfanity(message *string) {
	messageVal := *message
	messageVal = strings.ReplaceAll(messageVal, "fubb", "****")
	messageVal = strings.ReplaceAll(messageVal, "shiz", "****")
	messageVal = strings.ReplaceAll(messageVal, "witch", "*****")
	*message = messageVal
}

Pass by Reference

Functions in Go generally pass variables by value, meaning that functions receive a copy of most non-composite types:

func increment(x int) {
    x++
    fmt.Println(x)
    // 6
}


func main() {
    x := 5
    increment(x)
    fmt.Println(x)
    // 5
}

The main function still prints 5 because the increment function received a copy of x.

One of the most common use cases for pointers in Go is to pass variables by reference, meaning that the function receives the address of the original variable, not a copy of the value. This allows the function to update the original variable's value.

func increment(x *int) {
    *x++
    fmt.Println(*x)
    // 6
}

func main() {
    x := 5
    increment(&x)
    fmt.Println(x)
    // 6
}

Fields of Pointers

When your function receives a pointer to a struct, you might try to access a field like this and encounter an error:

msgTotal := *analytics.MessagesTotal

Instead, access it – like you'd normally do — using a selector expression.

msgTotal := analytics.MessagesTotal
package main

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}


func getMessageText(AnalyticsStructPointer *Analytics, message Message){
	if message.Success{
		AnalyticsStructPointer.MessagesTotal += 1 
		AnalyticsStructPointer.MessagesSucceeded += 1
			
	}else{
		AnalyticsStructPointer.MessagesFailed += 1
		AnalyticsStructPointer.MessagesTotal += 1 
	}
	
}
// 

This approach is the recommended, simplest way to access struct fields in Go, and is shorthand for:

(*analytics).MessagesTotal

Nil Pointers

Pointers can be very dangerous.

If a pointer points to nothing (the zero value of the pointer type) then dereferencing it will cause a runtime error (a panic) that crashes the program. Generally speaking, whenever you're dealing with pointers you should check if it's nil before trying to dereference it.

package main

import (
	"strings"
)

Let's make our profanity checker safe. Update the removeProfanity function. If message is nil, return early to avoid a panic. After all, there are no bad words to remove.

func removeProfanity(message *string) {
	if message == nil{
		return
	}
	messageVal := *message
	messageVal = strings.ReplaceAll(messageVal, "fubb", "****")
	messageVal = strings.ReplaceAll(messageVal, "shiz", "****")
	messageVal = strings.ReplaceAll(messageVal, "witch", "*****")
	*message = messageVal
}

Pointer Receivers

A receiver type on a method can be a pointer.

Methods with pointer receivers can modify the value to which the receiver points. Since methods often need to modify their receiver, pointer receivers are more common than value receivers. However, methods with pointer receivers don't require that a pointer is used to call the method. The pointer will automatically be derived from the value.
Pointer Receiver

type car struct {
	color string
}

func (c *car) setColor(color string) {
	c.color = color
}

func main() {
	c := car{
		color: "white",
	}
	c.setColor("blue")
	fmt.Println(c.color)
	// prints "blue"
}

Non-Pointer Receiver

type car struct {
	color string
}

func (c car) setColor(color string) {
	c.color = color
}

func main() {
	c := car{
		color: "white",
	}
	c.setColor("blue")
	fmt.Println(c.color)
	// prints "white"
}

Pointer Receiver Code

Methods with pointer receivers don't require that a pointer is used to call the method. The pointer will automatically be derived from the value.

type circle struct {
	x int
	y int
    radius int
}

func (c *circle) grow() {
    c.radius *= 2
}

func main() {
    c := circle{
        x: 1,
        y: 2,
        radius: 4,
    }

    // notice c is not a pointer in the calling function
    // but the method still gains access to a pointer to c
    c.grow()
    fmt.Println(c.radius)
    // prints 8
}

Pointer Performance

Occasionally, new Go developers hear "pointers don't pass copies" and take that to a logical extreme, concluding:

    Pointers are always faster because copying is slow. I'll always use pointers!

No. Bad. Stop.

Here are my rules of thumb:

    First, worry about writing clear, correct, maintainable code.
    If you have a performance problem, fix it.

Before even thinking about using pointers to optimize your code, use pointers when you need a shared reference to a value; otherwise, just use values.

If you do have a performance problem, consider:

    Stack vs. Heap
    Copying

Interestingly, local non-pointer variables are generally faster to pass around than pointers because they're stored on the stack, which is faster to access than the heap. Even though copying is involved, the stack is so fast that it's no big deal.

Once the value becomes large enough that copying is the greater problem, it can be worth using a pointer to avoid copying. That value will probably go to the heap, so the gain from avoiding copying needs to be greater than the loss from moving to the heap.

One of the reasons Go programs tend to use less memory than Java and C# programs is that Go tends to allocate more on the stack.
Where are the values to which pointers point usually stored in Go? - HEAP
Which is typically more performant to pass to a function? (Especially for small amounts of data) = VALUE




**/
