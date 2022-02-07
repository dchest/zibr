zibr
====

zipr is a command-line utility that repacks compressed ZIP files into
"stored" ZIP files (ZIP without compression) or PNG files into
uncompressed, but valid, PNG files, and then compresses the result
with brotli. These smaller files may be served by a web server in
response to requests with `Accept-Encoding: br` to save traffic.

See https://twitter.com/dchest/status/1488449100072857600

    Repack ZIP or PNG files with brotli
    Usage: zibr infile [outfile]
    -c int
            brotli compression level (0 - 11) (default 11)
    -utf8
            make sure repacked ZIP files have UTF-8 encoding even if the original doesn't


Example:

    zibr file.zip

will produce "file.zip.br".

Example:

    zibr file.png

will produce "file.png.br".

If passed -utf8 flag, ZIP files repacked with zibr will have UTF8 flag set,
even if the original doesn't. This is not useful for compression, it's just
a fix for annoying macOS ZIP packer behavior which doesn't set the flag,
while encoding file names in UTF-8, making some decompressors (such as
The Unarchiver) sometimes detect a wrong encoding, mangling file names
(as a maker of [Mémoires.app](https://www.codingrobots.com/memoires/), I've
been burned by that.)
