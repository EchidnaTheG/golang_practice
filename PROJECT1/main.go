/**
💡 Project: Inventory Management System

🔧 What you’ll build:
An application that tracks items in a store’s inventory. Each item will be a struct.
You’ll store multiple items using slices. You’ll define interfaces for actions like adding, removing, or listing inventory.
Arrays can be used for fixed-size containers (like storage bins or warehouses).

Concept    How It’s Used
Structs    Define Item, Category, or Warehouse
Interfaces    Define behavior like Storable, Printable, or Searcher
Slices    Hold dynamic list of Items in your inventory
Project: Inventory Management System
**/

package main
import ("fmt" ; "math/rand"	)





type Item struct{
	item string  	//literal type of item -> toy, plushie, hammer
	Category string //One of 4 cats -> Inventory, Entertainment, Staff, Maintenance
	Warehouse string //The different warehouses around -> 5: RX01, RX02, RX03, RX04, CDC1
	id int64 // unique id -> 4232291121 : this max length
}

func (i Item) info() string{
	item := i.item
    category := i.Category
    warehouse := i.Warehouse
    id := i.id

    if item == "" {
        item = "nil"
    }
    if category == "" {
        category = "nil"
    }
    if warehouse == "" {
        warehouse = "nil"
	}
	return fmt.Sprintf("item: %v | category: %v | warehouse: %v | id: %v", item, category, warehouse, id)

}
func (i Item) Storable() bool{
	return true
}






type System struct{
	db []Item
	initializedDB bool
}






func (s *System) createItem(item, Category, Warehouse string ){

	id := rand.Int63n(23312231)
	s.db =append(s.db,Item{item: item, Category:Category, Warehouse:Warehouse, id: id} ) 

}

func (s System) readItems() string{
	wholeStr :=""
	for i, item := range s.db{
	wholeStr += fmt.Sprintf("%v| %v\n", i, item.info())
	}
	return wholeStr
}

func (s System) readItemByIndex(rowNum int) string{
	wholeStr :=""
	counter := 0
	for i, item := range s.db{
		if counter == rowNum{
	wholeStr = fmt.Sprintf("%v| %v\n", i, item.info())
		}
		counter += 1
	}
	return wholeStr
}

func (s *System) createDB(){
	if !s.initializedDB{
	s.db = []Item{}
	}
}

type Storable interface{
	info() string
	Storable() bool
}






func main(){
	system := System{initializedDB: false}
	system.createDB()
	system.createItem("pizza","Inventory","")
	system.createItem("Pizza Cutter", "Inventory", "RX01")
	system.createItem("Cheese Grater", "Inventory", "CDC1")
	system.createItem("Oven Mitt", "Maintenance", "RX04")
	system.createItem("Pizza Box", "Maintenance", "RX01")
	system.createItem("Pepperoni Slicer", "Inventory", "RX02")
	system.createItem("Mozzarella Block", "Inventory", "RX03")
	system.createItem("Tomato Sauce Can", "Inventory", "CDC1")
	system.createItem("Delivery Scooter", "Entertainment", "RX04")
	system.createItem("Cash Register", "Staff", "RX03")
	system.createItem("Arcade Machine", "Entertainment", "RX01")
	system.createItem("Uniform Shirt", "Staff", "RX02")
	system.createItem("Mop Bucket", "Maintenance", "RX01")
	system.createItem("Flour Bag", "Inventory", "RX04")
	system.createItem("Sauce Ladle", "Inventory", "RX02")
	system.createItem("Pizza Peel", "Inventory", "RX03")
	system.createItem("Rolling Pin", "Inventory", "CDC1")
	system.createItem("Deep Fryer", "Maintenance", "RX04")
	system.createItem("Receipt Printer", "Staff", "RX02")
	system.createItem("Plastic Crates", "Inventory", "RX01")
	system.createItem("Walkie Talkie", "Staff", "CDC1")
	system.createItem("First Aid Kit", "Maintenance", "RX03")
	system.createItem("Cleaning Spray", "Maintenance", "RX01")
	system.createItem("Loyalty Card Scanner", "Staff", "RX04")
	system.createItem("Measuring Cup", "Inventory", "RX02")
	system.createItem("Spatula", "Inventory", "RX03")
	system.createItem("Apron", "Staff", "CDC1")
	system.createItem("Sound System", "Entertainment", "RX01")
	system.createItem("Grease Trap", "Maintenance", "RX03")
	system.createItem("Serving Tray", "Inventory", "RX02")
	system.createItem("Fire Extinguisher", "Maintenance", "RX04")
	system.createItem("Chef Hat", "Staff", "RX01")
	system.createItem("Cutting Board", "Inventory", "RX03")
	system.createItem("Timer", "Maintenance", "RX02")
	system.createItem("Thermometer", "Maintenance", "RX04")
	system.createItem("Chair", "Inventory", "CDC1")
	system.createItem("Table", "Inventory", "RX01")
	system.createItem("Napkin Dispenser", "Inventory", "RX02")
	system.createItem("Hand Sanitizer", "Maintenance", "RX03")
	system.createItem("Toolbox", "Maintenance", "RX04")
	system.createItem("Fan", "Maintenance", "RX01")
	system.createItem("Light Bulb", "Maintenance", "RX02")
	system.createItem("Shelf", "Inventory", "RX03")
	system.createItem("Storage Bin", "Inventory", "RX04")
	system.createItem("Trash Can", "Maintenance", "RX01")
	system.createItem("Soap Dispenser", "Maintenance", "CDC1")
	system.createItem("Credit Card Reader", "Staff", "RX02")
	system.createItem("Mask Box", "Staff", "RX03")
	system.createItem("Pizza Dough Ball", "Inventory", "RX01")
	system.createItem("Drink Cooler", "Inventory", "RX04")
	system.createItem("Receipt Roll", "Staff", "CDC1")
	fmt.Println(system.readItems())
	fmt.Println(system.readItemByIndex(45))
}