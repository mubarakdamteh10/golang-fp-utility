golang-fp-utility

Welcome to golang-fp-utility! This Go package is crafted to help developers embrace functional programming paradigms in Go. With a suite of utility functions, this package encourages immutable coding practices, helping you move away from traditional mutable loops like for loops. Together, let’s make Go code more elegant and maintainable!

Features

	•	Immutability First: All functions return new collections rather than modifying the originals, promoting a functional programming style.
	•	Comprehensive List and Map Operations: Includes functions like Map, Filter, Reduce, GroupBy, and more, enabling functional processing of collections.
	•	Error Handling: Many functions come with built-in error handling, making it easier to write robust code.
	•	Eliminate Side-Effects: Functions like ForEachWithError and MapReturnWithError help ensure your code handles errors gracefully.
	•	Type-Safe Generics: With Go’s generics support, these utilities work seamlessly across various data types.

Why Functional Programming?

Functional programming (FP) emphasizes the use of functions and immutability, avoiding shared state and side effects. This leads to more predictable and maintainable code. By adopting FP practices in Go, you can write clearer, more concise, and bug-resistant code.

Utility Functions

General Purpose

	•	IfThen[T any](condition bool, ifTrue, ifFalse T) T: Conditional inline operation, similar to the ternary operator in other languages.

List Operations

	•	Filter[T any](source []T, filterFunc func(item T) bool) []T: Filters a list based on a provided function.
	•	Map[T1 any, T2 any](source []T1, transform func(item T1) T2) []T2: Applies a transformation function to each item in the list and returns a new list.
	•	Reduce[T any](source []T, reduceFunc func(acc T, item T) T, initialValue T) T: Reduces a list to a single value using an accumulator function.
	•	Distinct[T comparable](slice []T) []T: Returns a slice containing only unique elements.
	•	Sort[T any](list []T, less func(i, j int) bool) []T: Sorts a list using a custom less function.

Map Operations

	•	MapToHashMap[T1 any, T2 any, K comparable](source []T1, mappingFunc func(item T1) (K, T2)) map[K]T2: Converts a list to a hashmap using a transformation function.
	•	FilterMap[K comparable, V any](source map[K]V, filteringFunc func(key K, value V) bool) map[K]V: Filters a hashmap based on a provided function.
	•	MapHashMapToHashMap[K comparable, V1 any, V2 any](source map[K]V1, mappingFunc func(key K, value V1) V2) map[K]V2: Applies a transformation function to a hashmap and returns a new hashmap.

Grouping and Reflection

	•	GroupBy[K comparable, V any](slice []V, fieldName string) (map[K][]V, error): Groups elements of a list by a specified field name.
	•	GetField(element reflect.Value, fieldName string) reflect.Value: Retrieves the value of a nested field by name.

Utility Functions

	•	Sum[T Summable](list []T) T: Returns the sum of elements in a slice of summable types (e.g., integers, floats).
	•	Case[T any](source interface{}) (*T, error): Attempts to convert an interface{} to a specific type, returning a pointer.

Installation

To install the package, run:

	go get github.com/elytralover/golang-fp-utility

Usage

Here’s a quick example of how you can start using these utilities:

	func main() {
		numbers := []int{1, 2, 3, 4, 5}

		// Example using Map to square numbers
		squares := utility.Map(numbers, func(n int) int {
			return n * n
		})

		// Example using Filter to filter even numbers
		evenSquares := utility.Filter(squares, func(n int) bool {
			return n%2 == 0
		})

		fmt.Println(evenSquares) // Output: [4 16]
	}

Contributing

Contributions are welcome! If you have any ideas, suggestions, or improvements, feel free to open an issue or submit a pull request.

License

This project is licensed under the terms of the GNU Lesser General Public License (LGPL) Version 3. You can view the full license text in the LICENSE.txt file included in this repository.

Summary of the License

	•	You are free to use, modify, and distribute this software.
	•	If you distribute modified versions, you must include the source code and keep the same license.
	•	This license does not require you to release your own proprietary code when using this library.

For more details, please refer to the LGPL Version 3.

This format ensures that your README is clear, concise, and easy to read while maintaining a professional presentation.