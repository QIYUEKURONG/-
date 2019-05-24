package china

import "fmt"

// Person ...
type China struct {
	Name string
	Sex  string
}

// GetName ...
func (n China) GetName() {
	fmt.Println("China name is:", n.Name)
}

// GetSex ...
func (n China) GetSex() {
	fmt.Println("the name sex is:", n.Sex)
}
