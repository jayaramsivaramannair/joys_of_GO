package main

import (
	"fmt"
)

type Employee struct {
	Name string
	ID string
}

//Method is available on the Employee Struct
func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

//Note About Embedded Fields:
//1. In the below case, Manager contains a field of type Employee but no name is assigned to that field. 
//2. This means that Employee is an "Embedded Field"
//3. Embedded a Field within another struct means that any fields or methods declared on an embedded field 
// are promoted to the containing struct and can be invoked directly on the containing struct as show below.

//A new struct Manager
type Manager struct {
	Employee
	Reports [] Employee
}

func main() {
	m := Manager {
		Employee: Employee {
			Name: "Bob Bobson",
			ID: "12345",
		},

		Reports: [] Employee {},
	}

	//In this case, ID and Description are directly available on the Manager struct because it contains an embedded field of type Employee.
	fmt.Println(m.ID)
	fmt.Println(m.Description())
	

}