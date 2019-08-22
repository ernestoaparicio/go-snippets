package reflection

import (
	"fmt"
	"reflect"
)

func RunSprintReflect(x interface{}) {
	t := reflect.ValueOf(x)

	fmt.Println(t)
}
