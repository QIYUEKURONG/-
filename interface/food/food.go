package food

import "fmt"

// Banana is a fruit
type Banana struct {
	Name string
}

// DescribeFood function can describe a food
func (b *Banana) DescribeFood() {
	fmt.Println("Banana is a yello fruit.and can eat.can treat constipation")
}

// PrintfFood function can peintf a food
func (b *Banana) PrintfFood() {
	fmt.Println("I am a fruit of banana")
}
