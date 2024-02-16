package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name    string
	Age     int
	Address Address
}

type Address struct {
	City  string
	State string
}

/*
The Elem() method is then called on the reflect.Value. This method is used when the original reflect.Value represents a pointer, and it returns a new reflect.Value that represents the value that the pointer points to.
*/
func PopulateStruct(data map[string]interface{}, person interface{}) {
	// Get the reflection value of the persn interface and navigate to its underlying struct
	personValue := reflect.ValueOf(person).Elem()

	for key, val := range data {
		/*
			FieldByName method of the reflect.Value type to get the field with the specified name (key) from the personValue. The FieldByName method returns a reflect.Value representing the field.
		*/
		fmt.Println(personValue)
		personField := personValue.FieldByName(key)

		if personField.IsValid() {
			fmt.Println(personField)
			// If the field is a struct itself, recursively populate its fields
			if personField.Kind() == reflect.Struct {
				if nestedMap, ok := val.(map[string]interface{}); ok {
					newNestedStruct := reflect.New(personField.Type()).Interface()
					PopulateStruct(nestedMap, newNestedStruct)
					fmt.Println(newNestedStruct)
					personField.Set(reflect.ValueOf(newNestedStruct).Elem())

				}

			} else {
				personField.Set(reflect.ValueOf(val))
			}

		}
	}
}

func main() {
	// Sample data to populate the Person struct
	data := map[string]interface{}{
		"Name":    "Vaibhav",
		"Age":     22,
		"pincode": 250002, // This field won't be used as it doesn't exist in the struct
		"Address": map[string]interface{}{
			"City":  "Meerut",
			"State": "Uttar Pradesh",
		},
	}

	// Create a pointer to a Person struct
	var personPtr *Person = &Person{}
	// Populate the Person struct fields using the data
	PopulateStruct(data, personPtr)

	// Print the populated Person struct
	fmt.Printf("%+v\n", *personPtr)
}
