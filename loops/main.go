package main

import "fmt"

func main() {
	// x := 0
	// for x < 5 {
	// 	fmt.Println("value of x is:", x)
	// 	x++
	// }

	// for i := 5; i > 0; i-- {
	// 	fmt.Println("value of i is:", i)
	// }

	// for {
	// 	fmt.Println("This will run forever")
	// }

	names := []string{"Bob", "John", "Alice", "Michael", "Jenny"}

	// for i := 0; i < len(names); i++ {
	// 	fmt.Printf("The name at index %v is %v\n", i, names[i])
	// }

	// for index, name := range names {
	// 	fmt.Printf("The name at index %v is %v\n", index, name)
	// }

	// for _, name := range names { // if you don't need the index, you can use an underscore
	// 	fmt.Println(name)
	// }

	for _, name := range names {
		name = "Mr. " + name // this will not change the original slice
	}

	fmt.Println(names)

	for index := range names {
		names[index] = "Mr. " + names[index] // this will change the original slice
	}

	fmt.Println(names)

}

/*
* In go, there is no while loop. Instead, you can use a for loop with a condition to achieve the same effect.

* The for loop in go can be used in three ways:
	1. with a single condition
	2. with a traditional initial/condition/after for loop
	3. with no condition (infinite loop)

* The for loop can also be used to iterate over a slice or an array using the range keyword.
  The range keyword returns the index and the value at that index.

* The range keyword can also be used to iterate over a map. In this case, the range keyword returns the key and the value at that key.
*/
