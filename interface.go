package main

type Formatter interface {
	ReadLine() ([]byte, error)
}
