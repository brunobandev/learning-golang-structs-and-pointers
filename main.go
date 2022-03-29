package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Halpert",
		contactInfo: contactInfo{
			email:   "jim@halpert.com",
			zipCode: 910983,
		},
	}

	// Ref a memory address (&jim)
	// jimPointer := &jim
	// jimPointer.updateName("Jimmy")

	// Or simply
	jim.updateName("Jimmy")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	// Get a value of memory address (*pointerToPerson)
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// |===============================================
// | Memory Address  | Value
// |===============================================
// | 0001            | person{firstName: "bruno"}
// |-----------------------------------------------
// | 0002            |
// |-----------------------------------------------
// | 0003            | person{firstName: "bruno"}
// |-----------------------------------------------
// | ...

// Turn ADDRESS into VALUE with *address
// Turn VALUE into ADDRESS with &value

// =================================
// | value types | reference types |
// =================================
// | int         | slices          |
// | float       | maps            |
// | string      | channels        |
// | bool        | pointers        |
// | structs     | functions       |
// =================================

// Value types = Use pointers to change these things in a function
// Reference types = Dont worry with pointers with these
