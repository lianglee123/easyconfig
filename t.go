package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string `config:"default:joy"`
	//Age  int	`config:"default: 24"`
}

func main() {
	var p interface{} = &Person{Name: "joy"}
	ptrV := reflect.ValueOf(p)
	ptrT := reflect.TypeOf(p)
	fmt.Printf("%v %v\n", ptrV, ptrT)
	fmt.Printf("%v \n", p)
	//strType := ptrT.Elem()
	strValue := ptrV.Elem()

	//for i := 0; i < strType.NumField(); i++ {
	//	f := strType.Field(i)
	//	fmt.Printf("%v \n", f)
	//}
	for i := 0; i < strValue.NumField(); i++ {
		fieldValue := strValue.Field(i)
		fmt.Printf("%v \n", fieldValue)
		fieldValue.SetString("gggg")
		fieldValue.Set()
		v := reflect.New(fieldValue.Type())
		v.Interface()
	}
	fmt.Printf("%v", p)

}
