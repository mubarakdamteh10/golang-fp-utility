package reflection

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetField(t *testing.T) {
	t.Run("Success_get_primitive_typeField_in_Layer_1", func(t *testing.T) {
		type Layer2 struct {
			Field1 string
			Field2 int
		}
		type MyStruct struct {
			Name   string
			Age    int
			Layer2 Layer2
		}

		data := MyStruct{
			Name: "John",
			Age:  30,
			Layer2: Layer2{
				Field1: "Value1",
				Field2: 42,
			},
		}

		tests := []struct {
			fieldName string
			expected  interface{}
		}{
			{"Name", "John"},
			{"Age", 30},
		}

		for _, test := range tests {
			t.Run(test.fieldName, func(t *testing.T) {
				actual := GetField(reflect.ValueOf(data), test.fieldName).Interface()
				if actual != test.expected {
					t.Errorf("Expected %v but got %v", test.expected, actual)
				}
			})
		}
	})

	t.Run("Success_get_primitive_typeField_in_layer_2", func(t *testing.T) {
		type Layer2 struct {
			Field1 string
			Field2 int
		}
		type MyStruct struct {
			Name   string
			Age    int
			Layer2 Layer2
		}

		data := MyStruct{
			Name: "John",
			Age:  30,
			Layer2: Layer2{
				Field1: "Value1",
				Field2: 42,
			},
		}

		tests := []struct {
			fieldName string
			expected  interface{}
		}{
			{"Layer2.Field1", "Value1"},
			{"Layer2.Field2", 42},
		}

		for _, test := range tests {
			t.Run(test.fieldName, func(t *testing.T) {
				actual := GetField(reflect.ValueOf(data), test.fieldName).Interface()
				if actual != test.expected {
					t.Errorf("Expected %v but got %v", test.expected, actual)
				}
			})
		}
	})

	t.Run("Success_get_primitive_typeField_in_layer_3", func(t *testing.T) {
		type Layer3 struct {
			Field3 string
			Field4 bool
		}

		type Layer2 struct {
			Field1 string
			Field2 int
			Layer3 Layer3
		}
		type MyStruct struct {
			Name   string
			Age    int
			Layer2 Layer2
		}

		data := MyStruct{
			Name: "John",
			Age:  30,
			Layer2: Layer2{
				Field1: "Value1",
				Field2: 42,
				Layer3: Layer3{
					Field3: "Value3",
					Field4: true,
				},
			},
		}

		tests := []struct {
			fieldName string
			expected  interface{}
		}{
			{"Layer2.Layer3.Field3", "Value3"},
			{"Layer2.Layer3.Field4", true},
		}

		for _, test := range tests {
			t.Run(test.fieldName, func(t *testing.T) {
				actual := GetField(reflect.ValueOf(data), test.fieldName).Interface()
				if actual != test.expected {
					t.Errorf("Expected %v but got %v", test.expected, actual)
				}
			})
		}
	})
}

func Test_CaseObject(t *testing.T) {

	type TempStruct struct {
		Name  string
		Value int
	}

	t.Run("CaseObject", func(t *testing.T) {

		value1 := TempStruct{
			Name:  "value1",
			Value: 1,
		}

		interface1 := interface{}(value1)
		casedObject1, err := Case[TempStruct](interface1)
		assert.Nil(t, err)

		expected := TempStruct{
			Name:  "value1",
			Value: 1,
		}

		assert.EqualValues(t, &expected, casedObject1)
	})

	t.Run("CaseObject_error", func(t *testing.T) {

		value1 := TempStruct{
			Name:  "value1",
			Value: 1,
		}

		interface1 := interface{}(value1)
		_, err := Case[int](interface1)
		assert.NotNil(t, err)
		assert.Equal(t, "type assertion failed", err.Error())
	})

	t.Run("CaseObject_nil", func(t *testing.T) {

		interface1 := interface{}(nil)
		_, err := Case[int](interface1)
		assert.NotNil(t, err)
		assert.Equal(t, "type assertion failed", err.Error())
	})

	t.Run("CaseWrong_object", func(t *testing.T) {

		value1 := TempStruct{
			Name:  "value1",
			Value: 1,
		}

		type TempStruct2 struct {
			Name2  string
			Value2 int
		}

		interface1 := interface{}(value1)
		casedObject1, err := Case[TempStruct2](interface1)
		assert.NotNil(t, err)
		assert.Equal(t, "type assertion failed", err.Error())
		assert.Nil(t, casedObject1)
	})

}
