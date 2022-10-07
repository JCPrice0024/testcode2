package main

import (
	"fmt"
	"testcode2"
)

func main() {
	for {
		data := testcode2.UserInput()
		output := testcode2.Palindrome(data)
		if data == "end" {
			return
		} else {
			output = testcode2.Palindrome(data)
		}
		fmt.Println(output)
	}
}
