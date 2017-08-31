package functionsAndParameters

import (
	"fmt"
)

// func main() {
// 	message := "Hello"
// 	passValue(message)
// 	fmt.Println("-----")

// 	passReference(&message)
// 	fmt.Println(message)
// 	fmt.Println("-----")

// 	variadic("Hello", "Go", "from", "Pluralsight")
// 	fmt.Println("-----")

// 	result := add(1, 3, 5, 9)
// 	fmt.Println(result)
// 	fmt.Println("-----")

// 	numTerms, sum := multiReturnAdd(1, 3, 5, 9)
// 	fmt.Println("Added:", numTerms, "terms for a total of", sum)
// 	fmt.Println("-----")

// 	numTerms2, sum2 := namedMultiReturnAdd(1, 3, 5, 9)
// 	fmt.Println("Added:", numTerms2, "terms for a total of", sum2)
// 	fmt.Println("-----")

// 	addFunc := func(terms ...int) (numTerms int, sum int) {
// 		for _, term := range terms {
// 			sum += term
// 		}

// 		numTerms = len(terms)

// 		return
// 	}

// 	numTerms3, sum3 := addFunc(1, 3, 5, 9)
// 	fmt.Println("Added:", numTerms3, "terms for a total of", sum3)
// 	fmt.Println("-----")
// }

func namedMultiReturnAdd(terms ...int) (numTerms int, sum int) {
	for _, term := range terms {
		sum += term
	}

	numTerms = len(terms)

	return
}

func multiReturnAdd(terms ...int) (int, int) {
	result := 0
	for _, term := range terms {
		result += term
	}

	return len(terms), result
}

func add(terms ...int) int {
	result := 0
	for _, term := range terms {
		result += term
	}

	return result
}

func variadic(messages ...string) {
	for _, message := range messages {
		fmt.Println(message)
	}
}

func passReference(message *string) {
	fmt.Println(*message)
	*message = "Hello Go"
}

func passValue(message string) {
	fmt.Println(message)
}
