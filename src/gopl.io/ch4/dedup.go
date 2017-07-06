package main

import (
	"os"
	"fmt"
	"bufio"
)

func k(list []string) string {
	return fmt.Sprintf("%q", list)
}

func Add(m map[string]int, list []string) {
	m[k(list)]++
}

func Count(m map[string]int, list []string) int {
	return m[k(list)]
}

func main() {
	var rank = make(map[string]int)

	people := [...]string{"kris", "lily"}
	fmt.Printf("Operation for: %s\n", k(people[:]))
	fmt.Printf("Before: %d\n", Count(rank, people[:]))
	Add(rank, people[:])
	fmt.Printf("After: %d\n", Count(rank, people[:]))


	seen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}
