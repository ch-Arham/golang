package main

import "fmt"

func getName() (string, string) {
	return "Arham", "Khawar"
}

func sum(a, b int) int {
	return a + b
}

func getArea(width, height int) (area int) {
	area = width * height
	return // Naked return - returns the named return value alternatively we can also return area
	// return area
}

func main() {
	firstName, _ := getName()

	var result int = sum(3, 5)

	newArea := getArea(10, 5)

	fmt.Println("First Name:", firstName)
	fmt.Println("Sum:", result)
	fmt.Println("Area:", newArea)
}

/*
* The blank identifier is used to ignore the second value returned by the getName function.
* In Go, we can return multiple values from a function.
* In the above example, we have a function getName that returns two strings.
* We can use the blank identifier to ignore the second value as go does not allow unused variables.

* func sum(a, b int) int is a function signature that takes two integers and returns an integer.
 */
