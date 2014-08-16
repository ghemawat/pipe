package stream_test

import (
	"github.com/ghemawat/stream"

	"fmt"
	"os"
	"testing"
)

func BenchmarkSingle(b *testing.B) {
	stream.Run(stream.Repeat("", b.N))
}

func BenchmarkFive(b *testing.B) {
	f := stream.FilterFunc(func(arg stream.Arg) error {
		for s := range arg.In {
			arg.Out <- s
		}
		return nil
	})
	stream.Run(stream.Repeat("", b.N), f, f, f, f)
}

func BenchmarkWL(b *testing.B) {
	f, err := os.Create("/dev/null")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	stream.Run(
		stream.Repeat("hello", b.N),
		stream.WriteLines(f),
	)
}

func BenchmarkSample(b *testing.B) {
	stream.Run(
		stream.Repeat("hello", b.N),
		stream.Sample(10),
	)
}

func BenchmarkSort(b *testing.B) {
	stream.Run(
		stream.Repeat("hello", b.N),
		stream.Sort(),
	)
}
