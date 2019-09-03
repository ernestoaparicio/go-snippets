package concurrent

import (
	"fmt"
	"runtime"
)

type systemInfo struct{}

func (s systemInfo) getMaxProcs() int {
	return runtime.GOMAXPROCS(0)
}

func RunPrintMaxProcs() {
	s := systemInfo{}
	fmt.Println("GOMAXPROCS:", s.getMaxProcs())
}
