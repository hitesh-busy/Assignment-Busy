// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"reflect"
)

var ans []interface{}

func helper(arr interface{}) {

	valueOf := reflect.ValueOf(arr)
	switch valueOf.Kind() {
	case reflect.Slice:
		for _, v := range arr.([]interface{}) {
			helper(v)
		}
	case reflect.Interface:
		for _, v := range arr.([]interface{}) {
			ans = append(ans, v)
		}
	case reflect.Int32:
		//trick to tacle 'a' like chars
		ans = append(ans, string(rune(arr.(int32))))
		fmt.Println(arr, valueOf)
	default:
		ans = append(ans, arr)
	}
}
func merger(arr1 interface{}, arr2 interface{}) error {

	if arr1 == nil && arr2 == nil {
		return nil
	}

	if arr1 == nil {
		helper(arr2)
	}
	if arr2 == nil {
		helper(arr1)
	}

	helper(arr1)
	helper(arr2)
	return nil

}
func main() {
	//var arr1 interface{}
	// arr1 = []interface{}{1, "bcd", 3, 4.23, 5}
	// arr1 = true
	//arr1 = 325.234
	arr1 := []interface{}{[]interface{}{'9', 'f', 'a', 2}}
	var arr2 interface{} = []interface{}{"st", 6, true, 7, 3.14}

	err := merger(arr1, arr2)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ans is ", ans)
}
