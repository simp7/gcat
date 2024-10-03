package formatter

import (
	"bufio"
	"fmt"
	"io"
)

type numbered struct {
	reader     bufio.Reader
	lineNumber int
}

func Numbered(input io.Reader) *numbered {
	n := new(numbered)
	n.reader = *bufio.NewReader(input)
	n.lineNumber = 1
	return n
}

func (n *numbered) ReadLine() ([]byte, error) {
	line, _, err := n.reader.ReadLine()
	if err != nil {
		return nil, err
	}

	text := fmt.Sprintf("%6d\t%s", n.lineNumber, string(line))
	n.lineNumber++

	return []byte(text), nil
}
