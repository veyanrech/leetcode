package matrix

import "fmt"

func spiralOrder(matrix [][]int) []int {

	rowl := len(matrix[0])
	coll := len(matrix)

	n := rowl * coll

	h := 0
	v := 0

	edgeh := []int{0, rowl - 1}
	edgev := []int{1, coll - 1}

	res := []int{}

	right := true
	down := true

	goh := true

	for i := 0; i < n; i++ {

		fmt.Println(i, h, v, matrix[v][h])
		res = append(res, matrix[v][h])

		if goh {
			if right {
				h++
			} else {
				h--
			}

			if h > edgeh[1] || h < edgeh[0] {
				if h > edgeh[1] {
					h = edgeh[1]
					edgeh[1] = edgeh[1] - 1
					v++
				}
				if h < edgeh[0] {
					h = edgeh[0]
					edgeh[0] = edgeh[0] + 1
					v--
				}
				goh = !goh
				right = !right
			}

		} else {
			if down {
				v++
			} else {
				v--
			}

			if v > edgev[1] || v < edgev[0] {
				if v > edgev[1] {
					v = edgev[1]
					edgev[1] = edgev[1] - 1
					h--
				}
				if v < edgev[0] {
					v = edgev[0]
					edgev[0] = edgev[0] + 1
					h++
				}
				goh = !goh
				down = !down
			}
		}

	}

	return res
}

func Run() {
	test := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(spiralOrder(test))
}
