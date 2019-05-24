package japan

import "fmt"

// Japan ...
type Japan struct {
	Name string
}

// GetName ...
func (j Japan) GetName() {

	fmt.Println("Japan name is: ", j.Name)
}
