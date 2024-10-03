package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/simp7/gcat/formatter"
)

func exitWithError(err error) {
	fmt.Printf("Error occured: %s\n", err)
	os.Exit(1)
}

func main() {
	var input io.Reader
	var output io.Writer
	var f Formatter

	var line []byte
	var err error

	output = os.Stdout

	args := os.Args
	isNumbered := flag.Bool("n", false, "Number the output lines, starting at 1.")
	isNonBlankNumbered := flag.Bool("b", false, "Number the non-blank output lines, starting at 1.")

	if len(args) == 1 {
		input = os.Stdin
	} else {
		if args[1] == "-" {
			input = os.Stdin
		} else {
			file, err := os.Open(args[1])
			if err != nil {
				exitWithError(err)
				return
			}
			input = file
		}
	}

	if *isNumbered {
	} else if *isNonBlankNumbered {
	} else {
		f = formatter.Standard(input)
	}

	for {
		if line, err = f.ReadLine(); err != nil {
			break
		}

		output.Write(append(line, '\n'))
	}

	if err != io.EOF {
		exitWithError(err)
		return
	}
}
