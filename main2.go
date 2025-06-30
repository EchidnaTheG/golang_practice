package main

import "fmt"


//struct practice
/** Why Use Pointers to Structs in Go?
1. Efficiency (Performance)

    Passing a pointer (*Person) to a function avoids copying the entire struct.

    This is faster and uses less memory, especially for large structs.

structs memory layout- structs sit in memory in a contiguous block, with fields placed one after another as defined in the struct. Long story short, make structs start from largest field to smallest in terms of memory size
example

this is 4 bytes
type stats struct {
	Reach    uint16
	NumPosts uint8
	NumLikes uint8
}

An example of how NOT to do structs would be this
this is 6 bytes, 2 wasted due to padding
type stats struct {
	NumPosts uint8
	Reach    uint16
	NumLikes uint8
}
	
**/

type Car struct{
	model string
	dateAcquired string
	miles int
	cylinders int
	computer struct
{CPU string
 CORES int
	GPU string 
	Wattage int
	isOn bool 
	age int
}}

type computer struct{
	CPU string
	CORES int
	GPU string 
	Wattage int
	isOn bool 
	age int
}

func updateCarModelStruct(pointer *Car){
	pointer.model = "Honda"
}

// there is no inheritance in golang, but you can make a struct data-only inheritance...sort of, by using embedded structs

type car struct {
  brand string
  model string
}

type truck struct {
  // "car" is embedded, so the definition of a
  // "truck" now also additionally contains all
  // of the fields of the car struct
  car
  bedSize int
}

// We can also make "methods" in golang, kinda.
 //Methods are just functions that have a receiver. A receiver is a special parameter that syntactically goes before the name of the function.
type rect struct {
  width int
  height int
}

// area has a receiver of (r rect)
// rect is the struct
// r is the placeholder
func (r rect) area() int {
  return r.width * r.height
}

var r = rect{
  width: 5,
  height: 10,
}






func main (){
	//making an instance of a struct
	Mycar := Car{
		model : "Toyota",
		dateAcquired: "22/06/2014",
		miles : 78212,
		cylinders : 6,
		computer: computer {
			CPU: "INTEL",
			CORES : 8,
			GPU: "NVIDIA",
			Wattage: 750,
			isOn: true,
			age: 2,
		},
	}
Mycarpointer := &Mycar
//lets update a field of the struct, by pointer is best since we modify the original
//we call the function we made for this
updateCarModelStruct(Mycarpointer)
fmt.Println(Mycar)
//it works! and is efficient, optimal
//we right now are using named structs, we always should unless we truly don't ever require the struct to be used more than once

lanesTruck := truck{
  bedSize: 10,
  car: car{
    brand: "Toyota",
    model: "Tundra",
  },
}

fmt.Println(lanesTruck.brand) // Toyota
fmt.Println(lanesTruck.model) // Tundra
fmt.Println(r.area())
// prints 50
}
