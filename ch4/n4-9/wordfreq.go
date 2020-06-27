package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	for in.Scan() {
		counts[in.Text()]++
	}

	fmt.Printf("word\tcount\n")
	for word, count := range counts {
		fmt.Printf("%s\t%d\n", word, count)
	}
}
