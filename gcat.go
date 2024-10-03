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

	isNumbered := flag.Bool("n", false, "Number the output lines, starting at 1.")
	isNonBlankNumbered := flag.Bool("b", false, "Number the non-blank output lines, starting at 1.")
	flag.Parse()
	files := flag.Args()

	if len(files) == 0 {
		input = os.Stdin
	} else {
		readers := make([]io.Reader, len(files))
		for index, name := range files {
			if name == "-" {
				readers[index] = os.Stdin
			} else {
				file, err := os.Open(name)
				if err != nil {
					exitWithError(err)
					return
				}
				readers[index] = file
			}
		}
		input = io.MultiReader(readers...)
	}

	if *isNonBlankNumbered {
		f = formatter.NonBlankNumbered(input)
	} else if *isNumbered {
		f = formatter.Numbered(input)
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
