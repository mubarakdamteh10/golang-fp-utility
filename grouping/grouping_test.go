package grouping

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
