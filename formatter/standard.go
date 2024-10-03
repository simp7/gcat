package formatter

import (
	"bufio"
	"io"
)

type standard struct {
	reader *bufio.Reader
}

func Standard(input io.Reader) *standard {
	s := new(standard)
	s.reader = bufio.NewReader(input)
	return s
}

func (s *standard) ReadLine() ([]byte, error) {
	line, _, err := s.reader.ReadLine()
	return line, err
}
