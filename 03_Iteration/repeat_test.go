package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
    repeated := Repeat("a", 5)
    expected := "aaaaa"

    if repeated != expected {
        t.Errorf("expected %q but got %q", expected, repeated)
    }
}

func BenchmarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repeat("a", 5)
    }
}

func ExampleRepeat() {
    repeated := Repeat("a", 4)
    fmt.Println(repeated)
    // Output: aaaa
}

func TestGetBigger(t *testing.T) {
    bigger := GetBigger("a", "b")
    expected := "b"

    if bigger != expected {
        t.Errorf("expected %q but got %q", expected, bigger)
    }
}