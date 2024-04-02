package main

import (
	"fmt"
	"runtime"
	"sort"
	"strings"
)

func main() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.Compiler)
	fmt.Println(runtime.Version())

	// ========================Strings========================
	fmt.Println("========================Strings========================")
	greeting := "Hello, there friends!"
	// fmt.Println(strings.Count(greeting, "o"))
	// fmt.Println(strings.HasPrefix(greeting, "Hello"))
	// fmt.Println((strings.Contains(greeting, "play")))
	// fmt.Println(strings.ReplaceAll(greeting, " ", "_"))
	// fmt.Println(strings.ToUpper(greeting))
	// fmt.Println(strings.ToLower(greeting))
	// fmt.Println(strings.Index(greeting, "ll"))
	fmt.Println(strings.Split(greeting, " "))

	// Original string is not modified
	fmt.Println("Original string value: ", greeting)

	// ========================Sort========================
	fmt.Println("========================Sort========================")

	ages := []int{45, 20, 35, 30, 75, 60, 50, 25, 40, 55}

	sort.Ints(ages)

	// Original slice is modified
	fmt.Println(ages)

	index := sort.SearchInts(ages, 30)

	indexDoesNotExist := sort.SearchInts(ages, 100)

	fmt.Println(index)

	fmt.Println("length of ages", len(ages))
	fmt.Println(indexDoesNotExist) // returns the length of the slice

	// ========================Sort Strings========================
	fmt.Println("========================Sort Strings========================")

	names := []string{"Xander", "Buffy", "Angel", "Willow", "Giles"}

	sort.Strings(names)

	fmt.Println(names)

	indexString := sort.SearchStrings(names, "Willow")

	indexStringDoesNotExist := sort.SearchStrings(names, "Zander")

	fmt.Println(indexString)

	fmt.Println(indexStringDoesNotExist) // returns the length of the slice
}
