package interfaces

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func RunTypeAssert() {
	var w io.Writer = os.Stdout

	f, ok := w.(*os.File)      // success
	b, ok := w.(*bytes.Buffer) // failure

	fmt.Printf("%v, %v\n", f, ok)
	fmt.Printf("%v, %v\n", b, ok)
}
