package main

import (
	"io"
	"os"

	"github.com/dchest/pnglevel"
)

func RepackPng(w io.Writer, filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return pnglevel.Repack(w, f, 0)
}
