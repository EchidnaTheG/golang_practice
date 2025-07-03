
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
Arrays    Model fixed storage (e.g., storage bin with limited capacity)
Project: Inventory Management System
**/



package main


// This is the the struct for every item in the system, the idea is
// create a item struct say...a Pizza! Well, the convention for naming will be:
//Pizza4232291121 -> the item will be first named, like what the literal item is, followed by the id number, both of these  things will be set as properties of the struct
// but they also will be in the name, id is generated randomly, well, first checking if it already exists. this way, we can have multiple same item types

type Item struct{
	item string  	//literal type of item -> toy, plushie, hammer
	Category string //One of 4 cats -> Inventory, Entertainment, Staff, Maintenance
	Warehouse string //The different warehouses around -> 5: RX01, RX02, RX03, RX04, CDC1
	id int // unique id -> 4232291121 : this max length
}


