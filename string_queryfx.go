package queryfx

import (
	"fmt"
	"reflect"
)

// The string placeholder is for handling column placeholder like %s
type stringPlaceholder struct{
	Index int
	Key string
	Data interface{}
	Fragment string
}

func(c stringPlaceholder) Validate() error{
	// Get the raw fragement
	fragment := getRawFragment(c.Fragment)
	if fragment != "s"{
		return fmt.Errorf("Column placeholder is expecting a fragment of type s")
	}

	// check that the data type is string
	ref := reflect.ValueOf(c.Data)
	if ref.Kind() != reflect.String{
		return fmt.Errorf("String placeholder is expecting a string type")
	}

	return nil
}

func(c stringPlaceholder) GetKey() string{
	return c.Key
}
func(c stringPlaceholder) Format() string{
	// Assumes that we have validated the string so we need to format the column for safe usage
	// in a query
	data := c.Data.(string)
	return fmt.Sprintf("'%s'", data)
}

func NewStringPlaceholder(index int, key, fragment string, data interface{}) Placeholder{
	return stringPlaceholder{Index:index,Data:data,Fragment:fragment,Key:key}
}