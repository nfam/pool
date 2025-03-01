package flate

import (
	"compress/flate"
	"io"
	"sync"
)

var readerPool sync.Pool

type Reader struct {
	io.ReadCloser
}

func NewReader(r io.Reader) *Reader {
	if v := readerPool.Get(); v != nil {
		x := v.(*Reader)
		if s, ok := x.ReadCloser.(flate.Resetter); ok {
			s.Reset(r, nil)
			return x
		}
	}
	return &Reader{flate.NewReader(r)}
}

func (r *Reader) Close() error {
	err := r.ReadCloser.Close()
	readerPool.Put(r)
	return err
}
