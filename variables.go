package main

import "fmt"

func main() {
	var name string = "John" // explicit type
	// var name = "John" implicit type
	//  name := 'John' // assignment and implicit type
	name = "Arham"
	fmt.Println(name)
	fmt.Printf("%T\n", name)
}

// go has uint8, uint16, uint32, uint64, int8, int16, int32, int64
// go has float32, float64
// go has complex64, complex128
// go has bool
// go has string
// go has rune (int32)
// go has byte (uint8)
// go has uint (uint64)
// go has int (int64)
