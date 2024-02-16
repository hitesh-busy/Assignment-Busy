// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"reflect"
)

func RemoveKey(field string, m map[string]interface{}) {

	//first check if it exists
	if _, ok := m[field]; ok {
		delete(m, field)
	}

	//like in the last Q traverse the map inwards and use recursion to delete the field
	for _, v := range m {
		//valValue is a reflect.Value representing the value of an element in the slice or map.
		valValue := reflect.ValueOf(v)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			for i := 0; i < valValue.Len(); i++ {

				//Index(i) is a method of reflect.Value that returns the i-th element of the array, slice, or string represented by valValue.
				//Interface() is a method used to convert the value represented by reflect.Value to an interface type, effectively losing type information.
				/*
					 the RemoveKey function is designed to work with a map of type map[string]interface{}.
					Even though the original type might have been map[string]interface{}, when you use Interface() on an element from a slice of interfaces,
					the compiler sees it as an interface{}. To use it as a map[string]interface{}, you need to perform a type assertion.

					Type After .Interface(): interface{}
					The .Interface() method in Go's reflect package returns an interface{} type. This means that it converts the underlying value represented by a reflect.Value to an empty interface (interface{}), which is a type capable of holding values of any type. When you use .Interface(), you lose the specific type information, and the value is seen as an empty interface. To work with it as its original type, you typically need to perform a type assertion.

				*/
				if _, ok := valValue.Index(i).Interface().(map[string]interface{}); ok {
					RemoveKey(field, valValue.Index(i).Interface().(map[string]interface{}))
				}
			}
		case reflect.Map:
			RemoveKey(field, valValue.Interface().(map[string]interface{}))

		}
	}
}
func main() {
	var m = map[string]interface{}{
		"Name": "Hitesh Test Mame",
		"DOB":  14 - 06 - 2001,
		"city": "Delhi",
		"pin":  110023,
		// field named "Address" and assigns it a slice ([]interface{}). The square brackets [] indicate that it's a slice, and interface{} allows elements of any data type to be stored in the slice.
		"NewTest": map[string]interface{}{
			"street":  "Testtt",
			"plot_no": 96,
			"city":    "Example city",
			"pin":     633078,
		},
		"Address": []interface{}{
			//inside the slice, there are two elements, each represented by a map.

			map[string]interface{}{
				"street":  "Ashirvad chowk",
				"plot_no": 26,
				"city":    "Dwarka",
				"pin":     110078,
			},
			map[string]interface{}{
				"street":  "Lovely Chowk",
				"plot_no": 26,
				"city":    "London",
				"pin":     923478,
			},
		},
		"Salary":      800000,
		"Designation": "Developer",
	}
	//
	RemoveKey("city", m)
	fmt.Println(m)

}
