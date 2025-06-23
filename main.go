package main

import "fmt"

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



func triple(x,y,z int)(int, int, int){
	x *=3
	y*= 3
	z*= 3
	return x, y, z
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
}

