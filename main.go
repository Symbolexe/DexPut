package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"flag"
	"fmt"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

// Function to display the help message
func displayHelp() {
	helpMessage := `
Usage: go run DEXPUT.go [OPTIONS]

Options:
  -text         Text to hash
  -hashes       Comma-separated list of hashes to generate (all, md5, sha1, sha256, sha512, sha3-256, sha3-512, blake2b-256, blake2b-512, ripemd160, md4)
  -interactive  Interactive mode
  -benchmark    Benchmark mode
  -help         Display this help message

Examples:
  go run DEXPUT.go -text "hello"
  go run DEXPUT.go -interactive
  go run DEXPUT.go -text "hello" -hashes "md5,sha1"
  go run DEXPUT.go -text "hello" -benchmark
`
	fmt.Println(helpMessage)
}

func hashMD5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func hashSHA1(text string) string {
	hash := sha1.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func hashSHA256(text string) string {
	hash := sha256.Sum256([]byte(text))
	return hex.EncodeToString(hash[:])
}

func hashSHA512(text string) string {
	hash := sha512.Sum512([]byte(text))
	return hex.EncodeToString(hash[:])
}

func hashSHA3_256(text string) string {
	hash := sha3.Sum256([]byte(text))
	return hex.EncodeToString(hash[:])
}

func hashSHA3_512(text string) string {
	hash := sha3.Sum512([]byte(text))
	return hex.EncodeToString(hash[:])
}

func hashBLAKE2b_256(text string) string {
	hash := blake2b.Sum256([]byte(text))
	return hex.EncodeToString(hash[:])
}

func hashBLAKE2b_512(text string) string {
	hash := blake2b.Sum512([]byte(text))
	return hex.EncodeToString(hash[:])
}

func hashRIPEMD160(text string) string {
	hash := ripemd160.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}

func hashMD4(text string) string {
	hash := md4.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}

func main() {
	// Command-line flags
	textPtr := flag.String("text", "", "Text to hash")
	hashPtr := flag.String("hashes", "all", "Comma-separated list of hashes to generate (all, md5, sha1, sha256, sha512, sha3-256, sha3-512, blake2b-256, blake2b-512, ripemd160, md4)")
	interactivePtr := flag.Bool("interactive", false, "Interactive mode")
	benchmarkPtr := flag.Bool("benchmark", false, "Benchmark mode")
	helpPtr := flag.Bool("help", false, "Display this help message")
	flag.Parse()

	if *helpPtr {
		displayHelp()
		os.Exit(0)
	}

	if len(os.Args) == 1 {
		displayHelp()
		os.Exit(1)
	}

	if *interactivePtr {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text to hash: ")
		input, _ := reader.ReadString('\n')
		*textPtr = strings.TrimSpace(input)

		fmt.Print("Enter hashes to generate (comma-separated, 'all' for all hashes): ")
		input, _ = reader.ReadString('\n')
		*hashPtr = strings.TrimSpace(input)
	}

	if *textPtr == "" {
		if !*interactivePtr {
			fmt.Println("No text provided. Use -text or -interactive flag.")
			displayHelp()
			os.Exit(1)
		}
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text to hash: ")
		input, _ := reader.ReadString('\n')
		*textPtr = strings.TrimSpace(input)
	}

	selectedHashes := strings.Split(*hashPtr, ",")
	allHashes := map[string]func(string) string{
		"md5":         hashMD5,
		"sha1":        hashSHA1,
		"sha256":      hashSHA256,
		"sha512":      hashSHA512,
		"sha3-256":    hashSHA3_256,
		"sha3-512":    hashSHA3_512,
		"blake2b-256": hashBLAKE2b_256,
		"blake2b-512": hashBLAKE2b_512,
		"ripemd160":   hashRIPEMD160,
		"md4":         hashMD4,
	}

	if *benchmarkPtr {
		for name, hashFunc := range allHashes {
			start := time.Now()
			hashFunc(*textPtr)
			elapsed := time.Since(start)
			fmt.Printf("%s: %s\n", name, elapsed)
		}
		os.Exit(0)
	}

	var result strings.Builder
	for name, hashFunc := range allHashes {
		if *hashPtr == "all" || contains(selectedHashes, name) {
			hash := hashFunc(*textPtr)
			color.Cyan("%s: %s", name, hash)
			result.WriteString(fmt.Sprintf("%s: %s\n", name, hash))
		}
	}

	err := ioutil.WriteFile("DexPut_Result.txt", []byte(result.String()), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		os.Exit(1)
	}
}

func contains(slice []string, item string) bool {
	for _, a := range slice {
		if strings.TrimSpace(a) == item {
			return true
		}
	}
	return false
}
