package golangwiki

import "fmt"

//https://github.com/golang/go/wiki/SliceTricks#filter-in-place

func Filter(a []int) []int {
	n := 0
	for _, x := range a {
		if keep(x) {
			a[n] = x
			n++
		}
	}
	a = a[:n]

	return a
}

func keep(i int) bool {
	return i%2 == 0
}

func Filter2(a []int) {
	b := a[:0]
	for _, x := range a {
		if keep(x) {
			b = append(b, x)
		}
	}

	fmt.Println(b)
	fmt.Println(a)
}
