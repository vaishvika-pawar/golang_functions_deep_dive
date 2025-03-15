package main

import "fmt"

type transformFn func(int) int

// type anotherFn func(int, []string, map[string][]int) ([]int, string)

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 1, 2}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println("doubled=>", doubled)
	fmt.Println("tripled=>", tripled)

	transformerFn1 := getTransformerFunction(&numbers)
	transformerFn2 := getTransformerFunction(&moreNumbers)

	transformedNumbers := transformNumbers(&numbers, transformerFn1)
	moreTransformedNumbers := transformNumbers(&numbers, transformerFn2)

	fmt.Println("transformedNumbers=>", transformedNumbers)
	fmt.Println("moreTransformedNumbers=>", moreTransformedNumbers)

	//Anonymous functions: because it doesn;t have a name
	//It is not a type of a function but a function itself
	//You can take params as we want
	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	fmt.Println("transformed=>", transformed)

	//Understanding closures
	double1 := createTransformer(2)
	triple1 := createTransformer(3)

	doubled1 := transformNumbers(&numbers, double1)
	tripled1 := transformNumbers(&numbers, triple1)

	fmt.Println("doubled1=>", doubled1)
	fmt.Println("tripled1=>", tripled1)

	fact := factorial(3)

	fmt.Println("factorial=>", fact)

	sum := sumup([]int{10, 15, 40, -5}...)
	fmt.Println("sum=>", sum)

	anotherSum := sumup(numbers...)
	fmt.Println("anotherSum=>", anotherSum)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))

	}
	return dNumbers
}

func getTransformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

// Closures
// We can use it as a fcator contructor
// We can create different transformer functions
func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

// Recursion
func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}

// Veriadic functions
func sumup(numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val // sum = sum + val
	}
	return sum
}

// Splitting slices into parameter values
