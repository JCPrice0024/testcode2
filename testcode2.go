package main

import "fmt"

func pal(a string) bool {
	vals := []rune(a)
	back := len(vals) - 1
	for front := 0; front <= back; front++ {
		f := vals[front]
		if f >= 'A' && f <= 'Z' {
			f = 'a' + f - 'A'
		}
		if (f >= 'a' && f <= 'z') || (f >= '0' && f <= '9') {
			back = findValidBack(vals, back, front)
			b := vals[back]
			if b >= 'A' && b <= 'Z' {
				b = 'a' + b - 'A'
			}
			if f != b {
				return false
			}
			back--
		}

	}
	return true
}

//findValidBack takes an array of runes, and the current front, and back, and returns
//index of the next element that contains a valid ascii number or letter
func findValidBack(vals []rune, back int, front int) int {
	for ; back >= front; back-- {
		b := vals[back]
		if b >= 'A' && b <= 'Z' {
			b = 'a' + b - 'A'
		}
		if (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9') {
			return back
		}

	}
	return front
}
func main() {
	data := "A man, a plan, a canal: Panama"
	output := pal(data)
	fmt.Println(output)

}
