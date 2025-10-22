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

	// Buffer size must be a multiple of 128 bits (16 bytes).
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

	// Buffer size must be a multiple of 128 bits (16 bytes).
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
// CGO_ENABLED=0 go test -benchmem -run=^$ -bench ^BenchmarkRead_4096$
// goos: darwin
// goarch: arm64
// pkg: github.com/jnodorp/infinite
// cpu: Apple M4 Pro
// BenchmarkRead_4096-14             882296              1314 ns/op        3117.73 MB/s           0 B/op          0 allocs/op
// PASS
// ok      github.com/jnodorp/infinite     2.069s
//
// $ CGO_ENABLED=1 go test -benchmem -run=^$ -bench ^BenchmarkRead_4096$
// goos: darwin
// goarch: amd64
// pkg: github.com/jnodorp/infinite
// cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
// BenchmarkRead_4096-16    	 4713326	       250.4 ns/op	16355.63 MB/s	      24 B/op	       1 allocs/op
// PASS
// ok  	github.com/jnodorp/infinite	1.548s
//
// CGO_ENABLED=1 go test -benchmem -run=^$ -bench ^BenchmarkRead_4096$
// goos: darwin
// goarch: arm64
// pkg: github.com/jnodorp/infinite
// cpu: Apple M4 Pro
// BenchmarkRead_4096-14            5559720               204.7 ns/op      20006.60 MB/s          0 B/op          0 allocs/op
// PASS
// ok      github.com/jnodorp/infinite     1.518s
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
// CGO_ENABLED=0 go test -benchmem -run=^$ -bench ^BenchmarkRead_8192$
// goos: darwin
// goarch: arm64
// pkg: github.com/jnodorp/infinite
// cpu: Apple M4 Pro
// BenchmarkRead_8192-14             430222              2936 ns/op        2789.91 MB/s           0 B/op          0 allocs/op
// PASS
// ok      github.com/jnodorp/infinite     2.302s
//
// $ CGO_ENABLED=1 go test -benchmem -run=^$ -bench ^BenchmarkRead_8192$
// goos: darwin
// goarch: amd64
// pkg: github.com/jnodorp/infinite
// cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
// BenchmarkRead_8192-16    	 3105466	       377.6 ns/op	21692.77 MB/s	      24 B/op	       1 allocs/op
// PASS
// ok  	github.com/jnodorp/infinite	1.681s
//
// CGO_ENABLED=1 go test -benchmem -run=^$ -bench ^BenchmarkRead_8192$
// goos: darwin
// goarch: arm64
// pkg: github.com/jnodorp/infinite
// cpu: Apple M4 Pro
// BenchmarkRead_8192-14            3101902               398.6 ns/op      20551.33 MB/s          0 B/op          0 allocs/op
// PASS
// ok      github.com/jnodorp/infinite     1.823s
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
