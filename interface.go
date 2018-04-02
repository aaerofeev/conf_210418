package main

import (
	"os"
	"io"
)

func A(w os.File) struct{} {
	return struct{}{}
}

func B(w io.Writer) struct{} {
	return struct{}{}
}