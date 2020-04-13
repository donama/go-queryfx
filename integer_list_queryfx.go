package queryfx

import (
	"fmt"
	"reflect"
	"strings"
)

// The integer list placeholder is for handling integer list placeholder like %LC
type integerListPlaceholder struct{
	Index int
	Key string
	Data interface{}
	Fragment string
}

func(c integerListPlaceholder) Validate() error{
	// Get the raw fragment
	fragment := getRawFragment(c.Fragment)
	if fragment != "Ld"{
		return fmt.Errorf("Integer list placeholder is expecting a fragment of type Ld, got %q", fragment)
	}

	// check that the data type is string
	ref := reflect.ValueOf(c.Data)
	if ref.Kind() != reflect.Slice{
		return fmt.Errorf("Integer list placeholder is expecting a slice of integer types")
	}

	// sub check eacn entry to ensure that its aa string type
	temp := c.Data.([]int)
	for _, e := range temp{
		if reflect.ValueOf(e).Kind() != reflect.Int{
			return fmt.Errorf("Integer list has an entry that is not a string type : %q", e)
		}
	}
	return nil
}

func(c integerListPlaceholder) GetKey() string{
	return c.Key
}
func(c integerListPlaceholder) Format() string{
	// Assumes that we have validated the string so we need to format the integer for safe usage
	// in a query
	data := c.Data.([]int)
	temp := []string{}
	for _,e := range data{
		temp = append(temp, fmt.Sprintf("%d", e))
	}
	return strings.Join(temp, " , ")
}

func NewIntegerListPlaceholder(index int, key, fragment string, data interface{}) Placeholder{
	return integerListPlaceholder{Index:index,Data:data,Fragment:fragment,Key:key}
}