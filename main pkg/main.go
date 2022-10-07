package main

import (
	"fmt"
	"testcode2"
)

func main() {
	data := testcode2.UserInput()
	output := testcode2.Palindrome(data)
	fmt.Println(output)
}
