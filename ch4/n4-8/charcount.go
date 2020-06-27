// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	var letterCounts, numberCounts int
	invalid := 0 // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsLetter(r) {
			letterCounts++
			continue
		}
		if unicode.IsNumber(r) {
			numberCounts++
			continue
		}
	}
	fmt.Printf("letter\t%d\n", letterCounts)
	fmt.Printf("number\t%d\n", numberCounts)

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
