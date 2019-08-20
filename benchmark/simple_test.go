package benchmark

import (
	"github.com/ernestoaparicio/go-snippets/concurrent"
	"testing"
)

func BenchmarkTestUrls(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concurrent.RunFibonacci()
	}
}
