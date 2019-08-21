package reflection

import (
	"fmt"
	"strconv"
)

// example not using reflection

func RunSprint(x interface{}) {
	var msg string

	type stringer interface {
		String() string
	}

	switch x := x.(type) {
	case stringer:
		msg = x.String()
	case string:
		msg = x
	case int:
		msg = strconv.Itoa(x)
	case bool:
		if x {
			msg = "true"
		}
		msg = "false"
	default:
		msg = "???" // array, chan, func, map, pointer, slice, struct
	}

	fmt.Print(msg)
}
