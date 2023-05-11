//go:build !cgo

package infinite

import "io"

var _ io.Reader = (*reader)(nil)

type reader struct {
	Data []byte
	pos  int
}

// NewReader that will always return data. Once all data has been read, it will continue from the start.
func NewReader(data []byte) io.Reader {
	return &reader{
		Data: data,
	}
}

func (r *reader) Read(buf []byte) (int, error) {
	for i := range buf {
		buf[i] = r.Data[r.pos]

		r.pos++
		if r.pos == len(r.Data) {
			r.pos = 0
		}
	}

	return len(buf), nil
}
