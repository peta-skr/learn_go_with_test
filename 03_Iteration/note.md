## ベンチマーク

```go
func BenchmarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repeat("a")
    }
}
```

ベンチマークのコードはテストのコードと似ている。
このコードが実行されると、`b.N`回実行され、かかる時間を測定する。
`go test -bench=.`(windows Powershell を使用している場合は、`go test -bench="."`)を実行すると、ベンチマークを実行できる。
