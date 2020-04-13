// The QueryFX package provides a way to format SQL query statement such that they are safe to execute against a
// database

package queryfx

import (
	"fmt"
	"regexp"
	"strings"
)

//FormatQuery takes a query string with placeholders in it and then a list of parameters that should be used
// as substitute for the placesholders, then it performs data type verifications and then finally returns
// a query string with the provided data inplace and ready to be used against an SQL database server, this methods
// performs all the needed data escape to prevent SQL injection

func FormatQuery(query string , params ...interface{})(string, error){
	exp := regexp.MustCompile(`(%[\w+]+)`)

	if query == ""{
		return "", fmt.Errorf("Query string is not defined")
	}

	// First we need to get a list of the placeholders in the query string, if we do not find any we return
	// the query the way it came asumming that it does not have any placeholders
	var placeholders []string = []string{}

	placeholders =  exp.FindAllString(query, -1)

	if len(placeholders) == 0{
		return query, nil
	}

	// If we have a placeholders, we need to check to see if the number of parameters provided
	// matches with the number of placeholders we have figured out
	if len(placeholders) != len(params){
		return query, fmt.Errorf("Number of placeholders in query string mismatched the number of passed ins parameters")
	}

	// positionally loop through the placeholders building a collections of the index, placeholder and the
	// passed in data for this placeholder, we go positionally because its possible to have a repeating
	// placeholder
	var validators = []Placeholder{}
	var idx int = 0

	cquery := query
	for i:= 0; i<len(placeholders);i++{
		idx += 1
		key := makeQueryKey(idx)
		fragment := placeholders[i]
		data := params[i]
		// replace the fragment in the query string
		validator,err := makePlaceholderValidator(idx, key,fragment, data)
		if err != nil{
			return cquery, err
		}
		query = strings.Replace(query, fragment, key, 1)
		validators = append(validators, validator)
	}

	// validate all placeholder in the query string, this also check the data type to ensure it
	// matches the placeholder type
	for _, validator := range validators{
		err := validator.Validate()
		if err != nil{
			return cquery, err
		}
	}

	// next we format the query placeholders, replacing each placeholder with the actual escaped data,
	// this will bloat the query string and eventuall we will have a standard SQl query we can send to a database
	// for executrion
	for _, entry := range validators{
		query = strings.Replace(query, entry.GetKey(), entry.Format(), 1)
	}

	return query, nil
}