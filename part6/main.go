package main

/**
MAPS:

Maps are similar to JavaScript objects, Python dictionaries, and Ruby hashes. Maps are a data structure that provides key->value mapping.

The zero value of a map is nil.

We can create a map by using a literal or by using the make() function:

ages := make(map[string]int)
ages["John"] = 37
ages["Mary"] = 24
ages["Mary"] = 21 // overwrites 24

ages = map[string]int{
  "John": 37,
  "Mary": 21,
}

The len() function works on a map, it returns the total number of key/value pairs.

ages = map[string]int{
  "John": 37,
  "Mary": 21,
}
fmt.Println(len(ages)) // 2


Insert an Element

m[key] = elem

Get an Element

elem = m[key]

Delete an Element

delete(m, key)

Check If a Key Exists

elem, ok := m[key]

    If key is in m, then ok is true and elem is the value as expected.
    If key is not in the map, then ok is false and elem is the zero value for the map's element type.

Any type can be used as the value in a map, but keys are more restrictive.
 map keys may be of any type that is comparable
 comparable types are boolean, numeric, string, pointer, channel, and interface types, and structs or arrays that contain only those types.
absent from the list are slices, maps, and functions; these types cannot be compared using ==, and may not be used as map keys.

use structs over nested maps if possible, specifically struct as keys

You can combine an if statement with an assignment operation to use the variables inside the if block:

names := map[string]int{}
missingNames := []string{}

if _, ok := names["Denna"]; !ok {
    // if the key doesn't exist yet,
    // append the name to the missingNames slice
    missingNames = append(missingNames, "Denna")
}
 An attempt to fetch a map value with a key that is not present in the map will return the zero value for the type of the entries in the map. For instance, if the map contains integers, looking up a non-existent key will return 0. A set can be implemented as a map with value type bool. Set the map entry to true to put the value in the set, and then test it by simple indexing.

attended := map[string]bool{
    "Ann": true,
    "Joe": true,
    ...
}

if attended[person] { // will be false if person is not in the map
    fmt.Println(person, "was at the meeting")
}

Like slices, maps hold references to an underlying data structure. If you pass a map to a function that changes the contents of the map, the changes will be visible in the caller.

Map Literals

Maps can be constructed using the usual composite literal syntax with colon-separated key-value pairs, so it's easy to build them during initialization.

var timeZone = map[string]int{
    "UTC":  0*60*60,
    "EST": -5*60*60,
    "CST": -6*60*60,
    "MST": -7*60*60,
    "PST": -8*60*60,

To delete a map entry, use the delete built-in function, whose arguments are the map and the key to be deleted. It's safe to do this even if the key is already absent from the map.

delete(timeZone, "PDT")  // Now on Standard Time
A function can mutate the values stored in a map and those changes affect the original values



func getNameCounts(names []string) map[rune]map[string]int {
	// Create the outer map. It maps a rune (first letter) to another map.
	nameCounts := make(map[rune]map[string]int)

	for _, name := range names {
		// Convert the name string to a slice of runes to properly handle Unicode.
		runes := []rune(name)

		// Check if the inner map for the first rune exists yet.
		// This is necessary because map values default to nil — you must initialize them before use.
		if _, ok := nameCounts[runes[0]]; !ok {
			// Initialize the inner map if it doesn't exist.
			nameCounts[runes[0]] = make(map[string]int)
		}

		// Now safe to increment the count for this name in the inner map.
		nameCounts[runes[0]][name] += 1
	}

	// Return the fully built nested map.
	return nameCounts
}


**/
