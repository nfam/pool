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

func (b *Buffer) IsZero() bool {
	return b.Buffer == nil
}

func Get() Buffer {
	b := pool.Get().(*bytes.Buffer)
	b.Reset()
	return Buffer{b}
}

func (b *Buffer) Close() error {
	if b != nil && b.Buffer != nil {
		pool.Put(b.Buffer)
		b.Buffer = nil
	}
	return nil
}
