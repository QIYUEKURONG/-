package main

import "testing"

func TestSlice(t *testing.T) {
	s2 := make([]int, 10)
	for i := 0; i < 100; i++ {
		printfLenandCap(s2)
		s2 = append(s2, 2*i+1)
	}
}
