package flate

import (
	"compress/flate"
	"io"
	"sync"
)

var writerPool sync.Pool

type Writer struct {
	*flate.Writer
}

func NewWriter(w io.Writer) *Writer {
	if v := writerPool.Get(); v != nil {
		z := v.(*Writer)
		z.Reset(w)
		return z
	}
	z, _ := flate.NewWriter(w, flate.BestCompression)
	return &Writer{z}
}

func (r *Writer) Close() error {
	err := r.Writer.Close()
	writerPool.Put(r)
	return err
}
