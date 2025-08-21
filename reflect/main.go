package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func reflectType(i interface{}) {
	v := reflect.ValueOf(i)
	t := reflect.TypeOf(i)
	fmt.Println("Type:", t.Name())
	for j := 0; j < v.NumField(); j++ {
		field := v.Field(j)
		fieldType := t.Field(j)
		fmt.Printf("Field %d: %s (%s) = %v\n", j, fieldType.Name, fieldType.Type, field.Interface())
	}
}
func main() {
	p := Person{Name: "John", Age: 30}
	reflectType(p)
}
