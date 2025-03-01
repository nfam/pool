package flate

import (
	"compress/flate"
	"io"
	"testing"
)

func BenchmarkWriters(b *testing.B) {
	w := io.Discard
	d := make([]byte, 1024*1024)
	for n := 0; n < b.N; n++ {
		z, _ := flate.NewWriter(w, flate.BestCompression)
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
