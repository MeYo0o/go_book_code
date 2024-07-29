package main

import "fmt"

// compiler won't stop you from creating un-read package level variables , however , as they are not used , they won't be included in the compiled binary
var zzz int = 10

func main() {
	var i int = 10
	i *= 2
	fmt.Println(i)

	var z float64 = 10.5
	fmt.Println(z / 0)

	//! Rune is an alias for Int32 , but you should consider the following when using them
	// Good , the type name matches the usage
	var myFirstInitial rune = 'J'
	fmt.Println(myFirstInitial)

	// Legal but not recommended
	var mySecondInitial int32 = 'J'
	fmt.Println(mySecondInitial)

	//! Go doesn't have automatic type promotion , you have to do it manually
	// this is an example of manually conversion between int & float
	var x int = 10
	var y float64 = 30.2
	var sum1 float64 = float64(x) + y
	fmt.Println(sum1)
	var sum2 int = x + int(y)
	fmt.Println(sum2)

	// another example for int,byte
	var x1 int = 10
	var b1 byte = 100
	var sum3 int = x1 + int(b1)
	fmt.Println(sum3)
	var sum4 byte = byte(x1) + b1
	fmt.Println(sum4)

	// const
	const name = "moaz"
	// name = "ahmed" => won't compile
	fmt.Println(name)

	// x2 := 1
	// y2 := 2
	// const z2 = x2 + y2 => won't compile as the both x2,y2 are not known at compile time [const]
	const x3 = 1
	const y3 = 2
	const z3 = x3 + y3
	fmt.Println(z3)

	// untyped constants
	const x4 = 10
	// x4 = 20 => won't compile
	fmt.Println(x4)
	// but it can be assigned to other type that share the same value assignment
	var y4 int = x4
	var y5 byte = x4
	var y6 float64 = x4
	fmt.Println(y4, y5, y6)

	// typed constants , go won't stop you from making un-read constants , however , as they are not used , they won't be included in the compiled binary
	const x5 int = 10
	// var y55 float64 = x5 => can't be assigned to float64 as (x5) is a typed int constant.

	//Exercises
	exercises()
}

func exercises() {
	// 1. Write a program that declares an integer variable called i with the value 20. Assign i to a floating-point variable named f. Print out i and f.
	var i int = 20
	var f float64 = float64(i)
	fmt.Println(i, f)

	// 2. Write a program that declares a constant called value that can be assigned to both an integer and a floating-point variable. Assign it to an integer called i and a floating-point variable called f. Print out i and f.
	const value = 10
	i = value
	f = value
	fmt.Println(i, f)

	// 3. Write a program with three variables, one named b of type byte, one named smallI of type int32, and one named bigI of type uint64. Assign each variable the maximum legal value for its type; then add 1 to each variable. Print out their values.
	var b byte
	var smallI int32
	var bigI uint64

	b = 255
	smallI = 2_147_483_647
	bigI = 18_446_744_073_709_551_615
	fmt.Println(b, smallI, bigI)
	b++
	smallI++
	bigI++
	fmt.Println(b, smallI, bigI)

}
