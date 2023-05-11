package infinite_test

import (
	"fmt"
	"strings"
	"testing"

	. "github.com/jnodorp/infinite"
	"github.com/stretchr/testify/assert"
)

func TestReadBufferSmallerThanData(t *testing.T) {
	r := NewReader([]byte(string("0123456789ABCDEF-")))

	// Buffer size must be a multiple of 256 bits (32 bytes).
	buf := make([]byte, 32)

	n, err := r.Read(buf)
	assert.NoError(t, err)
	assert.Equal(t, len(buf), n)
	assert.Equal(t, "0123456789ABCDEF-0123456789ABCDE", string(buf))

	n, err = r.Read(buf)
	assert.NoError(t, err)
	assert.Equal(t, len(buf), n)
	assert.Equal(t, "F-0123456789ABCDEF-0123456789ABC", string(buf))

	n, err = r.Read(buf)
	assert.NoError(t, err)
	assert.Equal(t, len(buf), n)
	assert.Equal(t, "DEF-0123456789ABCDEF-0123456789A", string(buf))
}

func TestReadBufferLargerThanData(t *testing.T) {
	r := NewReader([]byte("foobar"))

	// Buffer size must be a multiple of 256 bits (32 bytes).
	buf := make([]byte, 64)

	n, err := r.Read(buf)
	assert.NoError(t, err)
	assert.Equal(t, len(buf), n)
	assert.Equal(t, strings.Repeat("foobar", 12)[:64], string(buf))

	n, err = r.Read(buf)
	assert.NoError(t, err)
	assert.Equal(t, len(buf), n)
	assert.Equal(t, strings.Repeat("foobar", 12)[4:68], string(buf))

	n, err = r.Read(buf)
	assert.NoError(t, err)
	assert.Equal(t, len(buf), n)
	assert.Equal(t, strings.Repeat("foobar", 12)[2:66], string(buf))

	n, err = r.Read(buf)
	assert.NoError(t, err)
	assert.Equal(t, len(buf), n)
	assert.Equal(t, strings.Repeat("foobar", 12)[:64], string(buf))
}

func ExampleNewReader() {
	r := NewReader([]byte("Hello infinity!"))

	// Buffer size must be a multiple of 256 bits (32 bytes).
	buf := make([]byte, 32)

	r.Read(buf)
	fmt.Printf(string(buf) + "\n")

	r.Read(buf)
	fmt.Printf(string(buf) + "\n")
	// Output: Hello infinity!Hello infinity!He
	// llo infinity!Hello infinity!Hell
}

// BenchmarkRead heavily depends on whether cgo is enabled!
//
// $ CGO_ENABLED=0 go test -benchmem -run=^$ -bench ^BenchmarkRead_4096$
// goos: darwin
// goarch: amd64
// pkg: github.com/jnodorp/infinite
// cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
// BenchmarkRead_4096-16    	  249250	      4082 ns/op	1003.49 MB/s	       0 B/op	       0 allocs/op
// PASS
// ok  	github.com/jnodorp/infinite	1.343s
//
// $ CGO_ENABLED=1 go test -benchmem -run=^$ -bench ^BenchmarkRead_4096$
// goos: darwin
// goarch: amd64
// pkg: github.com/jnodorp/infinite
// cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
// BenchmarkRead_4096-16    	 4713326	       250.4 ns/op	16355.63 MB/s	      24 B/op	       1 allocs/op
// PASS
// ok  	github.com/jnodorp/infinite	1.548s
func BenchmarkRead_4096(b *testing.B) {
	r := NewReader([]byte("foobar"))

	// Use a reasonably large buffer, to get useful results. Buffer size must be a multiple of 128 bits (16 bytes).
	buf := make([]byte, 4096)

	for i := 0; i < b.N; i++ {
		n, err := r.Read(buf)
		assert.NoError(b, err)
		b.SetBytes(int64(n))
	}
}

// BenchmarkRead heavily depends on whether cgo is enabled!
//
// $ CGO_ENABLED=0 go test -benchmem -run=^$ -bench ^BenchmarkRead_8192$
// goos: darwin
// goarch: amd64
// pkg: github.com/jnodorp/infinite
// cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
// BenchmarkRead_8192-16    	  141156	      8320 ns/op	 984.60 MB/s	       0 B/op	       0 allocs/op
// PASS
// ok  	github.com/jnodorp/infinite	1.372s
//
// $ CGO_ENABLED=1 go test -benchmem -run=^$ -bench ^BenchmarkRead_8192$
// goos: darwin
// goarch: amd64
// pkg: github.com/jnodorp/infinite
// cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
// BenchmarkRead_8192-16    	 3105466	       377.6 ns/op	21692.77 MB/s	      24 B/op	       1 allocs/op
// PASS
// ok  	github.com/jnodorp/infinite	1.681s
func BenchmarkRead_8192(b *testing.B) {
	r := NewReader([]byte("foobar"))

	// Use a reasonably large buffer, to get useful results. Buffer size must be a multiple of 128 bits (16 bytes).
	buf := make([]byte, 8192)

	for i := 0; i < b.N; i++ {
		n, err := r.Read(buf)
		assert.NoError(b, err)
		b.SetBytes(int64(n))
	}
}
