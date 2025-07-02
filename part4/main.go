package main
import (
	"fmt"
)

/**
ERRORS, PANICS and LOOPS 

Go programs express errors with error values. An Error is any type that implements the simple built-in error interface.

type error interface {
    Error() string
}

When something can go wrong in a function, that function should return an error as its last return value. Any code that calls a function that can return an error should handle errors by testing whether the error is nil.
Let's look at how the strconv.Atoi function uses this pattern. The function signature of Atoi is:
func Atoi(s string) (int, error)

This means Atoi takes a string argument and returns two values: an integer and an error. If the string can be successfully converted to an integer, Atoi returns the integer and a nil error. If the conversion fails, it returns zero and a non-nil error.
Here is how you would safely use Atoi:

// Atoi converts a stringified number to an integer
i, err := strconv.Atoi("42b")
if err != nil {
    fmt.Println("couldn't convert:", err)
    // because "42b" isn't a valid integer, we print:
    // couldn't convert: strconv.Atoi: parsing "42b": invalid syntax
    // Note:
    // 'parsing "42b": invalid syntax' is returned by the .Error() method
    return
}
// if we get here, then the
// variable i was converted successfully

Formatting Strings Review

A convenient way to format strings in Go is by using the standard library's fmt.Sprintf() function. It's a string interpolation function, similar to Python's f-strings. The %v substring uses the type's default formatting, which is often what you want.
For FLOATS
s := fmt.Sprintf("I am %f years old", 10.523)
// s = I am 10.523000 years old

// The ".2" rounds the number to 2 decimal places
s := fmt.Sprintf("I am %.2f years old", 10.523)
// s = I am 10.52 years old

THE ERROR INTERFACE

Because errors are just interfaces, you can build your own custom types that implement the error interface. Here's an example of a userError struct that implements the error interface:

type userError struct {
    name string
}

func (e userError) Error() string {
    return fmt.Sprintf("%v has a problem with their account", e.name)
}

It can then be used as an error:


func sendSMS(msg, userName string) error {
    if !canSendToUser(userName) {
        return userError{name: userName}
    }
    ...
}

Since this can sometimes be annoying, it might be worth out to check the errors package or use fmt.Errorf -> fmt.Errorf("failed to get user: %w", err)

The Go standard library provides an "errors" package that makes it easy to deal with errors.

Read the godoc for the errors.New() function, but here's a simple example:

var err error = errors.New("something went wrong")


PANIC

As we've seen, the proper way to handle errors in Go is to make use of the error interface. Pass errors up the call stack, treating them as normal values
However, there is another way to deal with errors in Go: the panic function. When a function calls panic, the program crashes and prints a stack trace.

As a general rule, do not use panic!

The panic function yeets control out of the current function and up the call stack until it reaches a function that defers a recover. If no function calls recover, the goroutine (often the entire program) crashes.

func enrichUser(userID string) User {
    user, err := getUser(userID)
    if err != nil {
        panic(err)
    }
    return user
}

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("recovered from panic:", r)
        }
    }()

    // this panics, but the defer/recover block catches it
    // a truly astonishingly bad way to handle errors
    enrichUser("123")
}

If you want your program to cleanly exit in an unrecoverable way, then a good alternative is log.Fatal(), it takes 2 parameters, the first the message you wanna log, and the err
Here's what log.Fatal() does:
    Prints an error message to standard error (stderr)
    Calls os.Exit(1) to terminate the program with a non-zero exit code
    Does NOT run deferred functions (unlike panic/recover)

LOOOPPSSSSSSSS
The basic loop in Go is written in standard C-like syntax:
for INITIAL; CONDITION; AFTER{
  // do something
}

INITIAL is run once at the beginning of the loop and can create
variables within the scope of the loop.

CONDITION is checked before each iteration. If the condition doesn't pass
then the loop breaks.

AFTER is run after each iteration.

for i := 0; i < 10; i++ {
  fmt.Println(i)
}
// Prints 0 through 9

Loops in Go can omit sections of a for loop. For example, the CONDITION (middle part) can be omitted which causes the loop to run forever.

for INITIAL; ; AFTER {
  // do something forever
}

There Is No While Loop in Go
Because Go allows for the omission of sections of a for loop, a while loop is just a for loop that only has a CONDITION.
for CONDITION {
  // do some stuff while CONDITION is true
}

Go supports the standard modulo operator:

7 % 3 // 1

The AND logical operator:

true && false // false
true && true // true

As well as the OR operator:

true || false // true
false || false // false


continue
The continue keyword stops the current iteration of a loop and continues to the next iteration. continue is a powerful way to use the guard clause pattern within loops.
break
The break keyword stops the current iteration of a loop and exits the loop.
**/

type divideError struct {
	dividend float64
}

func (e divideError) Error() string{
	return fmt.Sprintf("can not divide %v by zero", e.dividend)
}

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, divideError{dividend: dividend}
	}
	return dividend / divisor, nil
}

func main(){
	fmt.Println(divide(1,0))
	
}