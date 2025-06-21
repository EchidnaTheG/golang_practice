package main

import "fmt"



func main() {
	fname := "John"
	lname := "Doe"
	registerDate:="3/21/2023"
	planType := "Deluxe"
	isOnUpgradeList := false
	numOfDevices := 6
	numOfRelatives := 2
	logMessage := fmt.Sprintf("Name: %v %v, Date Registered: %v, Subscription Type: %v, OnUpgradeList: %v, Number of Devices: %v, Number of Relatives: %v",fname,lname, registerDate,planType, isOnUpgradeList, numOfDevices, numOfRelatives)
	fmt.Println(logMessage)

}
