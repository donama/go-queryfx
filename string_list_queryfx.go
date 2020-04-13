package queryfx

import (
	"fmt"
	"reflect"
	"strings"
)

// The string list placeholder is for handling string list placeholder like %LC
type stringListPlaceholder struct{
	Index int
	Key string
	Data interface{}
	Fragment string
}

func(c stringListPlaceholder) Validate() error{
	// Get the raw fragment
	fragment := getRawFragment(c.Fragment)
	if fragment != "Ls"{
		return fmt.Errorf("string list placeholder is expecting a fragment of type Ls")
	}

	// check that the data type is string
	ref := reflect.ValueOf(c.Data)
	if ref.Kind() != reflect.Slice{
		return fmt.Errorf("string list placeholder is expecting a slice of string types")
	}

	// sub check eacn entry to ensure that its aa string type
	temp := c.Data.([]string)
	for _, e := range temp{
		if reflect.ValueOf(e).Kind() != reflect.String{
			return fmt.Errorf("string list has an entry that is not a string type : %q", e)
		}
	}
	return nil
}

func(c stringListPlaceholder) GetKey() string{
	return c.Key
}
func(c stringListPlaceholder) Format() string{
	// Assumes that we have validated the string so we need to format the string for safe usage
	// in a query
	data := c.Data.([]string)
	temp := []string{}
	for _,e := range data{
		temp = append(temp,fmt.Sprintf("'%s'", e))
	}
	return strings.Join(temp, " , ")
}

func NewStringListPlaceholder(index int, key, fragment string, data interface{}) Placeholder{
	return stringListPlaceholder{Index:index,Data:data,Fragment:fragment,Key:key}
}