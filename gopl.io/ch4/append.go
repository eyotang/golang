package main

import ("fmt")

func appendInt(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2*len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	copy(z[len(x):], y)
	return z
}


func main() {
	a := [...]int{0, 1, 2, 5}
	b := appendInt(a[:], 8)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("len(a): %d, cap(a): %d, len(b): %d, cap(b): %d\n", len(a), cap(a), len(b), cap(b))

	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}

	var runes []int
	for i := 0; i < 10; i++ {
		runes = append(runes, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(runes), runes)
	}

	b = appendInt(a[:], 3, 4, 9)
	fmt.Println(b)
}
