package reflection

import (
	"fmt"
	"reflect"
)

func double(x interface{}) {
	//d := reflect.ValueOf(&x).Elem() // no need to get pointer since x is already a pointer
	d := reflect.ValueOf(x).Elem()
	d.Set(reflect.ValueOf(4))
}

func RunDouble(x interface{}) {
	double(&x) // pass a pointer so it's addressable
	fmt.Println(x)
}
