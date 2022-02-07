package main

import (
	"archive/zip"
	"fmt"
	"io"
)

// RepackZip repacks a ZIP file filename into a ZIP file
// without compression and writes it to the given writer.
//
// If utf8flag is true, sets NonUTF8 flag in file headers to false.
func RepackZip(w io.Writer, filename string, utf8Flag bool) error {
	r, err := zip.OpenReader(filename)
	if err != nil {
		return err
	}
	zw := zip.NewWriter(w)

	for _, f := range r.File {
		fmt.Printf("Repacking %s\n", f.Name)
		h := f.FileHeader
		h.Method = 0
		if utf8Flag {
			h.NonUTF8 = false
		}
		fw, err := zw.CreateHeader(&h)
		if err != nil {
			return err
		}
		rc, err := f.Open()
		if err != nil {
			return err
		}
		_, err = io.Copy(fw, rc)
		if err != nil {
			rc.Close()
			return err
		}
		if err := rc.Close(); err != nil {
			return err
		}
	}
	if err := zw.Close(); err != nil {
		return err
	}
	return nil
}
