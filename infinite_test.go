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

	buf := make([]byte, 16)

	n, err := r.Read(buf)
	assert.NoError(t, err)
	assert.Equal(t, len(buf), n)
	assert.Equal(t, "0123456789ABCDEF", string(buf))

	n, err = r.Read(buf)
	assert.NoError(t, err)
	assert.Equal(t, len(buf), n)
	assert.Equal(t, "-0123456789ABCDE", string(buf))

	n, err = r.Read(buf)
	assert.NoError(t, err)
	assert.Equal(t, len(buf), n)
	assert.Equal(t, "F-0123456789ABCD", string(buf))
}

func TestReadBufferLargerThanData(t *testing.T) {
	r := NewReader([]byte("foobar"))

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

	// Buffer size must be a multiple of 128 bits (16 bytes).
	buf := make([]byte, 16)

	r.Read(buf)
	fmt.Println(string(buf))

	r.Read(buf)
	fmt.Println(string(buf))
	// Output: Hello infinity!H
	// ello infinity!He
}

// $ go test -benchmem -run=^$ -bench ^BenchmarkRead_4096$
// goos: darwin
// goarch: arm64
// pkg: github.com/jnodorp/infinite
// cpu: Apple M4 Pro
// BenchmarkRead_4096-14           10906615                92.16 ns/op     44446.22 MB/s          0 B/op          0 allocs/op
// PASS
// ok      github.com/jnodorp/infinite     1.210s
func BenchmarkRead_4096(b *testing.B) {
	r := NewReader([]byte("foobar"))

	// Use a reasonably large buffer, to get useful results. Buffer size must be a multiple of 128 bits (16 bytes).
	buf := make([]byte, 4096)

	for b.Loop() {
		n, err := r.Read(buf)
		assert.NoError(b, err)
		b.SetBytes(int64(n))
	}
}

// $ go test -benchmem -run=^$ -bench ^BenchmarkRead_8192$
// goos: darwin
// goarch: arm64
// pkg: github.com/jnodorp/infinite
// cpu: Apple M4 Pro
// BenchmarkRead_8192-14            6095868               178.4 ns/op      45922.22 MB/s          0 B/op          0 allocs/op
// PASS
// ok      github.com/jnodorp/infinite     1.303s
func BenchmarkRead_8192(b *testing.B) {
	r := NewReader([]byte("foobar"))

	// Use a reasonably large buffer, to get useful results. Buffer size must be a multiple of 128 bits (16 bytes).
	buf := make([]byte, 8192)

	for b.Loop() {
		n, err := r.Read(buf)
		assert.NoError(b, err)
		b.SetBytes(int64(n))
	}
}
