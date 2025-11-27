package infinite

import (
	"bytes"
	"io"
)

var _ io.Reader = (*reader)(nil)

type reader struct {
	Data []byte
	pos  int
}

// NewReader that will always return data. Once all data has been read, it will continue from the start.
func NewReader(data []byte) io.Reader {
	return &reader{
		Data: bytes.Repeat(data, 16), // Align input for better performance.
	}
}

func (r *reader) Read(buf []byte) (int, error) {
	d := r.Data
	for i := 0; i < len(buf); {
		// copy from current position to end
		n := copy(buf[i:], d[r.pos:])
		i += n
		r.pos += n
		if r.pos == len(d) {
			// wrap to start
			r.pos = 0
		}
	}

	return len(buf), nil
}
