package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/dchest/cbrotli"
)

var (
	fCompressionLevel = flag.Int("c", 11, "brotli compression level (0 - 11)")
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Repack ZIP or PNG files with brotli\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s infile [outfile]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()
	if flag.NArg() < 1 || flag.NArg() > 2 {
		flag.Usage()
		return
	}
	if *fCompressionLevel < 0 || *fCompressionLevel > 11 {
		flag.Usage()
		return
	}
	infile := flag.Arg(0)
	filetype := strings.ToLower(filepath.Ext(infile))
	outfile := ""
	if flag.NArg() == 2 {
		outfile = flag.Arg(1)
	} else {
		outfile = infile + ".br"
	}

	outf, err := os.Create(outfile)
	if err != nil {
		log.Fatal(err)
	}
	defer outf.Close()

	bw := cbrotli.NewWriter(outf, cbrotli.WriterOptions{Quality: *fCompressionLevel})
	defer bw.Close()

	switch filetype {
	case ".zip":
		if err := RepackZip(bw, infile); err != nil {
			log.Fatal(err)
		}
	case ".png":
		if err := RepackPng(bw, infile); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("Unknown file extension (should be .zip or .png)")
	}
}
