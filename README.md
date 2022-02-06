zibr
====

zipr is a command-line utility that repacks compressed ZIP files into
"stored" ZIP files (ZIP without compression) or PNG files into
uncompressed, but valid, PNG files, and then compresses the result
with brotli. These smaller files may be served by a web server in
response to requests with `Accept-Encoding: br` to save traffic.

See https://twitter.com/dchest/status/1488449100072857600


    Usage: zibr infile [outfile]
    -c int
            brotli compression level (0 - 11) (default 11)


Example:

    zibr file.zip

will produce "file.zip.br".

Example:

    zibr file.png

will produce "file.png.br".

