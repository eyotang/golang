package main

import ("fmt")

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}

	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a)

	b := [...]string{"abc", "ef"}
	fmt.Println(equal(b[:], b[:]))

	var runes []rune
	for _, r := range "Hello, world" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)
}
