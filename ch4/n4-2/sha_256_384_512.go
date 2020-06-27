package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {
	var str string

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		str = input.Text()
	}

	switch {
	case len(os.Args) == 1 || os.Args[1] == "sha256":
		r := sha256.Sum256([]byte(str))
		fmt.Printf("SHA256: %x", r)
	case os.Args[1] == "sha384":
		r := sha512.Sum384([]byte(str))
		fmt.Printf("SHA384: %x", r)
	case os.Args[1] == "sha512":
		r := sha512.Sum512([]byte(str))
		fmt.Printf("SHA512: %x", r)
	}
}
