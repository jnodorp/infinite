//go:build cgo

package infinite

// #cgo CFLAGS: -march=native
// #include "reader.h"
import "C"
import (
	"bytes"
	"io"
	"unsafe"
)

type reader_t C.reader_t

var _ io.Reader = (*reader_t)(nil)

// NewReader that will always return data. Once all data has been read, it will continue from the start.
func NewReader(data []byte) io.Reader {
	data = bytes.Repeat(data, 16)
	return (*reader_t)(C.new_reader(C.CBytes(data), C.ulong(len(data))))
}

func (r *reader_t) Read(buf []byte) (int, error) {
	return int(C._read((*C.reader_t)(r), unsafe.Pointer(&buf[0]), C.ulong(len(buf)))), nil
}
