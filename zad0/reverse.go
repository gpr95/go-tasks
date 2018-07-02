package main

import (
    "bytes"
    "fmt"
)

func reverse(text string) (string) {
	var result bytes.Buffer
	for i := len(text) - 1; i >= int((len(text) - 1)/2); i-- {
		result.WriteString(string([]rune(text)[i] + string([]rune(text)[len(text) - i: i - 1] + string([]rune(text)[len(text) - i - 1]))
	}

	fmt.Printf(result.String())
	return result.String()
}


func main() {
	fmt.Printf("Hello, world.\n")
	reverse("teksanfsao")
}
