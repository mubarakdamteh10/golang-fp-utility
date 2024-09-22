package utility

// Package utility provides utility functions for functional programming in Go.
//
// This file is part of golang-fp-utility.
//
// golang-fp-utility is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License
// as published by the Free Software Foundation, either version 3
// of the License, or (at your option) any later version.
//
// golang-fp-utility is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with golang-fp-utility. If not, see <https://www.gnu.org/licenses/lgpl-3.0.txt>.

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/pkg/errors"
)

// IfThen executes an if-else operation in a single line.
// Examples:
//   - IfThen(len("ABC") < 4, "less than 4", "more than 4") returns "less than 4".
//   - IfThen(len("ABD") < 4, 0, 1) returns 0.
//   - IfThen(len("ABCD") < 4, "0", "1") returns "1".
func IfThen[T any](condition bool, ifTrue, ifFalse T) T {
	if condition {
		return ifTrue
	}
	return ifFalse
}

// Filter returns a filtered list based on the provided function.
func Filter[T any](source []T, filterFunc func(item T) bool) []T {
	result := []T{}
	for _, item := range source {
		if filterFunc(item) {
			result = append(result, item)
		}
	}
	return result
}

// ForEach executes a function for each item in the list.
func ForEach[T any](source []T, action func(item T)) {
	for _, item := range source {
		action(item)
	}
}

// ForEachWithError executes a function for each item and handles errors.
func ForEachWithError[T any](source []T, action func(item T) error) error {
	for _, item := range source {
		if err := action(item); err != nil {
			return err
		}
	}
	return nil
}

// Map applies a transformation function to each item in the list and returns a new list.
func Map[T1 any, T2 any](source []T1, transform func(item T1) T2) []T2 {
	result := []T2{}
	for _, item := range source {
		result = append(result, transform(item))
	}
	return result
}

// MapReturnWithError applies a transformation function to each item and handles errors.
func MapReturnWithError[T1 any, T2 any](source []T1, mappingFunc func(item T1) (T2, error)) ([]T2, error) {
	result := []T2{}

	for idx, item := range source {
		res, err := mappingFunc(item)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error mapping at index:'%v', error", idx))
		}
		result = append(result, res)
	}
	return result, nil
}

// MapToHashMap converts a list to a hashmap using a transformation function.
func MapToHashMap[T1 any, T2 any, K comparable](source []T1, mappingFunc func(item T1) (K, T2)) map[K]T2 {
	result := make(map[K]T2)
	for _, item := range source {
		key, value := mappingFunc(item)
		result[key] = value
	}
	return result
}

// MapToHashMapReturnWithError converts a list to a hashmap with error handling.
func MapToHashMapReturnWithError[T1 any, T2 any, K comparable](source []T1, mappingFunc func(item T1) (K, T2, error)) (map[K]T2, error) {
	result := make(map[K]T2)
	for idx, item := range source {
		key, value, err := mappingFunc(item)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error mapping at index:'%v', error", idx))
		}
		result[key] = value
	}
	return result, nil
}

// MapHashMapToHashMap applies a transformation function to a hashmap and returns a new hashmap.
func MapHashMapToHashMap[K comparable, V1 any, V2 any](source map[K]V1, mappingFunc func(key K, value V1) V2) map[K]V2 {
	result := make(map[K]V2)
	for key, value := range source {
		result[key] = mappingFunc(key, value)
	}
	return result
}

// MapHashMapToHashMapReturnWithError applies a transformation function to a hashmap and handles errors.
func MapHashMapToHashMapReturnWithError[K comparable, V1 any, V2 any](source map[K]V1, mappingFunc func(key K, value V1) (V2, error)) (map[K]V2, error) {
	result := make(map[K]V2)
	for key, value := range source {
		res, err := mappingFunc(key, value)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error mapping at key:'%v', error", key))
		}
		result[key] = res
	}
	return result, nil
}

// MapHashMapToList applies a transformation function to a hashmap and returns a list.
func MapHashMapToList[K comparable, V1 any, V2 any](source map[K]V1, mappingFunc func(key K, value V1) V2) []V2 {
	keys := []K{}

	for key := range source {
		keys = append(keys, key)
	}
	sortedKeys := Sort(keys, func(i, j int) bool { return fmt.Sprintf("%v", keys[i]) < fmt.Sprintf("%v", keys[j]) })
	return Map(sortedKeys, func(key K) V2 { return mappingFunc(key, source[key]) })
}

// MapHashMapToListReturnWithError applies a transformation function to a hashmap, returning a list with error handling.
func MapHashMapToListReturnWithError[K comparable, V1 any, V2 any](source map[K]V1, mappingFunc func(key K, value V1) (V2, error)) ([]V2, error) {
	keys := []K{}

	for key := range source {
		keys = append(keys, key)
	}
	sortedKeys := Sort(keys, func(i, j int) bool { return fmt.Sprintf("%v", keys[i]) < fmt.Sprintf("%v", keys[j]) })
	result := []V2{}
	for _, key := range sortedKeys {
		res, err := mappingFunc(key, source[key])
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error mapping at key:'%v', error", key))
		}
		result = append(result, res)
	}
	return result, nil
}

// FilterMap filters a hashmap based on a provided function.
func FilterMap[K comparable, V any](source map[K]V, filteringFunc func(key K, value V) bool) map[K]V {
	result := make(map[K]V)
	for key, value := range source {
		if filteringFunc(key, value) {
			result[key] = value
		}
	}
	return result
}

// FlatMap flattens a list of lists into a single list.
func FlatMap[T1 any](source [][]T1) []T1 {
	result := []T1{}
	for _, item := range source {
		result = append(result, item...)
	}
	return result
}

// Reduce reduces a list to a single value using the provided function.
func Reduce[T any](source []T, reduceFunc func(acc T, item T) T, initialValue T) T {
	acc := initialValue
	for _, item := range source {
		acc = reduceFunc(acc, item)
	}
	return acc
}

// CloneMap creates a shallow copy of the given map.
func CloneMap[K comparable, V any](source map[K]V) map[K]V {
	clone := make(map[K]V, len(source))
	for key, value := range source {
		clone[key] = value
	}
	return clone
}

// CloneList creates a shallow copy of the given list.
func CloneList[T any](source []T) []T {
	clone := make([]T, len(source))
	copy(clone, source)
	return clone
}

// GroupBy groups elements of a list by a specified field name.
func GroupBy[K comparable, V any](slice []V, fieldName string) (map[K][]V, error) {
	result := make(map[K][]V)
	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Kind() != reflect.Slice {
		return nil, fmt.Errorf("groupBy: provided argument is not a slice")
	}
	for i := 0; i < sliceValue.Len(); i++ {
		element := sliceValue.Index(i)
		fieldValue := GetField(element, fieldName)
		if !fieldValue.IsValid() {
			return nil, fmt.Errorf("groupBy: field %s does not exist", fieldName)
		}
		key := fieldValue.Interface().(K)
		result[key] = append(result[key], element.Interface().(V))
	}
	return result, nil
}

// GroupBy1By1 groups elements of a list by a specified field name, ensuring uniqueness.
func GroupBy1By1[K comparable, V any](slice []V, fieldName string) (map[K]V, error) {
	grouped := make(map[K][]V)
	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Kind() != reflect.Slice {
		return nil, fmt.Errorf("groupBy: provided argument is not a slice")
	}
	for i := 0; i < sliceValue.Len(); i++ {
		element := sliceValue.Index(i)
		fieldValue := GetField(element, fieldName)
		if !fieldValue.IsValid() {
			return nil, fmt.Errorf("groupBy: field %s does not exist", fieldName)
		}
		key := fieldValue.Interface().(K)
		grouped[key] = append(grouped[key], element.Interface().(V))
	}
	uniqueResult := make(map[K]V)
	for key, value := range grouped {
		if len(value) > 1 {
			return nil, fmt.Errorf("groupBy: field %s is not unique", fieldName)
		}
		uniqueResult[key] = value[0]
	}
	return uniqueResult, nil
}

// GetField retrieves the value of a nested field by name.
func GetField(element reflect.Value, fieldName string) reflect.Value {
	names := strings.Split(fieldName, ".")
	for _, name := range names {
		if element.Kind() == reflect.Ptr {
			element = element.Elem()
		}
		if element.Kind() == reflect.Slice {
			var subElements []reflect.Value
			for i := 0; i < element.Len(); i++ {
				subElem := GetField(element.Index(i), name)
				if subElem.IsValid() {
					subElements = append(subElements, subElem)
				}
			}
			// Convert the slice of reflect.Value to a slice of interfaces.
			result := make([]interface{}, len(subElements))
			for i, v := range subElements {
				result[i] = v.Interface()
			}
			return reflect.ValueOf(result)
		}
		element = element.FieldByName(name)
	}
	return element
}

// SliceToHashMap converts a slice to a map with boolean values indicating presence.
func SliceToHashMap[T comparable](list []T) map[T]bool {
	result := make(map[T]bool)
	for _, item := range list {
		result[item] = true
	}
	return result
}

// Distinct returns a slice containing only unique elements.
func Distinct[T comparable](slice []T) []T {
	seen := make(map[T]bool)
	unique := []T{}
	for _, item := range slice {
		if !seen[item] {
			seen[item] = true
			unique = append(unique, item)
		}
	}
	return unique
}

// DistinctFunc returns a slice containing unique elements using a custom comparison function.
func DistinctFunc[T comparable](slice []T, compareFunc func(a, b T) bool) []T {
	seen := make(map[T]bool)
	unique := []T{}
	for _, item := range slice {
		if !seen[item] {
			seen[item] = true
			unique = append(unique, item)
		}
	}
	return unique
}

// Sort sorts a slice using a custom less function.
func Sort[T any](list []T, less func(i, j int) bool) []T {
	sort.Slice(list, less)
	return list
}

// Summable includes all types that can be summed, such as integers and floats.
type Summable interface {
	int | int32 | int64 | float32 | float64
}

// Sum returns the sum of elements in a slice of summable types.
func Sum[T Summable](list []T) T {
	var total T
	for _, v := range list {
		total += v
	}
	return total
}

// Case attempts to convert an interface{} to a specific type and returns a pointer to the result.
func Case[T any](source interface{}) (*T, error) {
	converted, ok := source.(T)
	if !ok {
		return nil, errors.New("type assertion failed")
	}
	return &converted, nil
}
