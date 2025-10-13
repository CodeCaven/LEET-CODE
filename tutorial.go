package main

import (
	"fmt"
)

//  go run .\tutorial.go from terminal

// go is not oo language but you can define methods on types and structs
// go is complied (can build an executable) and statically typed language

// function with parameters
func familyName(fname string, age int) {
	fmt.Println("Hello", age, "year old", fname, "Refsnes")
}

// main function
func main() {

	// variable declaration
	// there are diffs b/w var and :=

	var student1 string = "John" //type is string
	var student2 = "Jane"        //type is inferred
	x := 2                       //type is inferred

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)

	// data types
	var a bool = true    // Boolean
	var b int = 5        // Integer
	var c float32 = 3.14 // Floating point number
	var d string = "Hi!" // String

	fmt.Println("Boolean: ", a)
	fmt.Println("Integer: ", b)
	fmt.Println("Float:   ", c)
	fmt.Println("String:  ", d)

	// arrays
	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}

	fmt.Println(arr1)
	fmt.Println(arr2)

	// inferred length arrays
	arr3 := [...]int{9, 10, 11}
	fmt.Println(arr3)
	fmt.Println(arr3[0])   // accessing first element
	fmt.Println(len(arr3)) // length of array

	// slices are like arrays but dynamic in size
	myslice1 := []int{}
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))
	fmt.Println(myslice1)

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)

	// decisions
	time := 22
	if time < 10 {
		fmt.Println("Good morning.")
	} else if time < 20 {
		fmt.Println("Good day.")
	} else {
		fmt.Println("Good evening.")
	}

	// you can only use for loops in Go!
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}

	// for arrays and slices you can use range
	fruits := [3]string{"apple", "orange", "banana"}
	for idx, val := range fruits {
		fmt.Printf("%v\t%v\n", idx, val)
	}

	familyName("Liam", 3)
	familyName("Jenny", 14)
	familyName("Anja", 30)

	// maps are key value pairs like dict in python
	myMap := map[string]int{"apple": 5, "banana": 10, "orange": 15}
	fmt.Println(myMap)

}
