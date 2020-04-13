package queryfx

import (
	"fmt"
	"reflect"
)

// The integer placeholder is for handling integer placeholder like %C
type integerPlaceholder struct{
	Index int
	Key string
	Data interface{}
	Fragment string
}

func(c integerPlaceholder) Validate() error{
	// Get the raw fragement
	fragment := getRawFragment(c.Fragment)
	if fragment != "d"{
		return fmt.Errorf("Integer placeholder is expecting a fragment of type d")
	}

	// check that the data type is string
	ref := reflect.ValueOf(c.Data)
	kind := ref.Kind()
	if kind != reflect.Int && kind != reflect.Uint64 && kind != reflect.Int64{
		return fmt.Errorf("Integer placeholder is expecting an int type")
	}

	return nil
}

func(c integerPlaceholder) GetKey() string{
	return c.Key
}
func(c integerPlaceholder) Format() string{
	// Assumes that we have validated the string so we need to format the integer for safe usage
	// in a query
	ref := reflect.ValueOf(c.Data)
	var vkind string
	kind := ref.Kind()
	switch kind{
	case reflect.Int:
		vkind = fmt.Sprintf("%d", c.Data.(int))
	case reflect.Uint64:
		vkind = fmt.Sprintf("%d", c.Data.(uint64))
	case reflect.Int64:
		vkind = fmt.Sprintf("%d", c.Data.(int64))
	}

	return vkind
}

func NewIntegerPlaceholder(index int, key, fragment string, data interface{}) Placeholder{
	return integerPlaceholder{Index:index,Data:data,Fragment:fragment,Key:key}
}