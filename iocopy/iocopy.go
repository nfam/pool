package iocopy

import (
	"io"
	"sync"
)

var pool = sync.Pool{
	New: func() any {
		buf := make([]byte, 32*1024)
		return &buf
	},
}

func Copy(dst io.Writer, src io.Reader) (written int64, err error) {
	pbuf := pool.Get().(*[]byte)
	defer pool.Put(pbuf)
	return io.CopyBuffer(dst, src, *pbuf)
}
