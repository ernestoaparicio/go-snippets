package interfaces

import (
	"fmt"
	"sort"
)

type StringSlice []string

func (p StringSlice) Len() int {
	return len(p)
}

func (p StringSlice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p StringSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func RunSortExample1() {
	names := []string{"John", "Paul", "George", "Ringo"}
	sort.Sort(StringSlice(names))
	//sort.Strings(names)

	fmt.Println(names)
}
