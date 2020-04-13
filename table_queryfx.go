package queryfx

import (
	"fmt"
	"reflect"
)

// The table placeholder is for handling table placeholder like %C
type tablePlaceholder struct{
	Index int
	Key string
	Data interface{}
	Fragment string
}

func(c tablePlaceholder) Validate() error{
	// Get the raw fragment
	fragment := getRawFragment(c.Fragment)
	if fragment != "T"{
		return fmt.Errorf("table placeholder is expecting a fragment of type T")
	}

	// check that the data type is string
	ref := reflect.ValueOf(c.Data)
	if ref.Kind() != reflect.String{
		return fmt.Errorf("table place holder is expecting a string type")
	}

	return nil
}

func(c tablePlaceholder) GetKey() string{
	return c.Key
}
func(c tablePlaceholder) Format() string{
	// Assumes that we have validated the string so we need to format the table for safe usage
	// in a query
	data := c.Data.(string)
	return fmt.Sprintf("`%s`", data)
}

func NewTablePlaceholder(index int, key, fragment string, data interface{}) Placeholder{
	return tablePlaceholder{Index:index,Data:data,Fragment:fragment,Key:key}
}