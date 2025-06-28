package main

import "fmt"


//GO IS BLOCK SCOPED, NOT FUNCTION SCOPED, A BLOCK BEING WELL, BLOCK, there is the main block, and new ones made by curlybraces {}
func Logger(fname, lname, registerDate, planType string, isOnUpgradeList bool, numOfDevices, numOfRelatives int) string {
	logMessage := fmt.Sprintf(
		"\n Name: %v %v\n Date Registered: %v\n Subscription Type: %v\n OnUpgradeList: %v\n Number of Devices: %v\n Number of Relatives: %v\n",
		fname, lname, registerDate, planType, isOnUpgradeList, numOfDevices, numOfRelatives)
	return logMessage
}

func Request(httpCode int, serverNode string) int {
	if httpCode == 200 && serverNode != "" {
		return 0
	}else{
	return 1
}
}

//golang uses a lot of the early return style instead of nested flow statements
/**for example
func example( x, y int)(int){
	if x == 1 && y == 0{
	return 10
	}
	if x....
	etc, the idea is to return early instead of setting in a deep nested set of conditions
}
**/


func applyAddFunction( sum func(x int) int, number int) int {
	return sum(number)
}



func triple(x,y,z int)(int, int, int){
	x *=3
	y*= 3
	z*= 3
	return x, y, z
}

// REMINDERRR, NO FUNCTION CALLS OR STUFF GOING ON OUTSIDE OF MAIN< EVERYTHING OUTSIDE OF MAIN EXISTS TO BE USED IN SOMETHING IN MAIN< OR BY SOMETHING THAT WILL EVENTRUALLY BE USED IN MAIN
//DEFER KEYWORK : BASICALLY A FUNCTION CALL THAT ONLY EXECUTES ONCE THE FUNCTION ITS CALLED IN RETURNS ANY VALUE , OR WHEN THE FUNCTION ITS CALLED ENDS, IF MULTIPLE DEFERS, THEY FUNCTION IN STACK ORDER (BOTTOM TO TOP), IN OTHER WORDS< THIS IS LIKE CALLING A FUNCTION ALWAYS BEFORE RETURTNING OR FINSIHING THE FUNCTION WHERE ITS LCOATED AT
func write_to_db( ) {
	defer fmt.Println("DB CLOSED")
	fmt.Println("Writing to DB...")
	fmt.Println("Writing to DB...")
	fmt.Println("Writing to DB...")
	fmt.Println("Writing to DB...")
	fmt.Println("Writing to DB...")
	fmt.Println("DB OPERATION DONE!")
}


// A closure is a function that references variables from outside its own function body. The function may access and assign to the referenced variables.
func adder() func(int) int {
	sum := 0
	return func(num int) int{
		sum += num
		return sum
	}
}
// Function currying is a concept from functional programming and involves partial application of functions. It allows a function with multiple arguments to be transformed into a sequence of functions, each taking a single argument. For example
func MATTH() {
  squareFunc := selfMath(multiply)
  doubleFunc := selfMath(add)

  fmt.Println(squareFunc(5))
  // prints 25

  fmt.Println(doubleFunc(5))
  // prints 10
}

func multiply(x, y int) int {
  return x * y
}

func add(x, y int) int {
  return x + y
}

func selfMath(mathFunc func(int, int) int) func (int) int {
  return func(x int) int {
    return mathFunc(x, x)
  }
}


type person struct{
age int
name string
gender string
}



func main() {
	RESPONSEVALUE := Request(2010, "America: Miami")
	switch RESPONSEVALUE  {
	case 0:
		fmt.Println("=================================================")
		fmt.Println(Logger("John", "Doe", "3/21/2023", "Deluxe", false, 6, 2))
		fmt.Println("=================================================")
		fmt.Println("O P E R A T I O N	C O M P L E T E D")
	case 1:
		fmt.Println("=======================================")
		fmt.Println("N O	 D A T A	 F O U N D")			
		fmt.Println("=======================================")
		fmt.Println("O P E R A T I O N A L	F A I L U R E")
	}
	fmt.Println(triple(42,12,11))
	test_null_return1, test_null_return2, _ := triple(31,32,24323421)
	println(test_null_return1,test_null_return2)
	fmt.Println(applyAddFunction(func(x int) int {
	return x + x}, 2))
	write_to_db()
	myPerson := person{}
	myPerson.name = "Mike"
	fmt.Println(myPerson)
}

