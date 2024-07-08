# DexPut

DexPut is a command-line tool written in Go for generating various hashes (MD5, SHA-1, SHA-256, SHA-512, SHA-3, BLAKE2, RIPEMD-160, and MD4) for a given text. It provides an interactive mode, allows selection of specific hashes, and includes benchmarking capabilities.

## Features

- Generate multiple hashes for a given text
- Interactive mode for user-friendly input
- Select specific hashes to generate
- Benchmark mode to measure the performance of each hashing algorithm
- Results are displayed in the console with colored output and saved to a file

## Usage

### Standard Usage

```sh
go run main.go -text "hello"
```

This command will generate all hashes for the input text "hello", print them in color to the console, and save the results to `DexPut_Result.txt`.

### Interactive Mode
```sh
go run main.go -interactive
```

This command will prompt the user to enter the text and hashes interactively.

### Select Specific Hashes
```sh
go run main.go -text "hello" -hashes "md5,sha1"
```

This command will generate only the MD5 and SHA-1 hashes for the input text "hello".

### Benchmark Mode
```sh
go run main.go -text "hello" -benchmark
```

This command will benchmark the time taken for each hashing function and print the results.

### Display Help
```sh
go run main.go -help
```

This command will display the help message with detailed usage information.

### No Arguments
```sh
go run main.go
```

This command will display the help message if no arguments are provided.

## Installation
- [x] Clone the repository
```sh
git clone https://github.com/Symbolexe/DexPut.git
cd DexPut
```
- [x] Initialize a Go module
```sh
go mod init DexPut
```
- [x] Add the required dependencies
```sh
go get github.com/fatih/color
go get golang.org/x/crypto/blake2b
go get golang.org/x/crypto/md4
go get golang.org/x/crypto/ripemd160
go get golang.org/x/crypto/sha3
```
- [x] Run the tool
```sh
go run main.go -text "hello"
```
- [x] use build

Download the build version from Releases.






