package main

import "fmt"

// func pal takes an array of runes and determines whether or not
// their value is a palindrome
func pal(a string) bool {
	vals := []rune(a)
	back := len(vals) - 1
	for front := 0; front <= back; front++ {
		f := vals[front]
		b := vals[back]
		front = findValidFront(vals, front, back)
		back = findValidBack(vals, back, front)
		if f != b {
			return false
		}
		back--
	}
	return true
}

// frontLowerCaseAscii takes an array of runes and the current front
// and returns the lowercase ascii value of the rune
func frontLowerCaseAscii(vals []rune, front int, back int) int {
	for f := vals[front]; f >= 'A' && f <= 'Z'; {
		f = 'a' + f - 'A'
	}
	return back
}

// backLowerCaseAscii takes an array of runes and the current back
// and returns the lowercase ascii value of the rune
func backLowerCaseAscii(vals []rune, back int, front int) int {

	for b := vals[back]; b >= 'A' && b <= 'Z'; {
		b = 'a' + b - 'A'
	}
	return front
}

// findValidBack takes an array of runes, and the current front, and back, and returns
// index of the next element that contains a valid ascii number or letter
func findValidFront(vals []rune, front int, back int) int {
	for ; front <= back; front++ {
		f := vals[front]
		front = frontLowerCaseAscii(vals, front, back)

		if (f >= 'a' && f <= 'z') || (f >= '0' && f <= '9') {
			return front
		}
	}
	return back
}

// findValidBack takes an array of runes, and the current front, and back, and returns
// index of the next element that contains a valid ascii number or letter
func findValidBack(vals []rune, back int, front int) int {
	for ; back >= front; back-- {
		b := vals[back]
		back = backLowerCaseAscii(vals, back, front)

		if (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9') {
			return back
		}

	}
	return front
}
func main() {
	data := "a man, a plan, a canal: Panama"
	output := pal(data)
	fmt.Println(output)

}
