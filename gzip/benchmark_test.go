package gzip

import (
	"compress/gzip"
	"io"
	"testing"
)

func BenchmarkWriters(b *testing.B) {
	w := io.Discard
	d := make([]byte, 1024*1024)
	for n := 0; n < b.N; n++ {
		z, _ := gzip.NewWriterLevel(w, gzip.BestCompression)
		z.Write(d)
		z.Close()
	}
}

func BenchmarkPoolWriters(b *testing.B) {
	w := io.Discard
	d := make([]byte, 1024*1024)
	for n := 0; n < b.N; n++ {
		z := NewWriter(w)
		z.Write(d)
		z.Close()
	}
}
