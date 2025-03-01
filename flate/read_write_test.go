package flate

import (
	"bytes"
	"io"
	"log"
	"strings"
	"testing"
)

func TestReadWrite(t *testing.T) {
	proverbs := []string{
		"Don't communicate by sharing memory, share memory by communicating.\n",
		"Concurrency is not parallelism.\n",
		"The bigger the interface, the weaker the abstraction.\n",
		"Documentation is for users.\n",
	}

	var r strings.Reader
	var b bytes.Buffer
	var o strings.Builder

	for _, v := range proverbs {
		r.Reset(v)
		b.Reset()
		o.Reset()

		zw := NewWriter(&b)
		if _, err := io.Copy(zw, &r); err != nil {
			log.Fatal(err)
		}
		if err := zw.Close(); err != nil {
			log.Fatal(err)
		}

		zr := NewReader(&b)
		if _, err := io.Copy(&o, zr); err != nil {
			log.Fatal(err)
		}
		if err := zr.Close(); err != nil {
			log.Fatal(err)
		}

		if v != o.String() {
			t.Error(v + o.String())
		}
	}
}
