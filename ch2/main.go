package main

import "fmt"

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

}
