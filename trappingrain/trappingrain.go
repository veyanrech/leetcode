package trappingrain

import "fmt"

func trap(h []int) int {

	n := len(h) //dlina heights

	drops := make([]int, n, n)

	for li := 0; li < n; li++ {

		hli := h[li]

		for ri := li + 1; ri < n; ri++ {

			hri := h[ri]

			if hri == hli {
				break
			}

			if hri > hli {
				break
			}

			if hri < hli {
				drops[ri] = hli - hri
			}

		}
	}

	fmt.Println(drops)
	return 0
}

func TRAP() {
	h := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	trap(h)
}
