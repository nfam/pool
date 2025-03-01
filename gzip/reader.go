package gzip

import (
	"compress/gzip"
	"io"
	"sync"
)

var readerPool sync.Pool

type Reader struct {
	*gzip.Reader
}

func NewReader(r io.Reader) (*Reader, error) {
	if v := readerPool.Get(); v != nil {
		z := v.(*Reader)
		if err := z.Reset(r); err != nil {
			readerPool.Put(z)
			return nil, err
		}
		return z, nil
	}
	z, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}
	return &Reader{z}, nil
}

func (r *Reader) Close() error {
	err := r.Reader.Close()
	readerPool.Put(r)
	return err
}
