package main

import "fmt"

func main() {
	// var ages [3]int = [3]int{20, 25, 30}

	// Shorthand syntax
	var ages = [3]int{20, 25, 30}

	names := [4]string{"Yoshi", "Mario", "Peach", "Bowser"}

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	fmt.Println("==========================Slices==========================")

	var scores = []int{100, 50, 60} // same as scores := []int{100, 50, 60}

	newScores := append(scores, 85) // append() returns a new slice

	scores[0] = 95

	fmt.Println("newScores: ", newScores, len(newScores))

	fmt.Println("scores: ", scores, len(scores))

	numbers := make([]int, 3) // make() creates a slice with a length of 3

	fmt.Println("numbers: ", numbers, len(numbers))

	fmt.Println("=======================Slice Ranges=======================")

	rangeOne := names[1:3]  // [inclusive:exclusive]
	rangeTwo := names[2:]   // [inclusive:end]
	rangeThree := names[:3] // [start:exclusive]
	rangeFour := names[:]   // [start:end]

	fmt.Println("rangeOne: ", rangeOne)
	fmt.Println("rangeTwo: ", rangeTwo)
	fmt.Println("rangeThree: ", rangeThree)
	fmt.Println("rangeFour: ", rangeFour)

	// rangeOne[1] = "Goku" // changes the original array

	rangeOne = append(rangeOne, "Goku") // changes the original array

	fmt.Println("rangeOne: ", rangeOne)
	fmt.Println("names: ", names)

	// Deep copy
	rangeNew := append([]string{}, names[1:3]...)

	rangeNew[0] = "Vegeta" // does not change the original array

	fmt.Println("names: ", names)

	fmt.Println("rangeNew: ", rangeNew)
}

/*
* In Go, arrays are fixed-size. This means that once you define an array, you cannot change its size.

* Slices are a flexible alternative to arrays. They are created using brackets [] and can be resized.
* Slices are like dynamic arrays with special features.
* Under the hood, a slice is just a reference to an underlying array.

* The append() function is used to add elements to a slice. It returns a new slice with the added elements.
* The original slice is not modified.

* The make() function is used to create a slice with a specific length and capacity.
* The length is the size of the slice, and the capacity is the size of the underlying array.
* The capacity is optional and defaults to the length.

* The len() function returns the length of a slice.
* The cap() function returns the capacity of a slice.

* Slices are used more often than arrays in Go because of their flexibility and ease of use.

* Slice ranges are used to get a subset of a slice.
* The syntax is slice[start:stop], where start is the index to start from and stop is the index to stop before.
* If you change the elements of a slice range, the original slice will be modified as they share the same underlying array.

 */
