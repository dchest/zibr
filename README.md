zibr
====

zipr is a command-line utility that repacks a compressed ZIP or PNG file into
a "stored" ZIP file (ZIP without compression) or a PNG without compression,
and then compresses the result with brotli.

    Usage: zibr infile [outfile]
    -c int
            brotli compression level (0 - 11) (default 11)


Example:

    zibr file.zip

will produce "file.zip.br".

Example:

    zibr file.png

will produce "file.png.br".


See https://twitter.com/dchest/status/1488449100072857600
