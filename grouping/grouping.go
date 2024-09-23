package grouping

import (
	"fmt"
	"reflect"

	reflection "github.com/lumiluminousai/golang-fp-utility/reflection"
)

// GroupBy groups elements of a list by a specified field name.
func GroupBy[K comparable, V any](slice []V, fieldName string) (map[K][]V, error) {
	result := make(map[K][]V)
	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Kind() != reflect.Slice {
		return nil, fmt.Errorf("groupBy: provided argument is not a slice")
	}
	for i := 0; i < sliceValue.Len(); i++ {
		element := sliceValue.Index(i)
		fieldValue := reflection.GetField(element, fieldName)
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
		fieldValue := reflection.GetField(element, fieldName)
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
