package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	t := 0.0
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		f, err := strconv.ParseFloat(s.Text(), 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", s.Text(), err)
			continue
		}
		t += f
	}
	fmt.Printf("%.2f\n", t)
}
