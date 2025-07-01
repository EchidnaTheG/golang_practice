package main

import "fmt"

// INTERFACES PRACTICE
// The length of a string can be obtained using the len function, which returns the number of bytes.
/**

Interfaces allow you to focus on what a type does rather than how it's built. They can help you write more flexible and reusable code by defining behaviors (like methods) that different types can share. This makes it easy to swap out or update parts of your code without changing everything else.

Interfaces are just collections of method signatures. A type "implements" an interface if it has methods that match the interface's method signatures.

In the following example, a "shape" must be able to return its area and perimeter. Both rect and circle fulfill the interface.

type shape interface {
  area() float64
  perimeter() float64
}

type rect struct {
    width, height float64
}
func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perimeter() float64 {
    return 2*r.width + 2*r.height
}

type circle struct {
    radius float64
}
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
    return 2 * math.Pi * c.radius
}

When a type implements an interface, it can then be used as that interface type

func printShapeData(s shape) {
	fmt.Printf("Area: %v - Perimeter: %v\n", s.area(), s.perimeter())
}


Because we say the input is of type shape, we know that any argument must implement the .area() and .perimeter() methods.

As an example, because the empty interface doesn't require a type to have any methods at all, every type automatically implements the empty interface, written as:
interface{}

Interface Implementation

Interfaces are implemented implicitly.

A type never declares that it implements a given interface. If an interface exists and a type has the proper methods defined, then the type automatically fulfills that interface.
A quick way of checking whether a struct implements an interface is to declare a function that takes an interface as an argument. If the function can take the struct as an argument, then the struct implements the interface.
A type implements an interface by implementing its methods. Unlike in many other languages, there is no explicit declaration of intent, there is no "implements" keyword.

Implicit interfaces decouple the definition of an interface from its implementation. You may add methods to a type and in the process be unknowingly implementing various interfaces, and that's okay.
This is different from most other languages, where you have to explicitly assign an interface type to an object, like with Java:
class Circle implements Shape


Name Your Interface Parameters!!
Consider the following interface:
type Copier interface {
  Copy(string, string) int
}

This is a valid interface, but based on the code alone, can you deduce what kinds of strings you should pass into the Copy function?

We know the function signature expects 2 string types, but what are they? Filenames? URLs? Raw string data? For that matter, what the heck is that int that's being returned?

Let's add some named parameters and return data to make it more clear.

type Copier interface {
  Copy(sourceFile string, destinationFile string) (bytesCopied int)
}
When working with interfaces in Go, every once-in-awhile you'll need access to the underlying type of an interface value. You can cast an interface to its underlying type using a type assertion.  Read below

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}

// here we use type assertions!! ok returns true if its the type its being cast to, example -> f, ok := e.(randomstruct) - > ok will be true if it is that struct, false if not, look below, f is trying to accest the struct itself to become it

func getExpenseReport(e expense) (string, float64) {
	em, ok := e.(email)
	if ok {
		return em.toAddress, em.cost()
	}

	sm, ok := e.(sms)
	if ok {
		return sm.toPhoneNumber, sm.cost()
	}

	return "", 0.0
}
**/



/**
WAIT WAIT WAIT!!! FOR INTERFACE TYPE ASSERTIONS, USE A SWITCH STATEMENT!!
Specifically, a type switch:
A type switch makes it easy to do several type assertions in a series.
A type switch is similar to a regular switch statement, but the cases specify types instead of values.

here is how the previous getExpenseReport would look with a type switch

func getExpenseReport(e expense) (string, float64) {
// 'v' will automatically be of the correct type within each case
	switch v := e.(type) { 
	case email:
		return v.toAddress, v.cost()
	case sms:
		return v.toPhoneNumber, v.cost()
	default:
		return "", 0.0
	}
}
**/

// Interfaces are hard, 3 big rules of thumb
// 1. Keep Interfaces small, they define minimal behavior necessary to represent an idea or concept
// 2. Interfaces Should Have No Knowledge of Satisfying Types - aka, Interfaces should only describe what’s required, not who uses them. An interface says what a type must do to fit in, but it shouldn’t know about or care about which specific types actually use it.
// or in another set of words, The interface just sets the rules, the types that follow those rules come later and don’t need to be mentioned by the interface itself, get it?? Just use type assertions in the interface itself for this
// 3. Interfaces Are Not Classes!!!

func main(){
	
}
