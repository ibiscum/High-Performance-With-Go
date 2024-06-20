package hello_test

import (
	"fmt"
	"testing"
)

var _ = fmt.Sprint()

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = "Hello High Performance Go" // Print with no formatting
	}
}
