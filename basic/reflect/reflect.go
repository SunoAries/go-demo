package main

import (
	"reflect"
	"fmt"
)

func main() {
	var x float64 = 5.21
	fmt.Println("type:", reflect.TypeOf(x)) //type: float64

	v := reflect.ValueOf(x)
	fmt.Println("value:", v)         //value: 5.21
	fmt.Println("type:", v.Type())   //type: float64
	fmt.Println("kind:", v.Kind())   //kind: float64
	fmt.Println("value:", v.Float()) //value: 5.21

	fmt.Println(v.Interface())                    //5.21
	fmt.Printf("value is %1.1e\n", v.Interface()) //value is 5.2e+00
	y := v.Interface().(float64)
	fmt.Println(y) //5.21
}