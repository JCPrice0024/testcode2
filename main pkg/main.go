package main

import (
	"fmt"
	"strings"
	"testcode2"
)

func main() {
	for {
		data := testcode2.UserInput()
		strings.TrimSpace(data)
		output := testcode2.Palindrome(data)
		if data == "end" {
			return
		} else {
			output = testcode2.Palindrome(data)
		}
		fmt.Println(output)
	}

	//output := testcode2.Palindrome(data)
	//fmt.Println(output)
}
