package main

import (
	"fmt"
	"maps"
	_ "slices"
	"strings"
)

func main() {
	arrays()

	slicesEX()

	stringsEX()

	Maps()

	Structs()
}

// Structs demonstrates various struct operations and concepts
func Structs() {
	// when you have related data you want to group together, use a struct
	// a struct can be scoped to any level "Global" , "Local" // [Original Note]
	type person struct {
		name string
		age  int
		pet  string
	}

	// it doesn't matter between a nil struct or empty struct // [Original Note]
	var lol person = person{}
	var fred person
	fmt.Println(lol, fred)

	// first declaration type with attribute names visible [Named] => ORDER Doesn't matter & YOU don't have to provide all attribute values , the missing ones will default to their zero-values // [Original Note]
	var moaz person = person{
		name: "Moaz",
		age:  31,
		pet:  "none",
	}
	fmt.Println(moaz)

	// or you can omit the attribute names [Positional] => MUST be in order & YOU have to provide all attribute values // [Original Note]
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
	// Declare a variable that implement a struct type without giving that struct a name // [Original Note]
	var anonymousPerson struct {
		name string
		age  int
		pet  string
	}
	anonymousPerson.name = "Moaz"
	anonymousPerson.age = 31
	anonymousPerson.pet = "none"
	fmt.Println(anonymousPerson.name, anonymousPerson.age, anonymousPerson.pet)

	// also shorthand => Must follow it with declaration // [Original Note]
	pet := struct {
		name string
		kind string
	}{
		name: "Fido",
		kind: "dog",
	}
	fmt.Println(pet.name, pet.kind)

	fmt.Println("****************")

	// Comparing & Converting Structs
	// Structs can be compared only if they have the same comparable data types [int,string,float,etc] .... slices , maps are not comparable by default so Structs containing them are not comparable. // [Original Note]
	// Go does allow you to perform type conversion "your own implement function to do so" from one struct type to another if the fields of both structs have the same names , order and types // [Original Note]
	//* Example
	type firstPerson struct {
		name string
		age  int
	}
	firstP := firstPerson{
		name: "Moaz",
		age:  31,
	}

	// you can use type conversion to convert first & second Person [same datatype , same ordering] , but you can't check for equality. // [Original Note]
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

	// however you can't convert or compare the firstPerson to thirdPerson because even if same data type BUT different ordering. // [Original Note]
	type thirdPerson struct {
		age  int
		name string
	}

	// you can't covert firstPerson to fourthPerson because the field names doesn't match. // [Original Note]
	type forthPerson struct {
		firstPerson string
		age         int
	}

	// finally you can't convert the first Person to the fifth person because there is an additional field "favoriteColor" // [Original Note]
	type fifthPerson struct {
		name          string
		age           int
		favoriteColor string
	}

	// you can check for equality between a normal struct and an anonymous struct if they follow the same [datatype,ordering] // [Original Note]
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

// Maps demonstrates various map operations and concepts
func Maps() {
	// declaration
	var nilMap map[string]int
	// zero value for nil map is "nil" , length is "0" // [Original Note]
	fmt.Println(nilMap, len(nilMap))
	// read from a nil map => zero value for the value type , this case int => 0 // [Original Note]
	fmt.Println(nilMap["anyKey"])
	// write to a nil map => app will panic // [Original Note]
	// nilMap["anyKey"] = 10 => this code will compile but panics // [Original Note]

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
		// notice that you can omit the redundant types in a literal // [Original Note]
		"team1": {"1", "2", "3"},
		// must be separated by comma "," even the last element added. // [Original Note]
		"team2": {"4", "5", "6"},
		"team3": {"7", "8", "9"},
	}
	fmt.Println(teams)

	// add to a map
	teams["team4"] = []string{"10", "11", "12"}
	fmt.Println(teams)

	fmt.Println("****************")

	// if you already know how many elements you are going to add in a map but don't know the values yet , use make // [Original Note]
	// the made map will still have a length of 0 and it will grow as you add keys:values to it // [Original Note]
	makeMap := make(map[string]int, 10)
	fmt.Println(makeMap, len(makeMap))

	fmt.Println("****************")

	// reading and writing to a map
	totalWins := map[string]int{}
	totalWins["Orcas"] = 1
	totalWins["Lions"] = 2
	fmt.Println(totalWins["Orcas"], totalWins["Lions"])
	fmt.Println(totalWins["someTeamThatDoesNotExist"]) //=> will default to the default value of the assigned "value" data type. // [Original Note]
	totalWins["Lions"]++
	fmt.Println(totalWins)
	totalWins["Orcas"] = 100
	fmt.Println(totalWins)
	// adding value to a key that doesn't exist
	totalWins["someTeamThatDoesNotExist"]++ //=> the key does exist now and will have this value as an initial value. // [Original Note]
	fmt.Println(totalWins)

	fmt.Println("****************")

	// the comma ok idiom
	// to tell whether a key is not found in a map or it's found but has a zero value as it's declaration , how to tell the difference between them ? use it. // [Original Note]
	// it can be named any key , but the idiom is to use "ok" for that bool key. // [Original Note]
	m := map[string]int{
		"hello": 5,
		"world": 0,
	}
	// key is found with no zero value assigned => ok == true , it's found key // [Original Note]
	v, ok := m["hello"]
	fmt.Println(v, ok)

	// key is found with zero value assigned => ok == true , it's found key // [Original Note]
	v, ok = m["world"]
	fmt.Println(v, ok)

	// key is not found [ofc it has zero value] => ok == false , key not found // [Original Note]
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
	// deletes a key from a map , if the key doesn't exist or the map is nil , no problem , it won't return an error or panics // [Original Note]
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

	// using maps as sets , Go doesn't include Set "similar like maps" but doesn't include repetitive elements. // [Original Note]
	intSet := map[int]bool{}
	vals := []int{5, 10, 5, 10, 1, 2}
	for _, v := range vals {
		intSet[v] = true
	}
	fmt.Println(len(vals), len(intSet))

	// in Go 1.21 there is slices.Clone => clone1 := slices.Clone(vals) // [Original Note]
	// but for previous versions, you can make a copy from one slice to another as follows : // [Original Note]
	vals2 := make([]int, len(vals))
	copy(vals2, vals)
	fmt.Println(vals2)

	fmt.Println("****************")

	// Iterating over a Map
	// you can't rely on the order of elements of a map because Go randomize them on every iteration. // [Original Note]
	m = map[string]int{
		"hello": 5,
		"world": 10,
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
	for k := range m {
		fmt.Println(k)
	}
	for _, v := range m {
		fmt.Println(v)
	}

	fmt.Println("****************")

	// Getting the keys and values from a map
	// using "Maps" package in Go 1.21 // [Original Note]
	m = map[string]int{
		"hello": 5,
		"world": 10,
	}
	// keys := maps.Keys(m)
	// values := maps.Values(m)
	// fmt.Println(keys, values)
	// // but for the previous versions // [Original Note]
	// keys = []string{}
	// for k := range m {
	// 	keys = append(keys, k)
	// }
	// slices.Sort(keys) // sorted so that the print statements are predictable
	// fmt.Println(keys)

	// values = []int{}
	// for _, v := range m {
	// 	values = append(values, v)
	// }
	// fmt.Println(values)
}

// stringsEX demonstrates various string operations and concepts
func stringsEX() {
	// splitting on delimiter or substring
	s := "a/b/c"
	// strings.Fields and strings.FieldsFunc can be used if a whitespace exists. // [Original Note]
	parts := strings.Split(s, "/")
	fmt.Println(parts)

	// contains
	fmt.Println(strings.Contains(s, "/"))

	// replace
	fmt.Println(strings.ReplaceAll(s, "/", "|"))

	// joining on delimiter or substring
	fmt.Println(strings.Join(parts, "/"))

	// counting a substring
	fmt.Println(strings.Count(s, "/"))

	// lowercase
	fmt.Println(strings.ToLower(s))

	// uppercase
	fmt.Println(strings.ToUpper(s))

	// trimming whitespace or other characters
	fmt.Println(strings.TrimSpace("  a/b/c  "))
	fmt.Println(strings.Trim("xxaxbbxx", "x"))

	// removing a substring
	fmt.Println(strings.ReplaceAll("foo", "o", ""))

	// converting to and from byte slices
	b := []byte(s)
	fmt.Println(string(b))
}

// slicesEX demonstrates various slice operations and concepts
func slicesEX() {
	// creation
	s := []string{"a", "b", "c"}
	fmt.Println(s)

	// accessing
	fmt.Println(s[1])

	// modifying
	s[1] = "x"
	fmt.Println(s)

	// slicing
	fmt.Println(s[1:3])

	// appending
	s = append(s, "d")
	fmt.Println(s)

	// copying
	s2 := make([]string, len(s))
	copy(s2, s)
	fmt.Println(s2)

	// iterating
	for i, v := range s {
		fmt.Println(i, v)
	}
	for _, v := range s {
		fmt.Println(v)
	}

	// length and capacity
	fmt.Println(len(s), cap(s))

	// make
	s3 := make([]string, 3)
	fmt.Println(s3)

	// nil slice
	var s4 []string
	fmt.Println(s4, len(s4), cap(s4))

	// empty slice
	s5 := []string{}
	fmt.Println(s5, len(s5), cap(s5))

	// append slice to another slice
	s6 := []string{"e", "f"}
	s = append(s, s6...)
	fmt.Println(s)

	// slicing beyond the length but within the capacity
	s = s[:cap(s)]
	fmt.Println(s)

	// using slices as stacks
	stack := []string{}
	stack = append(stack, "a")
	stack = append(stack, "b")
	stack = stack[:len(stack)-1]
	fmt.Println(stack)
}

// arrays demonstrates various array operations and concepts
func arrays() {
	// declaration and initialization
	var a [3]int
	fmt.Println(a)

	b := [3]int{1, 2, 3}
	fmt.Println(b)

	// accessing
	fmt.Println(b[1])

	// modifying
	b[1] = 4
	fmt.Println(b)

	// iterating
	for i, v := range b {
		fmt.Println(i, v)
	}
	for _, v := range b {
		fmt.Println(v)
	}

	// length
	fmt.Println(len(b))

	// multidimensional arrays
	c := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(c)

	// array literals
	d := [...]int{5, 6, 7}
	fmt.Println(d)

	// copying
	e := d
	fmt.Println(e)

	// comparing
	fmt.Println(a == e)
	fmt.Println(b == e)
}
