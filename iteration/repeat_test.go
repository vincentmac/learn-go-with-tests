package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	assertCorrect := func(t *testing.T, actual, expected string) {
		t.Helper()

		if actual != expected {
			t.Errorf("expected %q but got %q", expected, actual)
		}
	}

	t.Run("Repeat a", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertCorrect(t, repeated, expected)
	})

	t.Run("Repeat b", func(t *testing.T) {
		repeated := Repeat("b", 4)
		expected := "bbbb"
		assertCorrect(t, repeated, expected)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("z", 6)
	fmt.Println(repeated)
	// Output: zzzzzz
}
