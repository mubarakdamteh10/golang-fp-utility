package reflection

import (
	"errors"
	"reflect"
	"strings"
)

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

// Case attempts to convert an interface{} to a specific type and returns a pointer to the result.
func Case[T any](source interface{}) (*T, error) {
	converted, ok := source.(T)
	if !ok {
		return nil, errors.New("type assertion failed")
	}
	return &converted, nil
}
