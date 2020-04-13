package queryfx

import (
	"fmt"
	"reflect"
)

// The column placeholder is for handling column placeholder like %C
type columnPlaceholder struct{
	Index int
	Key string
	Data interface{}
	Fragment string
}

func(c columnPlaceholder) Validate() error{
	// Get the raw fragement
	fragment := getRawFragment(c.Fragment)
	if fragment != "C"{
		return fmt.Errorf("Column placeholder is expecting a fragment of type C")
	}

	// check that the data type is string
	ref := reflect.ValueOf(c.Data)
	if ref.Kind() != reflect.String{
		return fmt.Errorf("Column place holder is expecting a string type")
	}

	return nil
}

func(c columnPlaceholder) GetKey() string{
	return c.Key
}
func(c columnPlaceholder) Format() string{
	// Assumes that we have validated the string so we need to format the column for safe usage
	// in a query
	data := c.Data.(string)
	return fmt.Sprintf("%s", data)
}

func NewColumnPlaceholder(index int, key, fragment string, data interface{}) Placeholder{
	return columnPlaceholder{Index:index,Data:data,Fragment:fragment,Key:key}
}