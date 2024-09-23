package maps

import (
	"errors"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
