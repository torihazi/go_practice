package main

import (
	"fmt"
	"reflect"
)

func main() {
	t := []interface{}{1, "2", 10, "11"}
	for i := 0; i < len(t); i++ {
		typ := reflect.TypeOf(t[i])
		switch typ.Kind() {
		case reflect.Int:
			fmt.Printf("%02d\n", t[i])
		case reflect.String:
			fmt.Printf("%02s\n", t[i])
		default:
			fmt.Println("エラーです")
		}
	}
}
