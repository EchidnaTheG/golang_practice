package main
/**
ARRAYS and SLICES,

Arrays in Go

    Simple/primitive types (int, string, bool, arrays, structs) → copied by value
    Complex/reference types (slices, maps, channels, pointers, functions) → share underlying data

Arrays are fixed-size groups of variables of the same type. For example, [4]string is an array of 4 values of type string.

To declare an array of 10 integers:

var myInts [10]int
or to declare an initialized literal:

primes := [6]int{2, 3, 5, 7, 11, 13}

Slices in Go

99 times out of 100 you will use a slice instead of an array when working with ordered lists.

Arrays are fixed in size. Once you make an array like [10]int you can't add an 11th element.

A slice is a dynamically-sized, flexible view of the elements of an array.

The zero value of slice is nil.

Non-nil slices always have an underlying array, though it isn't always specified explicitly. To explicitly create a slice on top of an array we can do:

primes := [6]int{2, 3, 5, 7, 11, 13}
mySlice := primes[1:4]
// mySlice = {3, 5, 7}

The syntax is:

arrayname[lowIndex:highIndex]
arrayname[lowIndex:]
arrayname[:highIndex]
arrayname[:]
Where lowIndex is inclusive and highIndex is exclusive.

lowIndex, highIndex, or both can be omitted to use the entire array on that side of the colon.
If a function takes a slice argument, any changes it makes to the elements of the slice will be visible to the caller, analogous to passing a pointer (we'll cover pointers later) to the underlying array
Every slice in Go points to a real array in memory.

The slice keeps track of: where the array starts, how long it is (length), and how much space is available (capacity).

When you append and exceed capacity, Go:

    Allocates a new (usually bigger) array,

    Copies the old elements over,

    Makes your slice point to this new array.

This is why you can keep appending to slices, even though the original array was a certain size.

a := []int{1, 2, 3, 4}
b := a[:2]       // b = [1, 2], shares the same array as a

b[0] = 99        // changes a[0] too
fmt.Println(a)   // [99 2 3 4]
fmt.Println(b)   // [99 2]

We can create a new slice using the make function:

// func make([]T, len, cap) []T
mySlice := make([]int, 5, 10)

// the capacity argument is usually omitted and defaults to the length
mySlice := make([]int, 5)

Slices created with make will be filled with the zero value of the type.

If we want to create a slice with a specific set of values, we can use a slice literal:

mySlice := []string{"I", "love", "go"}

Notice the square brackets do not have a 3 in them. If they did, you'd have an array instead of a slice.

Length

The length of a slice is simply the number of elements it contains. It is accessed using the built-in len() function:

mySlice := []string{"I", "love", "go"}
fmt.Println(len(mySlice)) // 3

Capacity

The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice. It is accessed using the built-in cap() function:

mySlice := []string{"I", "love", "go"}
fmt.Println(cap(mySlice)) // 3

INDEXING
mySlice := []string{"I", "love", "go"}
fmt.Println(mySlice[2]) // go

mySlice[0] = "you"
fmt.Println(mySlice) // [you love go]

func getMessageCosts(messages []string) []float64 {
	messageCosts := make([]float64, len(messages))
	for i := 0; i < len(messages); i++ {
		cost := float64(len(messages[i])) * 0.01
		messageCosts[i] = cost
	}
	return messageCosts
}

The capacity of a slice, accessible by the built-in function cap, reports the maximum length the slice may assume. 
Variadic

Many functions, especially those in the standard library, can take an arbitrary number of final arguments. This is accomplished by using the "..." syntax in the function signature.

A variadic function receives the variadic arguments as a slice.

func concat(strs ...string) string {
    final := ""
    // strs is just a slice of strings
    for i := 0; i < len(strs); i++ {
        final += strs[i]
    }
    return final
}

func main() {
    final := concat("Hello ", "there ", "friend!")
    fmt.Println(final)
    // Output: Hello there friend!
}

The spread operator allows us to pass a slice into a variadic function. The spread operator consists of three dots following the slice in the function call.

func printStrings(strings ...string) {
	for i := 0; i < len(strings); i++ {
		fmt.Println(strings[i])
	}
}

func main() {
    names := []string{"bob", "sue", "alice"}
    printStrings(names...)
}

The built-in append function is used to dynamically add elements to a slice:

func append(slice []Type, elems ...Type) []Type

If the underlying array is not large enough, append() will create a new underlying array and point the returned slice to it.

Notice that append() is variadic, the following are all valid:

slice = append(slice, oneThing)
slice = append(slice, firstThing, secondThing)
slice = append(slice, anotherSlice...)


Range

Go provides syntactic sugar to iterate easily over elements of a slice:

for INDEX, ELEMENT := range SLICE {
}

The element is a copy of the value at that index.

For example:

fruits := []string{"apple", "banana", "grape"}
for i, fruit := range fruits {
    fmt.Println(i, fruit)
}
// 0 apple
// 1 banana
// 2 grape

Slice of Slices

Slices can hold other slices, effectively creating a matrix, or a 2D slice.

rows := [][]int{}

func createMatrix(rows, cols int) [][]int {
	matrix := [][]int{}
	for i := 0; i< rows; i++{
		row := [] int {}
		for j := 0; j < cols; j++{
			row = append(row, i*j)
		}
		matrix = append(matrix, row)
	}
	return matrix
}

Tricky Slices

The append() function changes the underlying array of its parameter AND returns a new slice. This means that using append() on anything other than itself is usually a BAD idea.

// dont do this!
someSlice = append(otherSlice, element)

Take a look at these head-scratchers:
Example 1: Works As Expected

a := make([]int, 3)
fmt.Println("len of a:", len(a))
fmt.Println("cap of a:", cap(a))
// len of a: 3
// cap of a: 3

b := append(a, 4)
fmt.Println("appending 4 to b from a")
fmt.Println("b:", b)
fmt.Println("addr of b:", &b[0])
// appending 4 to b from a
// b: [0 0 0 4]
// addr of b: 0x44a0c0

c := append(a, 5)
fmt.Println("appending 5 to c from a")
fmt.Println("addr of c:", &c[0])
fmt.Println("a:", a)
fmt.Println("b:", b)
fmt.Println("c:", c)
// appending 5 to c from a
// addr of c: 0x44a180
// a: [0 0 0]
// b: [0 0 0 4]
// c: [0 0 0 5]

With slices a, b, and c, 4 and 5 seem to be appended as we would expect. We can even check the memory addresses and confirm that b and c point to different underlying arrays.
Example 2: Something Fishy

i := make([]int, 3, 8)
fmt.Println("len of i:", len(i))
fmt.Println("cap of i:", cap(i))
// len of i: 3
// cap of i: 8

j := append(i, 4)
fmt.Println("appending 4 to j from i")
fmt.Println("j:", j)
fmt.Println("addr of j:", &j[0])
// appending 4 to j from i
// j: [0 0 0 4]
// addr of j: 0x454000

g := append(i, 5)
fmt.Println("appending 5 to g from i")
fmt.Println("addr of g:", &g[0])
fmt.Println("i:", i)
fmt.Println("j:", j)
fmt.Println("g:", g)
// appending 5 to g from i
// addr of g: 0x454000
// i: [0 0 0]
// j: [0 0 0 5]
// g: [0 0 0 5]

In this example, however, when 5 is appended to g it overwrites j's fourth index because j and g point to the same underlying array. The append() function only creates a new array when there isn't any capacity left. We created i with a length of 3 and a capacity of 8, which means we can append 5 items before a new array is automatically allocated.

Again, to avoid bugs like this, you should always use the append function on the same slice the result is assigned to:

mySlice := []int{1, 2, 3}
mySlice = append(mySlice, 4)



way to deal with characters (runes in golang)

func isValidPassword(password string) bool {
	length := len(password)
	hasUpper := false
	hasNumber := false

	for i := 0; i < length; i++ {
		char := password[i]
		if char >= 'A' && char <= 'Z' {
			hasUpper = true
		}
		if char >= '0' && char <= '9' {
			hasNumber = true
		}
	}

	return length >= 5 && length <= 12 && hasUpper && hasNumber
}

**/

