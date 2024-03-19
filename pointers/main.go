package main

import "fmt"

// The function createLocalVariable creates a local variable x on the stack and returns its address.
// This can cause a dangling pointer if the variable is deallocated when the function returns.
// Also the memory address can be reused for other purposes and can cause undefined behavior.
func createLocalVariable() *int {
	var x int = 42
	return &x // Return the address of the local variable x
}

// The function createHeapVariable creates a heap-allocated variable x and returns its address.
// This ensures that the variable persists beyond the function's scope and can be safely accessed.
// Escape analysis will determine that the variable needs to be allocated on the heap.
func createHeapVariable() *int {
	x := new(int)
	*x = 42
	return x // Return the address of the heap-allocated variable x
}

func main() {
	newPointer1 := createLocalVariable()

	fmt.Println(*newPointer1)

	newPointer2 := createHeapVariable()

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

When you return the address of a local variable from a function and assign it to a pointer,
it will cause the variable to escape, meaning it will be allocated on the heap instead of the stack.
This is because the variable's lifetime extends beyond the scope of the function,
 and it needs to persist even after the function returns.

However, returning pointers to local variables from functions can indeed lead to dangling pointers if not handled carefully.
If the local variable is deallocated when the function returns, the pointer will become invalid,
pointing to memory that may be reused for other purposes.
Accessing the memory through such a dangling pointer can result in undefined behavior,
including runtime panics or memory corruption.
*/
