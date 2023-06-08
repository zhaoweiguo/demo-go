
- 包装器模式(装饰器模式)
- 对输入参数的类型的包装，并在不改变被包装类型（输入参数类型）的定义的情况下，返回具备新功能特性的、实现相同接口类型的新类型。

# 示例

```go
// $GOROOT/src/io/io.go
func LimitReader(r Reader, n int64) Reader { return &LimitedReader{r, n} }
type LimitedReader struct {
    R Reader // underlying reader
    N int64  // max bytes remaining
}
func (l *LimitedReader) Read(p []byte) (n int, err error) {
    // ... ...
}
```
- 使用
```go

func main() {
	r := strings.NewReader("hello, gopher!\n")
	lr := io.LimitReader(r, 4)
	if _, err := io.Copy(os.Stdout, lr); err != nil {
		log.Fatal(err)
	}
}

```

