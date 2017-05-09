package main

import (
	"fmt"
)

func main() {
	//map declaration
	var personSalary1 map[string]int //zero value of map is nil
	if personSalary1 == nil {
		fmt.Println("map is nil. Going to make personSalary1 map.")
		personSalary1 = make(map[string]int)
	}

	//short hand declaration
	fmt.Println("\nShort hand declaration:-")
	personSalary2 := make(map[string]int)
	personSalary2["steve"] = 12000
	personSalary2["jamie"] = 15000
	personSalary2["mike"] = 9000
	fmt.Println("personSalary2 map contents:", personSalary2)

	//map declaration and initialization
	fmt.Println("\nMap initialization and declaration:-")
	personSalary3 := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary3["mike"] = 9000
	fmt.Println("personSalary map contents:", personSalary3)

	//accessing items of a map
	fmt.Println("\nAccessing items of a map:-")
	employee := "jamie"
	fmt.Println("Salary of", employee, "is", personSalary3[employee])

	//accessing an item which is not present returns the zero value
	fmt.Println("\nAccess an element which is not present returns the zero value:-")
	fmt.Println("Salary of joe is", personSalary3["joe"])

	//checking whether a key is present or not
	fmt.Println("\nChecking whether a key is present or not:-")
	newEmp := "joe"
	value, ok := personSalary3[newEmp]
	if ok == true {
		fmt.Println("Salary of", newEmp, "is", value)
	} else {
		fmt.Println(newEmp, "not found")
	}

	//iterating over all element using for range
	fmt.Println("\niterating over all elements using for range:-")
	for key, value := range personSalary3 {
		fmt.Printf("personSalary3[%s] = %d\n", key, value)
	}

	//deleting items
	fmt.Println("\nDeleting items:-")
	fmt.Println("map before deletion", personSalary3)
	delete(personSalary3, "steve")
	fmt.Println("map after deletion", personSalary3)

	//length of map
	fmt.Println("\nLength of map:-")
	fmt.Println("length of personSalary3 is", len(personSalary3))

	/*Maps are reference types. When a map is assigned to another,
	they point to the same internal datastructure.
	Changes made to one map will be reflected in the other
	*/
	fmt.Println("\nMaps are reference type:-")
	map1 := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	fmt.Println("Original map1", map1)
	map2 := map1
	map2["mike"] = 18000
	fmt.Println("map1 changed", map2)

}
