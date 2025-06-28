package main

import "fmt"


//struct practice
/** Why Use Pointers to Structs in Go?
1. Efficiency (Performance)

    Passing a pointer (*Person) to a function avoids copying the entire struct.

    This is faster and uses less memory, especially for large structs.**/


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
}
