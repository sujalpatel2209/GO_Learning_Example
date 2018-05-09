package main

import (
	"fmt"
	"strings"
)

func main() {

	sampleText := "Sujal Patel"

	fmt.Println(strings.Contains(sampleText, "Pa"))

	fmt.Println(strings.Index(sampleText, "Pa"))

	fmt.Println(strings.Count(sampleText, "a"))

	fmt.Println(strings.Replace(sampleText, "a","o",3))

}
