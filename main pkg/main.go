package main

import (
	"fmt"
	pal "testcode2"
)

func main() {
	for {
		data := pal.UserInput()
		output := pal.Palindrome(data)
		if data == "end" {
			return
		} else {
			output = pal.Palindrome(data)
		}
		fmt.Println(output)
	}
}
