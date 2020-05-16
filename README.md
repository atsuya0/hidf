# USAGE
```
// Hide
$ date > date.txt
$ cat date.txt
Sat May 16 18:18:01 JST 2020
$ hidf date.txt
$ ls
date.png
$ file date.png
date.png: PNG image data, 500 x 500, 8-bit/color RGB, non-interlaced

// Extract
$ hidf date.png
$ cat date.txt
Sat May 16 18:18:01 JST 2020
```
