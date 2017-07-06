package main

import (
	"fmt"
	"os"
	"crypto/sha256"
	"crypto/sha512"
)


func main() {
	switch os.Args[1] {
	case "sha256":
		c1 := sha256.Sum256([]byte("x"))
		c2 := sha256.Sum256([]byte("X"))
		fmt.Printf("Sum256:\n%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	case "sha384":
		c1 := sha512.Sum384([]byte("x"))
		c2 := sha512.Sum384([]byte("X"))
		fmt.Printf("Sum384:\n%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	case "sha512":
		c1 := sha512.Sum512([]byte("x"))
		c2 := sha512.Sum512([]byte("X"))
		fmt.Printf("Sum512:\n%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	default:
		fmt.Printf("can't find the algo!\n")
	}
}
