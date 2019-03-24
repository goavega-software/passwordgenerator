package passwordgen

import (
	"strings"
	"testing"
)

func TestShouldHaveCorrectLength(t *testing.T) {
	want := 8
	t.Run("Length should be 8", func(t *testing.T) {
		if got := Generate(); len(got) != want {
			t.Errorf("Generate() = %v, want %v", got, want)
		}
	})
}

func TestShouldStartWithConsonant(t *testing.T) {
	want := "aeiou0123456789"
	t.Run("Length should be 8", func(t *testing.T) {
		got := Generate()
		if i := strings.Index(string([]rune(got)[0]), want); i > -1 {
			t.Errorf("Generate() = %v, should not start with want %v", got, want)
		}
	})
}
