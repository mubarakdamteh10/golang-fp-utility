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
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfThen(t *testing.T) {
	result := IfThen(true, 1, 2)
	assert.Equal(t, 1, result)

	result = IfThen(false, 1, 2)
	assert.Equal(t, 2, result)

	type TempStruct struct {
		Name  string
		Value int
	}
	value1 := TempStruct{
		Name:  "value1",
		Value: 1,
	}
	value2 := TempStruct{
		Name:  "value2",
		Value: 2,
	}

	resultTempStruct := IfThen(true, value1, value2)
	assert.Equal(t, value1, resultTempStruct)

	resultTempStruct = IfThen(false, value1, value2)
	assert.Equal(t, value2, resultTempStruct)

	resultTempStruct = IfThen(value1.Value < value2.Value, value2, value1)
	assert.Equal(t, value2, resultTempStruct)

	resultInt := IfThen(false, 1, 2)
	assert.Equal(t, 2, resultInt)

	resultString := IfThen(true, "1", "2")
	assert.Equal(t, "1", resultString)

	resultString = IfThen(len("ABC") < 4, "less than length 4.", "more than length 4.")
	assert.Equal(t, "less than length 4.", resultString)

	resultString = IfThen(len("ABD") < 4, "0", "1")
	assert.Equal(t, "0", resultString)

	resultString = IfThen(len("ABCD") < 4, "0", "1")
	assert.Equal(t, "1", resultString)

	tempInfo1 := TempStruct{
		Name:  IfThen(len("ABCD") < 4, "1", "0"),
		Value: 2,
	}
	expected := TempStruct{
		Name:  "0",
		Value: 2,
	}
	assert.EqualValues(t, expected, tempInfo1)

	tempInfo1 = TempStruct{
		Name:  IfThen(len("ABCD") == 4, "ABCD", ""),
		Value: IfThen(len("ABCD") > 4, 4, 0),
	}
	expected = TempStruct{
		Name:  "ABCD",
		Value: 0,
	}
	assert.EqualValues(t, expected, tempInfo1)
}

func TestFilter(t *testing.T) {
	t.Run("filter > 3", func(t *testing.T) {

		source := []int{1, 2, 3, 4, 5}
		mappingFunc := func(data int) bool {
			return data > 3 // Convert each integer to string with prefix
		}

		result := Filter(source, mappingFunc)

		expected := []int{4, 5}
		assert.Equal(t, expected, result)
	})

	t.Run("filter empty list", func(t *testing.T) {
		source := []int{}
		mappingFunc := func(data int) bool {
			return data > 3 // Convert each integer to string with prefix
		}

		result := Filter(source, mappingFunc)

		expected := []int{}
		assert.Equal(t, expected, result)
	})
}

func TestForEach(t *testing.T) {
	t.Run("print integers", func(t *testing.T) {

		source := []int{1, 2, 3, 4, 5}
		forEachFunc := func(item int) {
			fmt.Println(item)
		}

		ForEach(source, forEachFunc)
	})

	t.Run("change value each item", func(t *testing.T) {

		source := []int{1, 2, 3, 4, 5}
		forEachFunc := func(item int) {
			item = item * 2
		}

		ForEach(source, forEachFunc)

		expected := []int{1, 2, 3, 4, 5}
		assert.Equal(t, expected, source)
	})

	t.Run("change value object", func(t *testing.T) {

		type TempStruct struct {
			Name  string
			Value int
		}
		value1 := TempStruct{
			Name:  "value1",
			Value: 1,
		}
		value2 := TempStruct{
			Name:  "value2",
			Value: 2,
		}

		source := []TempStruct{value1, value2}
		forEachFunc := func(item TempStruct) {
			item.Value = item.Value * 2
		}

		ForEach(source, forEachFunc)

		expected := []TempStruct{value1, value2}
		assert.Equal(t, expected, source)
	})

	t.Run("change value object pointer", func(t *testing.T) {

		type TempStruct struct {
			Name  string
			Value int
		}
		value1 := &TempStruct{
			Name:  "value1",
			Value: 1,
		}
		value2 := &TempStruct{
			Name:  "value2",
			Value: 2,
		}

		source := []*TempStruct{value1, value2}
		forEachFunc := func(item *TempStruct) {
			item.Value = item.Value * 2
		}

		ForEach(source, forEachFunc)

		expected := []*TempStruct{value1, value2}
		assert.Equal(t, expected, source)
	})
}

func TestForEachWithError(t *testing.T) {
	t.Run("print integers", func(t *testing.T) {

		source := []int{1, 2, 3, 4, 5}
		forEachFunc := func(item int) error {
			if item == 3 {
				return errors.New("error")
			}
			fmt.Println(item)
			return nil
		}

		err := ForEachWithError(source, forEachFunc)
		assert.Error(t, err)
	})
}

func TestMap(t *testing.T) {

	t.Run("test use sum() function", func(t *testing.T) {
		// Example list of doubles
		numbers := []float64{1.5, 2.0, 3.5, 4.0}

		// Use utility.Map to double each number in the list
		doubledNumbers := Map(numbers, func(item float64) float64 {
			return item * 2
		})

		// Use utility.Sum to get the summation of the doubled numbers
		sum := Sum(doubledNumbers)

		// Assert the expected values
		assert.Equal(t, []float64{3.0, 4.0, 7.0, 8.0}, doubledNumbers, "Doubled numbers should match the expected list")
		assert.Equal(t, 22.0, sum, "Summation of doubled numbers should be 22.0")
	})

	t.Run("map with nil list", func(t *testing.T) {
		source := []int(nil)

		mappingFunc := func(item int) string {
			return fmt.Sprintf("string_%v", item) // Convert each integer to string with prefix
		}

		result := Map(source, mappingFunc)

		expected := []string{}
		assert.Equal(t, expected, result)
	})

	t.Run("map integers to strings", func(t *testing.T) {

		source := []int{1, 2, 3, 4, 5}
		mappingFunc := func(item int) string {
			return fmt.Sprintf("string_%v", item) // Convert each integer to string with prefix
		}

		result := Map(source, mappingFunc)

		expected := []string{"string_1", "string_2", "string_3", "string_4", "string_5"}
		assert.Equal(t, expected, result)
	})

	t.Run("map empty list", func(t *testing.T) {
		source := []int{}
		mappingFunc := func(item int) string {
			return fmt.Sprintf("string_%v", item) // Convert each integer to string with prefix
		}

		result := Map(source, mappingFunc)

		expected := []string{}
		assert.Equal(t, expected, result)
	})
}

func TestFilterMap(t *testing.T) {
	tests := []struct {
		name          string
		source        map[int]string // Example uses int keys and string values for demonstration
		filteringFunc func(int, string) bool
		want          map[int]string
	}{
		{
			name:   "filter out odd keys",
			source: map[int]string{1: "a", 2: "b", 3: "c", 4: "d"},
			filteringFunc: func(key int, value string) bool {
				return key%2 == 0 // Keep if key is even
			},
			want: map[int]string{2: "b", 4: "d"},
		},
		{
			name:   "filter out values with length > 1",
			source: map[int]string{1: "a", 2: "bb", 3: "ccc", 4: "dddd"},
			filteringFunc: func(key int, value string) bool {
				return len(value) <= 1 // Keep if value's length is 1 or less
			},
			want: map[int]string{1: "a"},
		},
		{
			name:          "empty map",
			source:        map[int]string{},
			filteringFunc: func(key int, value string) bool { return true },
			want:          map[int]string{},
		},
		{
			name:   "filter everything",
			source: map[int]string{1: "a", 2: "b"},
			filteringFunc: func(key int, value string) bool {
				return false // Filter out everything
			},
			want: map[int]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterMap(tt.source, tt.filteringFunc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestCloneMap(t *testing.T) {
	tests := []struct {
		name   string
		source map[string]int // This test uses string keys and int values for simplicity
		want   map[string]int
	}{
		{
			name:   "non-empty map",
			source: map[string]int{"a": 1, "b": 2},
			want:   map[string]int{"a": 1, "b": 2},
		},
		{
			name:   "empty map",
			source: map[string]int{},
			want:   map[string]int{},
		},
		{
			name:   "nil map",
			source: nil,
			want:   map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CloneMap(tt.source)

			assert.Equal(t, tt.want, result)
		})
	}
}

func TestCloneStringList(t *testing.T) {
	tests := []struct {
		name   string
		source []string
		want   []string
	}{
		{
			name:   "empty list",
			source: []string{},
			want:   []string{},
		},
		{
			name:   "single element",
			source: []string{"element"},
			want:   []string{"element"},
		},
		{
			name:   "multiple elements",
			source: []string{"hello", "world"},
			want:   []string{"hello", "world"},
		},
		{
			name:   "nil list",
			source: nil,
			want:   []string{},
		},
		{
			name:   "empty list",
			source: []string{},
			want:   []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CloneList(tt.source)

			assert.Equal(t, tt.want, result)
		})
	}
}

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

func TestGroupBy(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	t.Run("Success_groupBy_age", func(t *testing.T) {
		people := []Person{
			{Name: "Alice", Age: 30},
			{Name: "Bob", Age: 30},
			{Name: "Charlie", Age: 25},
		}

		fieldName := "Age"

		result, err := GroupBy[int](people, fieldName)
		assert.NoError(t, err)

		expected := map[int][]Person{
			30: {people[0], people[1]},
			25: {people[2]},
		}

		assert.Equal(t, expected, result)
	})
	t.Run("Success_groupBy_name", func(t *testing.T) {
		people := []Person{
			{Name: "Alice", Age: 30},
			{Name: "Bob", Age: 30},
			{Name: "Charlie", Age: 25},
			{Name: "Charlie", Age: 25},
		}

		fieldName := "Name"

		result, err := GroupBy[string](people, fieldName)
		assert.NoError(t, err)

		expected := map[string][]Person{
			"Alice":   {people[0]},
			"Bob":     {people[1]},
			"Charlie": {people[2], people[3]},
		}

		assert.Equal(t, expected, result)
	})

	t.Run("Success_groupBy_layer2_field", func(t *testing.T) {
		type Layer2 struct {
			Field1 string
			Field2 int
		}

		type Layer1 struct {
			Name   string
			Age    int
			Layer2 Layer2
		}

		people := []Layer1{
			{
				Name: "Alice",
				Age:  30,
				Layer2: Layer2{
					Field1: "Value1",
					Field2: 42,
				},
			},
			{
				Name: "Bob",
				Age:  30,
				Layer2: Layer2{
					Field1: "Value2",
					Field2: 43,
				},
			},
			{
				Name: "Charlie",
				Age:  25,
				Layer2: Layer2{
					Field1: "Value3",
					Field2: 44,
				},
			},
			{
				Name: "x",
				Age:  28,
				Layer2: Layer2{
					Field1: "Value1",
					Field2: 45,
				},
			},
		}

		fieldName := "Layer2.Field1"

		result, err := GroupBy[string](people, fieldName)
		assert.NoError(t, err)

		expected := map[string][]Layer1{
			"Value1": {people[0], people[3]},
			"Value2": {people[1]},
			"Value3": {people[2]},
		}

		assert.Equal(t, expected, result)
	})

	t.Run("Success_groupBy_layer3_field", func(t *testing.T) {

		type Layer3 struct {
			Field3 string
			Field4 bool
		}

		type Layer2 struct {
			Field1 string
			Field2 int
			Layer3 Layer3
		}

		type Layer1 struct {
			Name   string
			Age    int
			Layer2 Layer2
		}

		people := []Layer1{
			{
				Name: "Alice",
				Age:  30,
				Layer2: Layer2{
					Field1: "Value1",
					Field2: 42,
					Layer3: Layer3{
						Field3: "layer3-1",
						Field4: true,
					},
				},
			},
			{
				Name: "Bob",
				Age:  30,
				Layer2: Layer2{
					Field1: "Value2",
					Field2: 43,
					Layer3: Layer3{
						Field3: "layer3-2",
						Field4: false,
					},
				},
			},
			{
				Name: "Charlie",
				Age:  25,
				Layer2: Layer2{
					Field1: "Value3",
					Field2: 44,
					Layer3: Layer3{
						Field3: "layer3-3",
						Field4: true,
					},
				},
			},
			{
				Name: "x",
				Age:  28,
				Layer2: Layer2{
					Field1: "Value1",
					Field2: 45,
					Layer3: Layer3{
						Field3: "layer3-2",
						Field4: true,
					},
				},
			},
		}

		fieldName := "Layer2.Layer3.Field3"

		result, err := GroupBy[string](people, fieldName)
		assert.NoError(t, err)

		expected := map[string][]Layer1{
			"layer3-1": {people[0]},
			"layer3-2": {people[1], people[3]},
			"layer3-3": {people[2]},
		}

		assert.Equal(t, expected, result)
	})

	t.Run("Error_invalid_field_name", func(t *testing.T) {
		people := []Person{
			{Name: "Alice", Age: 30},
			{Name: "Bob", Age: 30},
			{Name: "Charlie", Age: 25},
		}

		fieldName := "Nonexistent"

		result, err := GroupBy[int](people, fieldName)
		assert.Error(t, err)
		assert.Equal(t, "groupBy: field Nonexistent does not exist", err.Error())

		assert.Nil(t, result)
	})

}

func TestGroupBy1by1(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	t.Run("Success_groupBy_age", func(t *testing.T) {
		people := []Person{
			{Name: "Alice", Age: 30},
			{Name: "Bob", Age: 31},
			{Name: "Charlie", Age: 25},
		}

		fieldName := "Age"

		result, err := GroupBy1By1[int](people, fieldName)
		assert.NoError(t, err)

		expected := map[int]Person{
			30: people[0],
			31: people[1],
			25: people[2],
		}

		assert.Equal(t, expected, result)
	})
	t.Run("error_dupe_key_groupBy_age", func(t *testing.T) {
		people := []Person{
			{Name: "Alice", Age: 30},
			{Name: "Bob", Age: 30},
			{Name: "Charlie", Age: 25},
		}

		fieldName := "Age"

		_, err := GroupBy1By1[int](people, fieldName)
		assert.Error(t, err)
		assert.Equal(t, "groupBy: field Age is not unique", err.Error())
	})

	t.Run("Success_groupBy_name", func(t *testing.T) {
		people := []Person{
			{Name: "Alice", Age: 30},
			{Name: "Bob", Age: 30},
			{Name: "Charlie", Age: 25},
			{Name: "Damx", Age: 25},
		}

		fieldName := "Name"

		result, err := GroupBy1By1[string](people, fieldName)
		assert.NoError(t, err)

		expected := map[string]Person{
			"Alice":   people[0],
			"Bob":     people[1],
			"Charlie": people[2],
			"Damx":    people[3],
		}

		assert.Equal(t, expected, result)
	})

	t.Run("Success_groupBy_layer2_field", func(t *testing.T) {
		type Layer2 struct {
			Field1 string
			Field2 int
		}

		type Layer1 struct {
			Name   string
			Age    int
			Layer2 Layer2
		}

		people := []Layer1{
			{
				Name: "Alice",
				Age:  30,
				Layer2: Layer2{
					Field1: "Value1",
					Field2: 42,
				},
			},
			{
				Name: "Bob",
				Age:  30,
				Layer2: Layer2{
					Field1: "Value2",
					Field2: 43,
				},
			},
			{
				Name: "Charlie",
				Age:  25,
				Layer2: Layer2{
					Field1: "Value3",
					Field2: 44,
				},
			},
			{
				Name: "x",
				Age:  28,
				Layer2: Layer2{
					Field1: "Value4",
					Field2: 45,
				},
			},
		}

		fieldName := "Layer2.Field1"

		result, err := GroupBy1By1[string](people, fieldName)
		assert.NoError(t, err)

		expected := map[string]Layer1{
			"Value1": people[0],
			"Value2": people[1],
			"Value3": people[2],
			"Value4": people[3],
		}
		assert.Equal(t, expected, result)
	})
	t.Run("Error_dupe_key_groupBy_name", func(t *testing.T) {
		people := []Person{
			{Name: "Alice", Age: 30},
			{Name: "Bob", Age: 30},
			{Name: "Charlie", Age: 25},
			{Name: "Charlie", Age: 27},
		}

		fieldName := "Name"

		_, err := GroupBy1By1[string](people, fieldName)
		assert.Error(t, err)
		assert.Equal(t, "groupBy: field Name is not unique", err.Error())
	})

	t.Run("Error_key_dupe_groupBy_layer2_field", func(t *testing.T) {
		type Layer2 struct {
			Field1 string
			Field2 int
		}

		type Layer1 struct {
			Name   string
			Age    int
			Layer2 Layer2
		}

		people := []Layer1{
			{
				Name: "Alice",
				Age:  30,
				Layer2: Layer2{
					Field1: "Value1",
					Field2: 42,
				},
			},
			{
				Name: "Bob",
				Age:  30,
				Layer2: Layer2{
					Field1: "Value2",
					Field2: 43,
				},
			},
			{
				Name: "Charlie",
				Age:  25,
				Layer2: Layer2{
					Field1: "Value3",
					Field2: 44,
				},
			},
			{
				Name: "x",
				Age:  28,
				Layer2: Layer2{
					Field1: "Value1",
					Field2: 45,
				},
			},
		}

		fieldName := "Layer2.Field1"

		_, err := GroupBy1By1[string](people, fieldName)
		assert.Error(t, err)
		assert.Equal(t, "groupBy: field Layer2.Field1 is not unique", err.Error())
	})

	t.Run("Success_groupBy_layer3_field", func(t *testing.T) {

		type Layer3 struct {
			Field3 string
			Field4 bool
		}

		type Layer2 struct {
			Field1 string
			Field2 int
			Layer3 Layer3
		}

		type Layer1 struct {
			Name   string
			Age    int
			Layer2 Layer2
		}

		people := []Layer1{
			{
				Name: "Alice",
				Age:  30,
				Layer2: Layer2{
					Field1: "Value1",
					Field2: 42,
					Layer3: Layer3{
						Field3: "layer3-1",
						Field4: true,
					},
				},
			},
			{
				Name: "Bob",
				Age:  30,
				Layer2: Layer2{
					Field1: "Value2",
					Field2: 43,
					Layer3: Layer3{
						Field3: "layer3-2",
						Field4: false,
					},
				},
			},
			{
				Name: "Charlie",
				Age:  25,
				Layer2: Layer2{
					Field1: "Value3",
					Field2: 44,
					Layer3: Layer3{
						Field3: "layer3-3",
						Field4: true,
					},
				},
			},
			{
				Name: "x",
				Age:  28,
				Layer2: Layer2{
					Field1: "Value1",
					Field2: 45,
					Layer3: Layer3{
						Field3: "layer3-4",
						Field4: true,
					},
				},
			},
		}

		fieldName := "Layer2.Layer3.Field3"

		result, err := GroupBy1By1[string](people, fieldName)
		assert.NoError(t, err)

		expected := map[string]Layer1{
			"layer3-1": people[0],
			"layer3-2": people[1],
			"layer3-3": people[2],
			"layer3-4": people[3],
		}

		assert.Equal(t, expected, result)
	})
	t.Run("Error_dupe_key_groupBy_layer3_field", func(t *testing.T) {

		type Layer3 struct {
			Field3 string
			Field4 bool
		}

		type Layer2 struct {
			Field1 string
			Field2 int
			Layer3 Layer3
		}

		type Layer1 struct {
			Name   string
			Age    int
			Layer2 Layer2
		}

		people := []Layer1{
			{
				Name: "Alice",
				Age:  30,
				Layer2: Layer2{
					Field1: "Value1",
					Field2: 42,
					Layer3: Layer3{
						Field3: "layer3-1",
						Field4: true,
					},
				},
			},
			{
				Name: "Bob",
				Age:  30,
				Layer2: Layer2{
					Field1: "Value2",
					Field2: 43,
					Layer3: Layer3{
						Field3: "layer3-2",
						Field4: false,
					},
				},
			},
			{
				Name: "Charlie",
				Age:  25,
				Layer2: Layer2{
					Field1: "Value3",
					Field2: 44,
					Layer3: Layer3{
						Field3: "layer3-3",
						Field4: true,
					},
				},
			},
			{
				Name: "x",
				Age:  28,
				Layer2: Layer2{
					Field1: "Value1",
					Field2: 45,
					Layer3: Layer3{
						Field3: "layer3-2",
						Field4: true,
					},
				},
			},
		}

		fieldName := "Layer2.Layer3.Field3"

		_, err := GroupBy1By1[string](people, fieldName)
		assert.Error(t, err)
		assert.Equal(t, "groupBy: field Layer2.Layer3.Field3 is not unique", err.Error())
	})

	t.Run("Error_invalid_field_name", func(t *testing.T) {
		people := []Person{
			{Name: "Alice", Age: 30},
			{Name: "Bob", Age: 30},
			{Name: "Charlie", Age: 25},
		}

		fieldName := "Nonexistent"

		result, err := GroupBy1By1[int](people, fieldName)
		assert.Error(t, err)
		assert.Equal(t, "groupBy: field Nonexistent does not exist", err.Error())

		assert.Nil(t, result)
	})

}

// TestDistinct tests the Distinct function for various slice types.
func TestDistinct(t *testing.T) {
	tests := []struct {
		name     string
		slice    interface{}
		expected interface{}
	}{
		{
			name:     "ints",
			slice:    []int{1, 2, 3, 2, 4, 5, 4, 6},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "strings",
			slice:    []string{"apple", "banana", "apple", "orange", "banana"},
			expected: []string{"apple", "banana", "orange"},
		},
		{
			name:     "empty int slice",
			slice:    []int{},
			expected: []int{},
		},
		{
			name:     "empty string slice",
			slice:    []string{},
			expected: []string{},
		},
		{
			name:     "no duplicates",
			slice:    []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var unique interface{}
			switch tc.slice.(type) {
			case []int:
				unique = Distinct(tc.slice.([]int))
			case []string:
				unique = Distinct(tc.slice.([]string))
			default:
				t.Fatalf("Unsupported type in test case")
			}

			if !reflect.DeepEqual(unique, tc.expected) {
				t.Errorf("Distinct(%v) = %v, want %v", tc.slice, unique, tc.expected)
			}
		})
	}
}

func TestDistinctFunc(t *testing.T) {
	tests := []struct {
		name     string
		slice    interface{}
		fn       interface{}
		expected interface{}
	}{
		{
			name:  "ints",
			slice: []int{1, 2, 3, 2, 4, 5, 4, 6},
			fn: func(i, j int) bool {
				return i == j
			},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:  "strings",
			slice: []string{"apple", "banana", "apple", "orange", "banana"},
			fn: func(i, j string) bool {
				return i == j
			},
			expected: []string{"apple", "banana", "orange"},
		},
		{
			name:  "empty int slice",
			slice: []int{},
			fn: func(i, j int) bool {
				return i == j
			},
			expected: []int{},
		},
		{
			name:  "empty string slice",
			slice: []string{},
			fn: func(i, j string) bool {
				return i == j
			},
			expected: []string{},
		},
		{
			name:  "no duplicates",
			slice: []int{1, 2, 3},
			fn: func(i, j int) bool {
				return i == j
			},
			expected: []int{1, 2, 3},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var unique interface{}
			switch tc.slice.(type) {
			case []int:
				unique = DistinctFunc(tc.slice.([]int), tc.fn.(func(int, int) bool))
			case []string:
				unique = DistinctFunc(tc.slice.([]string), tc.fn.(func(string, string) bool))
			default:
				t.Fatalf("Unsupported type in test case")
			}

			if !reflect.DeepEqual(unique, tc.expected) {
				t.Errorf("DistinctFunc(%v) = %v, want %v", tc.slice, unique, tc.expected)
			}
		})
	}
}

func TestSort_Ints(t *testing.T) {
	intSlice := []int{5, 2, 8, 1, 9}
	expected := []int{1, 2, 5, 8, 9}

	sorted := Sort(intSlice, func(i, j int) bool { return intSlice[i] < intSlice[j] })

	if !reflect.DeepEqual(sorted, expected) {
		t.Errorf("Sort failed for int slice. Got: %v, Expected: %v", intSlice, expected)
	}
}

func TestSort_Strings(t *testing.T) {
	stringSlice := []string{"c", "a", "b"}
	expected := []string{"a", "b", "c"}

	sorted := Sort(stringSlice, func(i, j int) bool { return stringSlice[i] < stringSlice[j] })

	if !reflect.DeepEqual(sorted, expected) {
		t.Errorf("Sort failed for string slice. Got: %v, Expected: %v", stringSlice, expected)
	}
}

func TestSort_CustomType(t *testing.T) {
	// Define a custom type for testing
	type Person struct {
		Name string
		Age  int
	}

	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}
	expected := []Person{
		{"Bob", 25},
		{"Alice", 30},
		{"Charlie", 35},
	}

	sorted := Sort(people, func(i, j int) bool { return people[i].Age < people[j].Age })

	if !reflect.DeepEqual(sorted, expected) {
		t.Errorf("Sort failed for custom type slice. Got: %v, Expected: %v", people, expected)
	}
}

func TestSort_StringsByReverseOrder(t *testing.T) {
	stringSlice := []string{"apple", "banana", "cherry"}
	expected := []string{"cherry", "banana", "apple"}

	sorted := Sort(stringSlice, func(i, j int) bool { return stringSlice[i] > stringSlice[j] })

	if !reflect.DeepEqual(sorted, expected) {
		t.Errorf("Sort failed for string sort by reverse order. Got: %v, Expected: %v", stringSlice, expected)
	}
}

func TestSort_StringsCaseInsensitive(t *testing.T) {
	stringSlice := []string{"banana", "Apple", "CHERRY"}
	expected := []string{"Apple", "banana", "CHERRY"}

	sorted := Sort(stringSlice, func(i, j int) bool {
		return strings.ToLower(stringSlice[i]) < strings.ToLower(stringSlice[j])
	})

	if !reflect.DeepEqual(sorted, expected) {
		t.Errorf("Sort failed for string sort case insensitive. Got: %v, Expected: %v", stringSlice, expected)
	}
}

func TestSort_StringsByLength(t *testing.T) {
	stringSlice := []string{"ccccc", "aaa", "bbbb"}
	expected := []string{"aaa", "bbbb", "ccccc"}

	sorted := Sort(stringSlice, func(i, j int) bool { return len(stringSlice[i]) < len(stringSlice[j]) })

	if !reflect.DeepEqual(sorted, expected) {
		t.Errorf("Sort failed for string sort by length. Got: %v, Expected: %v", stringSlice, expected)
	}
}
func TestSort_StringsByLength_reversed(t *testing.T) {
	stringSlice := []string{"ccccc", "aaa", "bbbb"}
	expected := []string{"ccccc", "bbbb", "aaa"}

	sorted := Sort(stringSlice, func(i, j int) bool { return len(stringSlice[i]) > len(stringSlice[j]) })

	if !reflect.DeepEqual(sorted, expected) {
		t.Errorf("Sort failed for string sort by length. Got: %v, Expected: %v", stringSlice, expected)
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		name     string
		slice    interface{}
		expected interface{}
	}{
		{
			name:     "ints",
			slice:    []int{1, 2, 3, 4, 5},
			expected: 15,
		},
		{
			name:     "floats",
			slice:    []float64{1.1, 2.2, 3.3, 4.4, 5.5},
			expected: 16.5,
		},
		{
			name:     "empty int slice",
			slice:    []int{},
			expected: 0,
		},
		{
			name:     "empty float slice",
			slice:    []float64{},
			expected: 0.0,
		},
		{
			name:     "no elements",
			slice:    []int{},
			expected: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var sum interface{}
			switch tc.slice.(type) {
			case []int:
				sum = Sum(tc.slice.([]int))
			case []float64:
				sum = Sum(tc.slice.([]float64))
			default:
				t.Fatalf("Unsupported type in test case")
			}

			if !reflect.DeepEqual(sum, tc.expected) {
				t.Errorf("Sum(%v) = %v, want %v", tc.slice, sum, tc.expected)
			}
		})
	}
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

func TestMapReturnWithError(t *testing.T) {
	t.Run("Success", func(t *testing.T) {

		source := []int{1, 2, 3, 4, 5}

		mappingFunc := func(data int) (int, error) {
			return data * 2, nil
		}

		result, err := MapReturnWithError(source, mappingFunc)
		assert.NoError(t, err)

		expected := []int{2, 4, 6, 8, 10}
		assert.Equal(t, expected, result)
	})

	t.Run("some_element_has_Error", func(t *testing.T) {

		source := []int{1, 2, 3, 4, 5}

		mappingFunc := func(data int) (int, error) {
			if data == 3 {
				return 0, errors.New("fake error for 3")
			}
			return data * 2, nil
		}

		result, err := MapReturnWithError(source, mappingFunc)
		assert.Error(t, err)
		assert.Equal(t, "error mapping at index:'2', error: fake error for 3", err.Error())

		assert.Nil(t, result)
	})

}

func TestHigherOrderFunction_FlatMap(t *testing.T) {
	t.Run("Success_Int", func(t *testing.T) {

		source := [][]int{
			{1, 2, 3},
			{4, 5},
			{6, 7, 8},
		}

		result := FlatMap(source)

		expected := []int{1, 2, 3, 4, 5, 6, 7, 8}
		assert.Equal(t, expected, result)
	})

	t.Run("Success_String", func(t *testing.T) {

		source := [][]string{
			{"a", "b", "c"},
			{"d", "e"},
			{"f", "g", "h"},
		}

		result := FlatMap(source)

		expected := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
		assert.Equal(t, expected, result)
	})
}

func TestHigherOrderFunction_Reduce(t *testing.T) {
	t.Run("Success_Int", func(t *testing.T) {

		source := []int{1, 2, 3, 4, 5}

		reduceFunc := func(acc, value int) int {
			return acc + value
		}

		result := Reduce[int](source, reduceFunc, 0)

		expected := 15
		assert.Equal(t, expected, result)
	})

	t.Run("Success_String", func(t *testing.T) {
		source := []string{"D", "a", "r", "k", " ", "m", "a", "g", "i", "c"}

		reduceFunc := func(acc, value string) string {
			return acc + value
		}

		result := Reduce[string](source, reduceFunc, "")

		expected := "Dark magic"
		assert.Equal(t, expected, result)
	})

	t.Run("Success_Empty_List_Int", func(t *testing.T) {
		source := []int{}

		reduceFunc := func(acc, value int) int {
			return acc + value
		}

		result := Reduce[int](source, reduceFunc, 0)

		expected := 0
		assert.Equal(t, expected, result)
	})
}

func TestHigherOrderFunction_Sort(t *testing.T) {
	t.Run("Success_Int", func(t *testing.T) {

		source := []int{5, 2, 8, 1, 9}

		sortFunc := func(i, j int) bool {
			return source[i] < source[j]
		}

		sorted := Sort(source, sortFunc)

		expected := []int{1, 2, 5, 8, 9}
		assert.Equal(t, expected, sorted)
	})

	t.Run("Success_Int_reverse", func(t *testing.T) {
		source := []int{5, 2, 8, 1, 9}

		sortFunc := func(i, j int) bool {
			return source[i] > source[j]
		}

		sorted := Sort(source, sortFunc)

		expected := []int{9, 8, 5, 2, 1}
		assert.Equal(t, expected, sorted)
	})

	t.Run("Success_String", func(t *testing.T) {

		source := []string{"c", "a", "b"}

		sortFunc := func(i, j int) bool {
			return source[i] < source[j]
		}

		sorted := Sort(source, sortFunc)

		expected := []string{"a", "b", "c"}
		assert.Equal(t, expected, sorted)
	})

	t.Run("Success_String_reverse", func(t *testing.T) {

		source := []string{"c", "a", "b"}

		sortFunc := func(i, j int) bool {
			return source[i] > source[j]
		}

		sorted := Sort(source, sortFunc)

		expected := []string{"c", "b", "a"}
		assert.Equal(t, expected, sorted)
	})

	t.Run("Success_CustomType", func(t *testing.T) {

		type Person struct {
			Name string
			Age  int
		}

		source := []Person{
			{"Alice", 30},
			{"Bob", 25},
			{"Charlie", 35},
		}

		sortFunc := func(i, j int) bool {
			return source[i].Age < source[j].Age
		}

		sorted := Sort(source, sortFunc)

		expected := []Person{
			{"Bob", 25},
			{"Alice", 30},
			{"Charlie", 35},
		}
		assert.Equal(t, expected, sorted)
	})

	t.Run("Success_CustomType_reverse", func(t *testing.T) {

		type Person struct {
			Name string
			Age  int
		}

		source := []Person{
			{"Alice", 30},
			{"Bob", 25},
			{"Charlie", 35},
		}

		sortFunc := func(i, j int) bool {
			return source[i].Age > source[j].Age
		}

		sorted := Sort(source, sortFunc)

		expected := []Person{
			{"Charlie", 35},
			{"Alice", 30},
			{"Bob", 25},
		}
		assert.Equal(t, expected, sorted)
	})

	t.Run("Success_sort_2_layers_of_customerType_sort_customerCode_and_SalesOrderNumber", func(t *testing.T) {

		type SalesOrder struct {
			CustomerCode     string
			SalesOrderNumber string
			Amount           float64
		}

		source := []SalesOrder{
			{"C2", "S2", 200},
			{"C1", "S3", 300},
			{"C2", "S4", 400},
			{"C1", "S1", 100},
		}

		sortFunc := func(i, j int) bool {
			if source[i].CustomerCode == source[j].CustomerCode {
				return source[i].SalesOrderNumber < source[j].SalesOrderNumber
			}
			return source[i].CustomerCode < source[j].CustomerCode
		}

		sorted := Sort(source, sortFunc)

		expected := []SalesOrder{
			{"C1", "S1", 100},
			{"C1", "S3", 300},
			{"C2", "S2", 200},
			{"C2", "S4", 400},
		}
		assert.Equal(t, expected, sorted)
	})
}

func TestMapHashMapToHashMap(t *testing.T) {
	t.Run("Success", func(t *testing.T) {

		source := map[string]int{
			"apple":  1,
			"banana": 2,
			"cherry": 3,
		}

		mappingFunc := func(key string, value int) string {
			return key + " " + strconv.Itoa(value)
		}

		result := MapHashMapToHashMap(source, mappingFunc)

		expected := map[string]string{
			"apple":  "apple 1",
			"banana": "banana 2",
			"cherry": "cherry 3",
		}

		assert.Equal(t, expected, result)
	})

	t.Run("Success_empty_map", func(t *testing.T) {

		source := map[string]int{}

		mappingFunc := func(key string, value int) string {
			return key + " " + strconv.Itoa(value)
		}

		result := MapHashMapToHashMap(source, mappingFunc)

		expected := map[string]string{}

		assert.Equal(t, expected, result)
	})
}

func TestMapHashMapToHashMapWithError(t *testing.T) {
	t.Run("Success", func(t *testing.T) {

		source := map[string]int{
			"apple":  1,
			"banana": 2,
			"cherry": 3,
		}

		mappingFunc := func(key string, value int) (string, error) {
			return key + " " + strconv.Itoa(value), nil
		}

		result, err := MapHashMapToHashMapReturnWithError(source, mappingFunc)
		assert.NoError(t, err)

		expected := map[string]string{
			"apple":  "apple 1",
			"banana": "banana 2",
			"cherry": "cherry 3",
		}

		assert.Equal(t, expected, result)
	})

	t.Run("Success_empty_map", func(t *testing.T) {

		source := map[string]int{}

		mappingFunc := func(key string, value int) (string, error) {
			return key + " " + strconv.Itoa(value), nil
		}

		result, err := MapHashMapToHashMapReturnWithError(source, mappingFunc)
		assert.NoError(t, err)

		expected := map[string]string{}

		assert.Equal(t, expected, result)
	})

	t.Run("Error_some_element_has_Error", func(t *testing.T) {

		source := map[string]int{
			"apple":  1,
			"banana": 2,
			"cherry": 3,
		}

		mappingFunc := func(key string, value int) (string, error) {
			if key == "banana" {
				return "", errors.New("fake error for banana")
			}
			return key + " " + strconv.Itoa(value), nil
		}

		result, err := MapHashMapToHashMapReturnWithError(source, mappingFunc)
		assert.Error(t, err)
		assert.Equal(t, "error mapping at key:'banana', error: fake error for banana", err.Error())

		assert.Nil(t, result)
	})
}

func TestMapHashMapToList(t *testing.T) {
	t.Run("Success", func(t *testing.T) {

		source := map[string]int{
			"apple":  1,
			"banana": 2,
			"cherry": 3,
		}

		mappingFunc := func(key string, value int) string {
			return key + " " + strconv.Itoa(value)
		}

		result := MapHashMapToList(source, mappingFunc)

		expected := []string{
			"apple 1",
			"banana 2",
			"cherry 3",
		}

		assert.Equal(t, expected, result)
	})

	t.Run("Success_empty_map", func(t *testing.T) {

		source := map[string]int{}

		mappingFunc := func(key string, value int) string {
			return key + " " + strconv.Itoa(value)
		}

		result := MapHashMapToList(source, mappingFunc)

		expected := []string{}

		assert.Equal(t, expected, result)
	})
}

func TestMapHashMapToListReturnWithError(t *testing.T) {
	t.Run("Success", func(t *testing.T) {

		source := map[string]int{
			"apple":  1,
			"banana": 2,
			"cherry": 3,
		}

		mappingFunc := func(key string, value int) (string, error) {
			return key + " " + strconv.Itoa(value), nil
		}

		result, err := MapHashMapToListReturnWithError(source, mappingFunc)
		assert.NoError(t, err)

		expected := []string{
			"apple 1",
			"banana 2",
			"cherry 3",
		}

		assert.Equal(t, expected, result)
	})

	t.Run("Success_empty_map", func(t *testing.T) {

		source := map[string]int{}

		mappingFunc := func(key string, value int) (string, error) {
			return key + " " + strconv.Itoa(value), nil
		}

		result, err := MapHashMapToListReturnWithError(source, mappingFunc)
		assert.NoError(t, err)

		expected := []string{}

		assert.Equal(t, expected, result)
	})

	t.Run("Error_some_element_has_Error", func(t *testing.T) {

		source := map[string]int{
			"apple":  1,
			"banana": 2,
			"cherry": 3,
		}

		mappingFunc := func(key string, value int) (string, error) {
			if key == "banana" {
				return "", errors.New("fake error for banana")
			}
			return key + " " + strconv.Itoa(value), nil
		}

		result, err := MapHashMapToListReturnWithError(source, mappingFunc)
		assert.Error(t, err)
		assert.Equal(t, "error mapping at key:'banana', error: fake error for banana", err.Error())

		assert.Nil(t, result)
	})
}

func TestSliceToHashMap(t *testing.T) {
	t.Run("Success_Int", func(t *testing.T) {

		source := []int{1, 2, 3, 4, 5}

		result := SliceToHashMap(source)

		expected := map[int]bool{
			1: true,
			2: true,
			3: true,
			4: true,
			5: true,
		}

		assert.Equal(t, expected, result)
	})

	t.Run("Success_String", func(t *testing.T) {

		source := []string{"apple", "banana", "cherry"}

		result := SliceToHashMap(source)

		expected := map[string]bool{
			"apple":  true,
			"banana": true,
			"cherry": true,
		}

		assert.Equal(t, expected, result)
	})

	t.Run("Success_empty_list", func(t *testing.T) {

		source := []int{}

		result := SliceToHashMap(source)

		expected := map[int]bool{}

		assert.Equal(t, expected, result)
	})
}

func TestMapToHashMap(t *testing.T) {
	t.Run("Success_Int", func(t *testing.T) {

		source := []int{1, 2, 3, 4, 5}

		mappingFunc := func(item int) (int, bool) {
			return item, true
		}

		result := MapToHashMap(source, mappingFunc)

		expected := map[int]bool{
			1: true,
			2: true,
			3: true,
			4: true,
			5: true,
		}

		assert.Equal(t, expected, result)
	})

	t.Run("Success_String", func(t *testing.T) {

		source := []string{"apple", "banana", "cherry"}

		mappingFunc := func(item string) (string, bool) {
			return item, true
		}

		result := MapToHashMap(source, mappingFunc)

		expected := map[string]bool{
			"apple":  true,
			"banana": true,
			"cherry": true,
		}

		assert.Equal(t, expected, result)
	})

	t.Run("Success_empty_list", func(t *testing.T) {

		source := []int{}

		mappingFunc := func(item int) (int, bool) {
			return item, true
		}

		result := MapToHashMap(source, mappingFunc)

		expected := map[int]bool{}

		assert.Equal(t, expected, result)
	})
}

func TestMapToHashMapReturnWithError(t *testing.T) {
	t.Run("Success_Int", func(t *testing.T) {

		source := []int{1, 2, 3, 4, 5}

		mappingFunc := func(item int) (int, bool, error) {
			return item, true, nil
		}

		result, err := MapToHashMapReturnWithError(source, mappingFunc)
		assert.NoError(t, err)

		expected := map[int]bool{
			1: true,
			2: true,
			3: true,
			4: true,
			5: true,
		}

		assert.Equal(t, expected, result)
	})

	t.Run("Success_String", func(t *testing.T) {

		source := []string{"apple", "banana", "cherry"}

		mappingFunc := func(item string) (string, bool, error) {
			return item, true, nil
		}

		result, err := MapToHashMapReturnWithError(source, mappingFunc)
		assert.NoError(t, err)

		expected := map[string]bool{
			"apple":  true,
			"banana": true,
			"cherry": true,
		}

		assert.Equal(t, expected, result)
	})

	t.Run("Success_empty_list", func(t *testing.T) {

		source := []int{}

		mappingFunc := func(item int) (int, bool, error) {
			return item, true, nil
		}

		result, err := MapToHashMapReturnWithError(source, mappingFunc)
		assert.NoError(t, err)

		expected := map[int]bool{}

		assert.Equal(t, expected, result)
	})

	t.Run("Error_some_element_has_Error", func(t *testing.T) {

		source := []int{1, 2, 3, 4, 5}

		mappingFunc := func(item int) (int, bool, error) {
			if item == 3 {
				return 0, false, errors.New("fake error for 3")
			}
			return item, true, nil
		}

		result, err := MapToHashMapReturnWithError(source, mappingFunc)
		assert.Error(t, err)
		assert.Equal(t, "error mapping at index:'2', error: fake error for 3", err.Error())

		assert.Nil(t, result)
	})

}
