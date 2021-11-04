package cipher

import (
	"bytes"
	"io"
)

// Rot128Reader implements io.Reader that transforms
type Rot128Reader struct{ reader io.Reader }

func NewRot128Reader(r io.Reader) (*Rot128Reader, error) {
	return &Rot128Reader{reader: r}, nil
}

func readAll(r io.Reader, capacity int64) (b []byte, err error) {
	var buf bytes.Buffer
	// If the buffer overflows, we will get bytes.ErrTooLarge.
	// Return that as an error. Any other panic remains.
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()
	if int64(int(capacity)) == capacity {
		buf.Grow(int(capacity))
	}
	_, err = buf.ReadFrom(r)
	return buf.Bytes(), err
}

func (r *Rot128Reader) Read(p []byte) (int, error) {
	if n, err := r.reader.Read(p); err != nil {
		return n, err
	} else {
		rot128(p[:n])
		return n, nil
	}
}

type Rot128Writer struct {
	writer io.Writer
	buffer []byte // not thread-safe
}

func NewRot128Writer(w io.Writer) (*Rot128Writer, error) {
	return &Rot128Writer{
		writer: w,
		buffer: make([]byte, 4096, 4096),
	}, nil
}

func (w *Rot128Writer) Write(p []byte) (int, error) {
	n := copy(w.buffer, p)
	rot128(w.buffer[:n])
	return w.writer.Write(w.buffer[:n])
}

func rot128(buf []byte) {
	for idx := range buf {
		buf[idx] += 128
	}
}
