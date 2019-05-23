package main

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	var arrary [5]int
	testIfValue(arrary)
	result := ifOldValue(arrary, arrary)
	if result {
		fmt.Println("ok")
	}
}
