## メモ

```go
func Area(circle Circle) float64 { ... }
func Area(rectangle Rectangle) float64 { ... }
```

Java などの言語だったら、引数が異なるため上記コードは OK。
だけど、go ではできないらしい。

## メソッド

メソッドは、タイプ（構造体とか、int など）にくっついている関数のこと。

```go
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
```

こんな感じで、宣言する。
関数名の前についている、`(r Rectangle)`がタイプ。

## インターフェース

```go
type Shape interface {
    Area() float64
}
```

これが go でインターフェースを実装する記載方法。
この`{...}`内に記載しているメソッドがあるタイプは暗黙的にインターフェースを実装したタイプとみなされる。
