package maps

import (
	"fmt"

	"github.com/pkg/errors"

	collection "github.com/lumiluminousai/golang-fp-utility/collection"
)

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
	sortedKeys := collection.Sort(keys, func(i, j int) bool { return fmt.Sprintf("%v", keys[i]) < fmt.Sprintf("%v", keys[j]) })
	return collection.Map(sortedKeys, func(key K) V2 { return mappingFunc(key, source[key]) })
}

// MapHashMapToListReturnWithError applies a transformation function to a hashmap, returning a list with error handling.
func MapHashMapToListReturnWithError[K comparable, V1 any, V2 any](source map[K]V1, mappingFunc func(key K, value V1) (V2, error)) ([]V2, error) {
	keys := []K{}

	for key := range source {
		keys = append(keys, key)
	}
	sortedKeys := collection.Sort(keys, func(i, j int) bool { return fmt.Sprintf("%v", keys[i]) < fmt.Sprintf("%v", keys[j]) })
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

// SliceToHashMap converts a slice to a map with boolean values indicating presence.
func SliceToHashMap[T comparable](list []T) map[T]bool {
	result := make(map[T]bool)
	for _, item := range list {
		result[item] = true
	}
	return result
}
