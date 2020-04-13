package queryfx

import (
	"errors"
	"fmt"
	"strings"
)


//Placeholder defines the methods that all concrete types must implement to be able t handle
// specific type of query placeholder and its parameters
type Placeholder interface{
	Validate() error
	Format() string
	GetKey() string
}

// Returns a string key used for indexing placeholder in a query string as its scanned from left
// to right
func makeQueryKey(idx int) string{
	return fmt.Sprintf("~%dA", idx)
}

//getRawFragment returns the placeholder string without the % prefix
func getRawFragment(f string) string{
	return strings.Trim(strings.Replace(f,"%","", 1),"")
}

// this defines the function signature of the placeholder validators
type placeholderFunc func(index int, key, fragment string, data interface{}) Placeholder

//placeholderFunctionMaps is a map defining a fragment type and its validator creator function
var placeholderFunctionMaps = map[string]placeholderFunc{

	"LC":NewColumnListPlaceholder,
	"C":NewColumnPlaceholder,
	"s":NewStringPlaceholder,
	"Ls":NewStringListPlaceholder,
	"d":NewIntegerPlaceholder,
	"Ld":NewIntegerListPlaceholder,
	"T":NewTablePlaceholder,
}

//makePlaceholderValidator build and returns the validator for each given placeholder fragment
// it will return an error if no validator is found for a given fragment
func makePlaceholderValidator(index int, key, fragment string, data interface{}) (Placeholder, error){
	// get the raw placeholder type we want to find a validator for
	validator_type := getRawFragment(fragment)

	// Dowe have a validator for this placeholder if it exists in the map we will execute
	// and return it otherwise we return an error
	if validator, ok := placeholderFunctionMaps[validator_type]; ok{
		return validator(index, key, fragment, data), nil
	}
	return nil, errors.New(fmt.Sprintf("Mo validator defined for placeholder %s", fragment))
}
