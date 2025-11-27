# infinite

Infinite provides an [`io.Reader`](https://pkg.go.dev/io#Reader) that never returns EOF. Instead, a pattern is repeated
forever.

I needed this for performance testing of a streaming JSON parser.

## Usage

```go
r := infinite.NewReader([]byte("Hello infinity!"))

buf := make([]byte, 16)

r.Read(buf)
fmt.Println(string(buf))

r.Read(buf)
fmt.Println(string(buf))

// Output:
// Hello infinity!H
// ello infinity!He
```
