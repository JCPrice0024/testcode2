// package testcode2 allows you to input an array of runes and detetermine whether or not
// their value is a plaindrome. It does this using a number of functions that
// take the value of the array and compare the front and back ends.
package testcode2

import (
	"bufio"
	"os"
)

// UserInput takes a user made string and puts it through pal to determine
// whether or not it is a palindrome
func UserInput() string {
	var potentialPal *bufio.Reader = bufio.NewReader(os.Stdin)
	word, _ := potentialPal.ReadString('\n')
	return word
}

// func pal takes an array of runes and determines whether or not
// their value is a palindrome
func Palindrome(a string) bool {
	vals := []rune(a)
	back := len(vals) - 1
	for front := 0; front <= back; front++ {
		f := ' '
		b := ' '
		front, f = findValidFront(vals, front, back)
		back, b = findValidBack(vals, back, front)
		if f != b {
			return false
		}
		back--
	}
	return true
}

// lowerCaseAscii takes an array of runes and the current front
// and returns the lowercase ascii value of the rune
func lowerCaseAscii(val rune) rune {
	if val >= 'A' && val <= 'Z' {
		val = 'a' + val - 'A'
	}
	return val
}

// findValidFront takes an array of runes, and the current front, and back, and returns
// index of the next element that contains a valid ascii number or letter
func findValidFront(vals []rune, front int, back int) (int, rune) {
	for ; front <= back; front++ {
		f := lowerCaseAscii(vals[front])

		if (f >= 'a' && f <= 'z') || (f >= '0' && f <= '9') {
			return front, f
		}
	}
	return back, vals[back]
}

// findValidBack takes an array of runes, and the current front, and back, and returns
// index of the next element that contains a valid ascii number or letter
func findValidBack(vals []rune, back int, front int) (int, rune) {
	for ; back >= front; back-- {
		b := lowerCaseAscii(vals[back])

		if (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9') {
			return back, b
		}

	}
	return front, vals[front]
}
