package main

import (
	"fmt"
	"sort"
)

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}

	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}

	return true
}

func main() {
	ages := map[string]int {
		"alice": 31,
		"charlie": 34,
		"boris": 38,
	}
	fmt.Println("map visit via (k, v) random:")
	for k, v := range ages {
		fmt.Printf("%s\t%d\n", k, v)
	}

	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}

	fmt.Println("map visit via sorted key:")
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}

	empty_ages := map[string]int {
		"frank" : 80,
	}

	fmt.Println(equal(ages, empty_ages))
}
