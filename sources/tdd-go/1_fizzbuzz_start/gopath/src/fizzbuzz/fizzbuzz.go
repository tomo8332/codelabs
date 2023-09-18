// fizzbuzz.go

package fizzbuzz

import "strconv"

func Convert(n int) string {
	// これを追加
	if n%5 == 0 {
		return "Buzz"
	}

	if n%3 == 0 {
		return "Fizz"
	}
	return strconv.Itoa(n)
}
