package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	arrays()

	slicesEX()

	stringsEX()

	Maps()

	Structs()

	exercises()

}

func exercises() {
	// 1. Write a program that defines a variable named greetings of type slice of strings with the following values: "Hello", "Hola", "नमस्कार", "こんにちは", and "Привіт". Create a subslice containing the first two values; a second subslice with the second, third, and fourth values; and a third subslice with the fourth and fifth values. Print out all four slices.
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}

	subSlice1 := greetings[:2]
	subSlice2 := greetings[1:4]
	subSlice3 := greetings[3:len(greetings)]
	fmt.Println(greetings)
	fmt.Println(subSlice1)
	fmt.Println(subSlice2)
	fmt.Println(subSlice3)

	// 2. Write a program that defines a string variable called message with the value "Hi 👧 and 👦" and prints the fourth rune in it as a character, not a number.
	message := "Hi 👧 and 👦"
	fmt.Println([]rune(message[3:4]))

	// 3. Write a program that defines a struct called Employee with three fields: firstName, lastName, and id. The first two fields are of type string, and the last field (id) is of type int. Create three instances of this struct using whatever values you’d like. Initialize the first one using the struct literal style without names, the second using the struct literal style with names, and the third with a var declaration. Use dot notation to populate the fields in the third struct. Print out all three structs.
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	emp1 := Employee{
		"Moaz", "Ahmed", 1,
	}

	emp2 := Employee{
		firstName: "Moaz",
		lastName:  "Ahmed",
		id:        2,
	}

	var emp3 Employee

	emp3.id = 3
	emp3.firstName = "Moaz"
	emp3.lastName = "Ahmed"

	fmt.Println(emp1, emp2, emp3)

}

func Structs() {
	// when you have related data you want to group together , use a struct
	// a struct can be scoped to any level "Global" , "Local"
	type person struct {
		name string
		age  int
		pet  string
	}

	// it doesn't matter between a nil struct or empty struct
	var lol person = person{}
	var fred person
	fmt.Println(lol, fred)

	// first declaration type with attribute names visible [Named] => ORDER Doesn't matter & YOU don't have to provide all attribute values , the missing ones will default to their zero-values
	var moaz person = person{
		name: "Moaz",
		age:  31,
		pet:  "none",
	}

	fmt.Println(moaz)

	// or you can omit the attribute names [Positional] => MUST be in order & YOU have to provide all attribute values
	moaz = person{
		"Moaz",
		31,
		"none",
	}

	fmt.Println(moaz)

	// access a property in a struct
	fmt.Println(moaz.age)

	fmt.Println("****************")

	// Anonymous Structs
	// Declare a variable that implement a struct type without giving that struct a name
	var anonymousPerson struct {
		name string
		age  int
		pet  string
	}
	anonymousPerson.name = "Moaz"
	anonymousPerson.age = 31
	anonymousPerson.pet = "none"

	fmt.Println(anonymousPerson.name)
	fmt.Println(anonymousPerson.age)
	fmt.Println(anonymousPerson.pet)

	// also shorthand => Must follow it with declaration
	pet := struct {
		name string
		kind string
	}{
		name: "Fido",
		kind: "dog",
	}

	fmt.Println(pet.name)
	fmt.Println(pet.kind)

	fmt.Println("****************")

	// Comparing & Converting Structs
	// Structs can be compared only if they have the same comparable data types [int,string,float,etc] .... slices , maps are not comparable by default so Structs containing them are not comparable.
	// Go does allow you to perform type conversion "your own implement function to do so" from one struct type to another if the fields of both structs have the same names , order and types
	//* Example
	type firstPerson struct {
		name string
		age  int
	}

	firstP := firstPerson{
		name: "Moaz",
		age:  31,
	}

	// you can use type conversion to convert first & second Person [same datatype , same ordering] , but you can't check for equality.
	type secondPerson struct {
		name string
		age  int
	}

	secondP := secondPerson{
		name: "Moaz",
		age:  31,
	}
	secondP = secondPerson(firstP)
	fmt.Println(secondP)

	// however you can't convert or compare the firstPerson to thirdPerson because even if same data type BUT different ordering.
	type thirdPerson struct {
		age  int
		name string
	}

	// you can't covert firstPerson to fourthPerson because the field names doesn't match.
	type forthPerson struct {
		firstPerson string
		age         int
	}

	// finally you can't convert the first Person to the fifth person because there is an additional field "favoriteColor"
	type fifthPerson struct {
		name          string
		age           int
		favoriteColor string
	}

	// you can check for equality between a normal struct and an anonymous struct if they follow the same [datatype,ordering]
	type firstStruct struct {
		name string
		age  int
	}
	f := firstStruct{
		name: "Moaz",
		age:  31,
	}

	var g struct {
		name string
		age  int
	}

	g = f
	fmt.Println(f == g)

}

func Maps() {
	// declaration
	var nilMap map[string]int
	// zero value for nil map is "nil" , length is "0"
	fmt.Println(nilMap)
	fmt.Println(len(nilMap))
	// read from an nil map => zero value for the value type , this case int => 0
	fmt.Println(nilMap["anyKey"])
	// write to a nil map => app will panic
	// nilMap["anyKey"] = 10 => this code will compile but panics

	fmt.Println("****************")

	// short hand declaration
	namesAndAges := map[string]int{
		"MeYo":    1,
		"YoMe":    2,
		"Damadem": 3,
		"Ayood":   4,
	}
	fmt.Println(namesAndAges)

	fmt.Println("****************")

	// declaration with various data types
	teams := map[string][]string{
		// notice that you can omit the redundant types in a literal
		"team1": {"1", "2", "3"},
		// must be separated by comma "," even the last element added.
		"team2": {"4", "5", "6"},
		"team3": {"7", "8", "9"},
	}
	fmt.Println(teams)

	// add to a map
	teams["team4"] = []string{"10", "11", "12"}
	fmt.Println(teams)

	fmt.Println("****************")

	// if you already know how many elements you are going to add in a map but don't know the values yet , use make
	// the made map will still have a length of 0 and it will grow as you add keys:values to it
	makeMap := make(map[string]int, 10)
	fmt.Println(makeMap)
	fmt.Println(len(makeMap))

	fmt.Println("****************")

	// reading and writing to a map
	totalWins := map[string]int{}
	totalWins["Orcas"] = 1
	totalWins["Lions"] = 2
	fmt.Println(totalWins["Orcas"])
	fmt.Println(totalWins["Lions"])
	fmt.Println(totalWins["someTeamThatDoesNotExist"]) //=> will default to the default value of the assigned "value" data type.
	totalWins["Lions"]++
	fmt.Println(totalWins)
	totalWins["Orcas"] = 100
	fmt.Println(totalWins)
	// adding value to a key that doesn't exist
	totalWins["someTeamThatDoesNotExist"]++ //=> the key does exist now and will have this value as an initial value.
	fmt.Println(totalWins)

	fmt.Println("****************")

	// the comma ok idiom
	// to tell whether a key is not found in a map or it's found but has a zero value as it's declaration , how to tell the difference between them ? use it.
	// it can be named any key , but the idiom is to use "ok" for that bool key.
	m := map[string]int{
		"hello": 5,
		"world": 0,
	}
	// key is found with no zero value assigned => ok == true , it's found key
	v, ok := m["hello"]
	fmt.Println(v, ok)

	// key is found with zero value assigned => ok == true , it's found key
	v, ok = m["world"]
	fmt.Println(v, ok)

	// key is not found [ofc it has zero value] => ok == false , key not found
	v, ok = m["goodBye"]
	fmt.Println(v, ok)

	fmt.Println("****************")

	// deleting from a map
	newMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Println(newMap)
	// deletes a key from a map , if the key doesn't exist or the map is nil , no problem , it won't return an error or panics
	delete(newMap, "one")
	fmt.Println(newMap)

	fmt.Println("****************")

	// emptying a map
	m = map[string]int{
		"hello": 5,
		"world": 10,
	}

	fmt.Println(m, len(m))
	clear(m)
	fmt.Println(m, len(m))

	fmt.Println("****************")

	// Comparing Maps
	m = map[string]int{
		"hello": 5,
		"world": 10,
	}

	n := map[string]int{
		"world": 5,
		"hello": 10,
	}

	fmt.Println(maps.Equal(m, n))

	fmt.Println("****************")

	// using maps as sets , Go doesn't include Set "similar like maps" but doesn't include repetitive elements.
	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}

	// to access elements of a slice/array
	for _, v := range vals {
		intSet[v] = true
	}
	fmt.Println(len(vals), len(intSet)) // 11 | 8
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])
	if intSet[100] {
		fmt.Println("100 is in the set")
	}

	// to access keys:values of a map/set
	for i, v := range intSet {
		fmt.Println(i, v)
	}

}

func stringsEX() {
	var s string = "hello there"
	// you can extract a single value from a string , same as array or slice
	// first you get the value as a UTF-8 value
	var b byte = s[6]
	fmt.Println(b)
	// then convert it to get the string value of it
	var s6 string = string(b)
	fmt.Println(s6)

	fmt.Println("********************")

	// Substring [supports 1 byte characters] , doesn't work well with emojis
	s = "Hello there"
	var s2 string = s[4:7] // "o t"
	var s3 string = s[:5]  // "hello"
	var s4 string = s[6:]  // " there"
	fmt.Println(s)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)

	fmt.Println("********************")

	// Length
	fmt.Println(len(s))

	// strings can be converted back and forth to [slice/rune] or a slice of bytes or slice of runes
	s = "Hello ,  "
	var bs []byte = []byte(s)
	var rs []rune = []rune(s)
	fmt.Println(bs)
	fmt.Println(rs)
}

func slicesEX() {
	// initialization
	sl1 := []int{1, 2, 3}
	fmt.Println(sl1)

	sl1 = []int{1, 5: 4, 6, 10: 100, 15} // 1 , 0 , 0 , 0 , 0 , 4 , 6 , 0 , 0 , 0 , 100 , 15
	fmt.Println(sl1)

	// read index value
	fmt.Println(sl1[0])
	// fmt.Println(sl1[15]) => this will compile normally but will panic because "out of bounds".

	// declaring a slice without a literal , the zero value for non-initialized slice is "nil"
	// sl2 := []int{}
	var sl2 []int
	fmt.Println(sl2)

	// you can't compare slices , at least not directly
	// fmt.Println(sl1 == sl2) => won't compile , invalid operation
	// the only thing you can compare slice to is "nil"
	fmt.Println(sl2 == nil)

	// however , since Go 1.21 , there are 2 introduced functions that can be used to compare slices , but they have to be the same datatype []int == []int
	fmt.Println(slices.Equal(sl1, sl2))

	// length : total number of elements for array/slice , length for nil slice is 0
	fmt.Println(len(sl2))

	// append , can work with separate elements or a whole slice , remember that if you don't assign the appended value to a slice , you will get a compile time error as you are not using that appended value , also the append function will copy the passed slice and return a copy of the appended slice , learn pointers and come back here to see if there is a solution to that issue.

	// append(sl2, 1, 2, 3) => won't compile , not used

	// append separate elements
	sl2 = append(sl2, 1, 2, 3)
	fmt.Println(sl2)
	// append a whole slice
	sl2 = append(sl2, sl1...)
	fmt.Println(sl2)

	// capacity : the number of consecutive memory locations reserved.
	// each time you append/add to slice , the capacity is increased by one or more values , when the capacity == length , the next time you append will make Go Runtime to allocate a new backing array , copying all data from the old array underlying array , ditching that old underlying array and append all new values to the end of the backing array.
	var x []int
	fmt.Println(x, len(x), cap(x)) // [] 0 0
	x = append(x, 10)
	fmt.Println(x, len(x), cap(x)) // [10] 1 1
	x = append(x, 20)
	fmt.Println(x, len(x), cap(x)) // [10 20] 2 2
	x = append(x, 30)
	fmt.Println(x, len(x), cap(x)) // [10 20 30] 3 4
	x = append(x, 40)
	fmt.Println(x, len(x), cap(x)) // [10 20 30 40] 4 4
	x = append(x, 50)
	fmt.Println(x, len(x), cap(x)) // [10 20 30 40 50] 5 8

	// notice how the capacity increase and surpasses the length at a certain elements added
	// better than this , it's better if you already know what's going to be added , so you can make a slice using the correct initial capacity , using "make"

	// make : it allows you to create an empty slice with reserved length and optional reserved capacity
	sl3 := make([]int, 5)
	// a common mistake is to use "append" as an initializer to a "make" made slice , this will append the values added to the end of the slice , leaving all initialized length values as zero-values.
	sl3 = append(sl3, 10)
	fmt.Println(sl3) // [0 0 0 0 0 10]

	// however if you make a slice with 0 length , you can append to that slice
	sl5 := make([]int, 0, 10)
	sl5 = append(sl5, 5, 6, 7, 8)
	fmt.Println(sl5)

	// specify an initial capacity
	sl4 := make([]int, 5, 10)
	fmt.Println(sl4)

	// never specify a capacity less than the length , it's a compile time error , and if you specify a variable that turns out to be less than the length , the app will panic at runtime.

	fmt.Println("********************")

	// empty a slice : this will reset the current existing elements back to their zero-value , however the length and capacity of the slice remain the same
	sl6 := []string{"MeYo", "YoMe", "Damadem", "Ayood"}
	fmt.Println(sl6, len(sl6), cap(sl6))
	clear(sl6)
	fmt.Println(sl6, len(sl6), cap(sl6))

	// this is a nil slice
	var slNil []int
	fmt.Println(slNil, len(slNil), cap(slNil))
	fmt.Println(slNil == nil)

	// this is a non-nil slice
	slNonNil := []int{}
	fmt.Println(slNonNil, len(slNonNil), cap(slNonNil))
	fmt.Println(slNil == nil)

	// hint : if you have slice starting values or if the slice values aren't going to change , then a slice literal is a good choice , else , use a nil slice "slice without literal".

	// hint : if you a good idea how large your slice will be but don't know what those values will be , use "make" to make a slice

	fmt.Println("********************")

	// slicing a slice : create a sliced slice out of a slice , lol
	sl10 := []string{"a", "b", "c", "d"}
	sl20 := sl10[:2]
	sl30 := sl10[1:]
	sl40 := sl10[1:3]
	sl50 := sl10[:]
	fmt.Println("sl10:", sl10) // a , b , c , d
	fmt.Println("sl20:", sl20) // a , b
	fmt.Println("sl30:", sl30) // b , c , d
	fmt.Println("sl40:", sl40) // b , c
	fmt.Println("sl50:", sl50) // a , b , c , d

	fmt.Println("********************")

	// when you take a slice from a slice , you are not taking a copy of the data. instead, you now have two variables that are sharing memory , this means that changes to an element in a slice affect all slices that share that element.
	sl100 := []string{"a", "b", "c", "d"}
	sl200 := sl100[:2] // a , b
	sl300 := sl100[1:] // b , c , d
	sl100[1] = "y"     // a , y , c , d
	sl200[0] = "x"     // x , y
	sl300[1] = "z"     // y , z , d

	fmt.Println("sl100:", sl100) // sl100 : x , y , z , d
	fmt.Println("sl200:", sl200) // sl200 : x , y
	fmt.Println("sl300:", sl300) // sl300 : y , z , d

	fmt.Println("********************")

	// and it gets extra confusing when combined with append
	slx := []string{"a", "b", "c", "d"}
	sly := slx[:2] // a , b
	fmt.Println(cap(slx), cap(sly))
	sly = append(sly, "z") // a , b , z

	fmt.Println("slx:", slx) // a , b , z , d => "c" is replaced by "z"
	fmt.Println("sly:", sly) // a , b , z

	fmt.Println("********************")

	// another example
	slx1 := make([]string, 0, 5)
	slx1 = append(slx1, "a", "b", "c", "d")
	sly1 := slx1[:2] // a , b
	slz1 := slx1[2:] // c , d
	fmt.Println(cap(slx1), cap(sly1), cap(slz1))

	sly1 = append(sly1, "i", "j", "k") // a , b , i , j , k
	slx1 = append(slx1, "x")           // a , b , i , j , k , x
	slz1 = append(slz1, "y")           // i , j , k , x , y

	fmt.Println("slx1:", slx1) // a , b , i , j , y
	fmt.Println("sly1:", sly1) // a , b , i , j , y
	fmt.Println("slz1:", slz1) //  i , j , y

	fmt.Println("********************")

	// full slice [3 parts] expression protects against append
	// first part : low index [inclusive]
	// second part : high index [exclusive]
	// third part : capacity of the new slice
	slx2 := make([]string, 0, 5)
	slx2 = append(slx2, "a", "b", "c", "d")
	sly2 := slx2[:2:2]  // a , b | underlying array (a,b)
	slz2 := slx2[2:4:4] // c , d | underlying array (c,d)
	fmt.Println("slx2:", slx2, cap(slx2))
	fmt.Println("sly2:", sly2, cap(sly2))
	fmt.Println("slz2:", slz2, cap(slz2))

	fmt.Println("********************")

	// to check if a slice contains an element
	fmt.Println(slices.Contains(slx2, "a"))
	// to get an element's index from a slice => returns -1 if not exist
	fmt.Println(slices.Index(slx2, "c"))

	fmt.Println("********************")

	// copy : make an independent copy of a slice that doesn't affect the original slice
	sliX := []int{1, 2, 3, 4, 6, 7, 8, 9}
	sliY := make([]int, 4)
	// copy destination first , source seconds , gives back the copied elements , limited to the minimum slice used in the operation , cuz u may wanna copy a large slice into a small slice
	num := copy(sliY, sliX)
	fmt.Println(sliY, num)

	fmt.Println("********************")

	// another example
	sliX1 := []int{1, 2, 3, 4}
	sliY1 := make([]int, 2)
	// if you don't need the minimum number of elements copied , no need to assign it or use it.
	copy(sliY1, sliX1)
	fmt.Println(sliX1)
	fmt.Println(sliY1)

	fmt.Println("********************")

	// another example
	x = []int{1, 2, 3, 4}
	num = copy(x[:3], x[1:]) //=> 1,2,3 to be 2,3,4 ====> final : 2,3,4,4
	fmt.Println(x, num)      //=> 2,3,4,4, | 3 minimum copies

	fmt.Println("********************")
	// another example
	x = []int{1, 2, 3, 4}
	d := [4]int{5, 6, 7, 8}
	y := make([]int, 2)
	copy(y, d[:])
	fmt.Println(y) //=> 5 , 6
	copy(d[:], x)
	fmt.Println(d) //=> 1,2,3,4

	fmt.Println("********************")

	// convert arrays to slices => arrName[:]
	xArray := [4]int{5, 6, 7, 8}
	xSlice := xArray[:]
	fmt.Println(xSlice)

	fmt.Println("********************")

	// another example
	xArr := [4]int{5, 6, 7, 8}
	ySli := xArr[:2]
	zSli := xArr[2:]
	xArr[0] = 10
	fmt.Println("xArr:", xArr) // 10,6,7,8
	fmt.Println("ySli:", ySli) // 10,6
	fmt.Println("zSli:", zSli) // 7,8

	fmt.Println("********************")

	// Converting Slices to Arrays of the same type [must]
	// Using Type conversion
	// also it will copy the slice into the array , modifying the array's elements won't modify the slice elements.
	xSlice = []int{1, 2, 3, 4}
	xArray = [4]int(xSlice)
	// can't use [...] at compile time , will give an error.
	// smallArray := [...]int(xSlice) => won't compile
	smallArray := [2]int(xSlice)
	xSlice[0] = 10
	fmt.Println(xSlice)
	fmt.Println(xArray)
	fmt.Println(smallArray)

	fmt.Println("*******************")

	// The compiler won't detect if you are using an array with a large size , if that occurred , will panic at runtime.
	// panicArray := [5]int(xSlice) => will compile but will have a runtime error [panic]
	// fmt.Println(panicArray)

	fmt.Println("*******************")

	// if you want an array to share a slice memory , use pointers
	xSlice = []int{1, 2, 3, 4}
	xArrayPointer := (*[4]int)(xSlice)
	xSlice[0] = 10             // 0,2,3,4
	xArrayPointer[1] = 20      // 0,1,3,4
	fmt.Println(xSlice)        // 0,1,3,4
	fmt.Println(xArrayPointer) // 0,1,3,4

}

func arrays() {
	// initialize array size [option 1]
	var x [3]int
	fmt.Println(x)

	x = [3]int{1, 2, 3}
	fmt.Println(x)

	y := [12]int{1, 5: 4, 6, 10: 100, 15} // 1 , 0 , 0 , 0 , 0 , 4 , 6 , 0 , 0 , 0 , 100 , 15
	fmt.Println(y)

	// [option2] you can also initialize the array with ... as the number of elements , but you have to make it literal.
	var z = [...]int{1, 2, 3}
	fmt.Println(z)

	// you can compare arrays of the same elements length
	arr1 := [...]int{1, 2, 3}
	arr2 := [...]int{1, 2, '3'}
	fmt.Println(arr1 == arr2)

	// read value
	fmt.Println(arr1[0])
	// fmt.Println(arr1[3]) won't compile => out of bounds

	// length of the array
	fmt.Println(len(arr1))

}
