package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/dchest/cbrotli"
)

var (
	fCompressionLevel = flag.Int("c", 11, "brotli compression level (0 - 11)")
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s infile.zip [outfile]\n", os.Args[0])
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
	outfile := ""
	if flag.NArg() == 2 {
		outfile = flag.Arg(2)
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
	if err := RepackZip(bw, infile); err != nil {
		log.Fatal(err)
	}
}
