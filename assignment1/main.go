// You can edit this code!
// Click here and start typing.
package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	// Sample JSON data
	var jsonstr string
	jsonstr = `{
        "name": "Tolexo Online Pvt. Ltd",
        "age_in_years": 8.5,
        "origin": "Noida",
        "head_office": "Noida, Uttar Pradesh",
        "address": [
            {
                "street": "91 Springboard",
                "landmark": "Axis Bank",
                "city": "Noida",
                "pincode": 201301,
                "state": "Uttar Pradesh"
            },
            {
                "street": "91 Springboard",
                "landmark": "Axis Bank",
                "city": "Noida",
                "pincode": 201301,
                "state": "Uttar Pradesh"
            }
        ],
        "sponsors": {
            "name": "One"
        },
        "revenue": "19.8 million$",
        "no_of_employee": 630,
        "str_text": ["one", "two"],
        "int_text": [1, 3, 4]
    }`

	var data map[string]interface{}

	if err := json.Unmarshal([]byte(jsonstr), &data); err != nil {
		panic("could not convernt")
	}
	printValueType(data)
}
func printValueType(data map[string]interface{}) {
	// Iterate over each key-value pair in the map
	for i, val := range data {
		// Get the reflection value of the current value
		val := reflect.ValueOf(val)

		// Use a switch statement to handle different kinds of values
		switch val.Kind() {
		case reflect.String:
			// If the value is a string, print its key, value, and type
			fmt.Printf("key is :%v , value is %v , type is %v\n", i, val, val.Kind())
		case reflect.Slice:
			fmt.Printf("key is :%v, value is %v, type is %v\n ", i, val, val.Type())
			fmt.Println("following is the breakdown")
			nestedSlice := val.Interface().([]interface{})
			for _, v := range nestedSlice {
				nestedValue := reflect.ValueOf(v)
				if nestedValue.Kind() == reflect.Map {
					printValueType(v.(map[string]interface{}))
				} else {
					fmt.Println("\t", v, reflect.TypeOf(v))
					fmt.Printf("key is :%v, value is %v, type is %v\n ", reflect.TypeOf(v), val, val.Type())
				}

			}
		case reflect.Map:
			fmt.Println("this is Map")
			printValueType(val.Interface().(map[string]interface{}))

		}
	}

}
