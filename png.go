package main

import (
	"io"
	"os"

	"github.com/dchest/pnglevel"
)

// RepackRng repacks a PNG file filename into a PNG with
// zero-level gzip compression. The result is a valid
// PNG file, but larger.
func RepackPng(w io.Writer, filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return pnglevel.Repack(w, f, 0)
}
