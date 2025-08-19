package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

var _ = bytes.ContainsAny

func main() {
	if len(os.Args) < 3 || os.Args[1] != "-E" {
		fmt.Fprintf(os.Stderr, "usage: mygrep -E <pattern>\n")
		os.Exit(2)
	}

	pattern := os.Args[2]

	line, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: read input text: %v\n", err)
		os.Exit(2)
	}

	ok, err := matchLine(line, pattern)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(2)
	}

	if !ok {
		os.Exit(1)
	}
}

func matchLine(line []byte, pattern string) (bool, error) {
	fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")

	if pattern == "\\d" {
		return bytes.ContainsAny(line, "0123456789"), nil
	} else if pattern == "\\w" {
		return bytes.ContainsAny(line, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"), nil
	}

	if len(pattern) != 1 {
		return false, fmt.Errorf("unsupported pattern: %q", pattern)
	}

	return bytes.ContainsAny(line, pattern), nil
}