package main

import (
	"fmt"
	"reflect"
)

func SetKeyValue(field string, m map[string]interface{}, val interface{}) {
	//idea is to find the outermost field first and set it
	_, ok := m[field]
	if ok {
		m[field] = val
	}

	//now traverse inwards into the map
	for _, v := range m {
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			//to make the slice iterable we have to convert in into a slice of interface
			for _, vj := range v.([]interface{}) {
				SetKeyValue(field, vj.(map[string]interface{}), val)
			}
		case reflect.Map:
			//in case of map simply send and convert into map, the current 'v' of m
			SetKeyValue(field, v.(map[string]interface{}), val)

		}

	}
}

func main() {
	var m = map[string]interface{}{
		"Name": "Test Name",
		"DOB":  12 - 11 - 2004,
		"city": "Delhi",
		"pin":  110043,
		// field named "Address" and assigns it a slice ([]interface{}). The square brackets [] indicate that it's a slice, and interface{} allows elements of any data type to be stored in the slice.
		"NewTest": map[string]interface{}{
			"street":  "Testtt",
			"plot_no": 96,
			"city":    "Example city",
			"pin":     633078,
		},
		"Address": []interface{}{
			//nside the slice, there are two elements, each represented by a map.

			map[string]interface{}{
				"street":  "sabka  chowk",
				"plot_no": 26,
				"city":    "Dwarka",
				"pin":     110078,
			},
			map[string]interface{}{
				"street":  "Heroine Chowk",
				"plot_no": 26,
				"city":    "London",
				"pin":     923478,
			},
		},
		"Salary":      800000,
		"Designation": "Developer",
	}
	var val interface{}
	val = "New York"

	SetKeyValue("city", m, val)
	fmt.Println(m)

}
