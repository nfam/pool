package buffer

import (
	"bytes"
	"sync"
)

var pool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

type Buffer struct {
	*bytes.Buffer
}

func Get() Buffer {
	b := pool.Get().(*bytes.Buffer)
	b.Reset()
	return Buffer{b}
}

func (b *Buffer) Close() error {
	pool.Put(b.Buffer)
	b.Buffer = nil
	return nil
}
