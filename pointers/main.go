package main

import "fmt"

// This function returns a pointer to an integer
// Which causes escape analysis to allocate the variable on the heap
// This is dangerous because the variable a is allocated on the stack
// and will be deallocated when the function returns
// causing the pointer to point to an invalid memory location
// and causing a runtime error
func createPointerDanger() *int {
	var a int = 10
	return &a
}

func createPointerSafe() *int {
	ptr := new(int)
	a := 100
	*ptr = a
	return ptr
}

func main() {
	newPointer1 := createPointerDanger()

	fmt.Println(*newPointer1)

	newPointer2 := createPointerSafe()

	fmt.Println(*newPointer2)

	fmt.Println("=======================================")

	var a int = 10

	var b *int // pointer to int

	b = &a // assign address of a to b

	fmt.Printf("Address of a: %x\n", &a) // %x is used to print in hexadecimal format
	fmt.Printf("Value of b: %x\n", b)
	fmt.Printf("Dereferencing b: %d\n", *b)
	*b = 20 // change the value at the address stored in pointer b
	fmt.Printf("Value of a: %d\n", a)
}

/*
* Escaping: When a variable is allocated on the heap instead of the stack
* This is dangerous because the variable a is allocated on the stack
* and will be deallocated when the function returns
* causing the pointer to point to an invalid memory location
* and causing a runtime error


* Zero value of a pointer is nil
* A pointer can be dereferenced using the * operator
* A pointer can be assigned to another pointer
* A pointer can be assigned to a variable of the same type


* Always guard against nil pointers by checking if the pointer is nil before dereferencing it
* Dereferencing a nil pointer will cause a runtime error / panic in Go

* Panic: A panic is a runtime error that causes the program to crash
 */
