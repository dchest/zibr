zibr
====

zipr is a command-line utility that repacks a compressed ZIP file into
a "stored" ZIP file (ZIP without compression), and then compresses it with brotli.

    Usage: zibr infile.zip [outfile]
    -c int
            brotli compression level (0 - 11) (default 11)


Example:

    zibr file.zip

will produce "file.zip.br".


See https://twitter.com/dchest/status/1488449100072857600
