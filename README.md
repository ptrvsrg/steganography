# LSB steganography

In this repository, you can find the source code for LSB steganography implemented in Go.

## Demo

|              Original              |                  Encoded                  |
|:----------------------------------:|:-----------------------------------------:|
| ![Original File](test/statham.png) | ![Encoded File](test/encoded_statham.png) |

The second image contains Jason Statham meme quote from file [message.txt](test/message.txt).

## How to use

Makefile help message:

```sh
$ make help
Available commands:
        make help               - print this help
        make build              - build executable
        make clean              - clean build directory
        make test               - run tests
```

After executing `make build`, you can find the compiled binary in the `build` directory.

```sh
$ make build
...
$ tree build
build
└── lsb
```

More information about this CLI application can be found in the help message:

```sh
$ ./build/lsb
NAME:
   lsb - Tool for LSB steganography on images

USAGE:
   lsb [global options] [command [command options]]

VERSION:
   0.0.1

AUTHOR:
   ptrvsrg

COMMANDS:
   encode, e, E  encode a message to a given image file
   decode, d, D  decode a message from a given image file
   help, h       Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version

COPYRIGHT:
   © 2024 ptrvsrg
```

```sh
$ ./build/lsb help encode
NAME:
   lsb encode - encode a message to a given image file

USAGE:
   lsb encode

OPTIONS:
   --message value, -m value, -M value           message
   --message-file value, --im value, --IM value  input message path
   --input-file value, -i value, -I value        input image path
   --output-file value, -o value, -O value       output image path
   --help, -h                                    show help
```

```sh
$ ./build/lsb help decode                                                                                                                                   
NAME:
   lsb decode - decode a message from a given image file

USAGE:
   lsb decode

OPTIONS:
   --input-file value, -i value, -I value        input image path
   --message-file value, --om value, --OM value  output message path
   --help, -h                                    show help
```

## Testing

```sh
$ make test
?       steganography/cmd/lsb   [no test files]
?       steganography/internal/helper   [no test files]
=== RUN   TestEncodeFromPngFile
--- PASS: TestEncodeFromPngFile (0.11s)
=== RUN   TestDecodeFromPngFile
--- PASS: TestDecodeFromPngFile (0.02s)
=== RUN   TestEncodeFromJpgFile
--- PASS: TestEncodeFromJpgFile (0.04s)
=== RUN   TestDecodeFromJpgFile
--- PASS: TestDecodeFromJpgFile (0.01s)
=== RUN   TestEncodeFromJpegFile
--- PASS: TestEncodeFromJpegFile (0.01s)
=== RUN   TestDecodeFromJpegFile
--- PASS: TestDecodeFromJpegFile (0.00s)
PASS
ok      steganography/internal/lsb      0.668s
```

## License

This project is released under the [MIT License](https://github.com/ptrvsrg/crypto/blob/master/LICENSE).
