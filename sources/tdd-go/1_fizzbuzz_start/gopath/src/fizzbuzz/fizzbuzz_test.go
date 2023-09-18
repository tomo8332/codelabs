// fizzbuzz_test.go

package fizzbuzz_test

import (
	"testing"

	"github.com/tomo8332/codelabs/sources/tdd-go/1_fizzbuzz_start/gopath/src/fizzbuzz"
)

func TestConvert(t *testing.T) {
	tests := []struct {
		n    int    // 入力値
		want string // 期待値
	}{
		{n: 1, want: "1"},
		{n: 2, want: "2"},
		{n: 3, want: "Fizz"},
		{n: 5, want: "Buzz"},
	}

	for _, tt := range tests {
		got := fizzbuzz.Convert(tt.n)
		if got != tt.want {
			t.Errorf(`Convert(%v) = %q but want %q`, tt.n, got, tt.want)
		}
	}

}
