package DataTypesAndOperators

import (
	"fmt"
	"reflect"
)

// func main() {
// 	arithmeticOperators()
// 	//collections()
// 	//constants()
// 	//primitiveDataTypes()
// }

func arithmeticOperators() {
	add := 1 + 2
	fmt.Println(add)

	subtract := 10 - 5
	fmt.Println(subtract)

	remainder := 5 % 2
	fmt.Println(remainder)

	fmt.Println("---------------")

	inc := 1
	inc++
	fmt.Println(inc)
	inc++
	fmt.Println(inc)
	inc += 5
	fmt.Println(inc)
	inc *= 10
	fmt.Println(inc)
}

func collections() {
	myMap := make(map[int]string)
	fmt.Println(myMap)
	myMap[42] = "Foo"
	myMap[12] = "Bar"

	fmt.Println(myMap)
	fmt.Println(myMap[99])

	empty := myMap[99]

	fmt.Println(reflect.TypeOf(empty))

	myFloatSlice := make([]float32, 100)
	myFloatSlice[0] = 12.
	myFloatSlice[1] = 15.
	myFloatSlice[2] = 18.

	//myFloatSlice := []float32{14., 15., 19.}

	fmt.Println(myFloatSlice)
	fmt.Println(len(myFloatSlice))

	myArray := [...]int{42, 27, 99}

	mySlice := myArray[:]
	fmt.Println(myArray)

	mySlice = append(mySlice, 100)

	fmt.Println(myArray)
	fmt.Println(mySlice)
}

func constants() {
	const (
		first = 1 << (10 * iota)
		second
		third
	)

	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)
}

func primitiveDataTypes() {
	var myInt int
	myInt = 42

	fmt.Println(myInt)

	var myFloat32 float32 = 42.
	fmt.Println(myFloat32)

	myString := "Hello Go"
	fmt.Println(myString)

	myComplex := complex(2, 3)
	fmt.Println(myComplex)
	fmt.Println(real(myComplex))
	fmt.Println(imag(myComplex))
}
