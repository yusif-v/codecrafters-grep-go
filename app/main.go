package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"unicode/utf8"
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
	
	if pattern == "\\d" {
		return bytes.ContainsAny(line, "0123456789"), nil
	}
	else pattern == "\\w" {
		return bytes.ContainsAny(line, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_")
	}
	
	if utf8.RuneCountInString(pattern) != 1 {
		return false, fmt.Errorf("unsupported pattern: %q", pattern)
	}

	var ok bool

	fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")

	ok = bytes.ContainsAny(line, pattern)

	return ok, nil
}
