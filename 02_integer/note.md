## Testable Examples

test ファイルの中に例として記述する関数のこと。

```go
func ExampleAdd() {
    sum := Add(1, 5)
    fmt.Println(sum)
    // Output: 6
}
```

これを記載すると、このコードが godoc 内のドキュメントに表示される。
`go test -v`を実行すると、このコードもテストされる。
なので。README などに記載するよりも、整合性が取れる。
(// Output 6 がないと実行されない。)
